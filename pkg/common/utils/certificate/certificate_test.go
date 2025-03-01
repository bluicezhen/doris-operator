// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package certificate

import (
	"crypto/x509/pkix"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func Test_ValidCA(t *testing.T) {

	dnsNames := []string{
		fmt.Sprintf("%s.%s", "service", "namespace"),
		fmt.Sprintf("%s.%s.svc", "service", "namespace"),
		fmt.Sprintf("%s.%s.svc.cluster.local", "service", "namespace"),
	}

	caop := CAOptions{
		Subject: pkix.Name{
			CommonName:   "test" + "-" + "HTTP",
			Organization: []string{"doris-operator"},
		},
		DnsNames: dnsNames,
	}
	ca, err := NewCAConfigSecret(caop)
	if err != nil {
		t.Errorf("validCA test new ca config from secret failed, err=%s", err.Error())
	}

	s := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service",
			Namespace: "namespace",
		},
	}

	s.Data = make(map[string][]byte, 2)
	s.Data[TlsKeyName] = ca.GetEncodePrivateKey()
	s.Data[TLsCertName] = ca.GetEncodeCert()

	check_ca := BuildCAFromSecret(s)

	res := ValidCA(check_ca)
	if res != true {
		t.Errorf("validCA test failed, the ca not valid.")
	}
}
