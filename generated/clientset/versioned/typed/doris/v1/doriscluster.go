/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/selectdb/doris-operator/api/doris/v1"
	scheme "github.com/selectdb/doris-operator/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DorisClustersGetter has a method to return a DorisClusterInterface.
// A group's client should implement this interface.
type DorisClustersGetter interface {
	DorisClusters(namespace string) DorisClusterInterface
}

// DorisClusterInterface has methods to work with DorisCluster resources.
type DorisClusterInterface interface {
	Create(ctx context.Context, dorisCluster *v1.DorisCluster, opts metav1.CreateOptions) (*v1.DorisCluster, error)
	Update(ctx context.Context, dorisCluster *v1.DorisCluster, opts metav1.UpdateOptions) (*v1.DorisCluster, error)
	UpdateStatus(ctx context.Context, dorisCluster *v1.DorisCluster, opts metav1.UpdateOptions) (*v1.DorisCluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DorisCluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DorisClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DorisCluster, err error)
	DorisClusterExpansion
}

// dorisClusters implements DorisClusterInterface
type dorisClusters struct {
	client rest.Interface
	ns     string
}

// newDorisClusters returns a DorisClusters
func newDorisClusters(c *DorisV1Client, namespace string) *dorisClusters {
	return &dorisClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dorisCluster, and returns the corresponding dorisCluster object, and an error if there is any.
func (c *dorisClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DorisCluster, err error) {
	result = &v1.DorisCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dorisclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DorisClusters that match those selectors.
func (c *dorisClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DorisClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DorisClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dorisclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dorisClusters.
func (c *dorisClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dorisclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dorisCluster and creates it.  Returns the server's representation of the dorisCluster, and an error, if there is any.
func (c *dorisClusters) Create(ctx context.Context, dorisCluster *v1.DorisCluster, opts metav1.CreateOptions) (result *v1.DorisCluster, err error) {
	result = &v1.DorisCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dorisclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dorisCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dorisCluster and updates it. Returns the server's representation of the dorisCluster, and an error, if there is any.
func (c *dorisClusters) Update(ctx context.Context, dorisCluster *v1.DorisCluster, opts metav1.UpdateOptions) (result *v1.DorisCluster, err error) {
	result = &v1.DorisCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dorisclusters").
		Name(dorisCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dorisCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dorisClusters) UpdateStatus(ctx context.Context, dorisCluster *v1.DorisCluster, opts metav1.UpdateOptions) (result *v1.DorisCluster, err error) {
	result = &v1.DorisCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dorisclusters").
		Name(dorisCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dorisCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dorisCluster and deletes it. Returns an error if one occurs.
func (c *dorisClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dorisclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dorisClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dorisclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dorisCluster.
func (c *dorisClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DorisCluster, err error) {
	result = &v1.DorisCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dorisclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
