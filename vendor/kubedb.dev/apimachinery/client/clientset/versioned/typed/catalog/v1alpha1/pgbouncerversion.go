/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PgBouncerVersionsGetter has a method to return a PgBouncerVersionInterface.
// A group's client should implement this interface.
type PgBouncerVersionsGetter interface {
	PgBouncerVersions() PgBouncerVersionInterface
}

// PgBouncerVersionInterface has methods to work with PgBouncerVersion resources.
type PgBouncerVersionInterface interface {
	Create(ctx context.Context, pgBouncerVersion *v1alpha1.PgBouncerVersion, opts v1.CreateOptions) (*v1alpha1.PgBouncerVersion, error)
	Update(ctx context.Context, pgBouncerVersion *v1alpha1.PgBouncerVersion, opts v1.UpdateOptions) (*v1alpha1.PgBouncerVersion, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PgBouncerVersion, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PgBouncerVersionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PgBouncerVersion, err error)
	PgBouncerVersionExpansion
}

// pgBouncerVersions implements PgBouncerVersionInterface
type pgBouncerVersions struct {
	client rest.Interface
}

// newPgBouncerVersions returns a PgBouncerVersions
func newPgBouncerVersions(c *CatalogV1alpha1Client) *pgBouncerVersions {
	return &pgBouncerVersions{
		client: c.RESTClient(),
	}
}

// Get takes name of the pgBouncerVersion, and returns the corresponding pgBouncerVersion object, and an error if there is any.
func (c *pgBouncerVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PgBouncerVersion, err error) {
	result = &v1alpha1.PgBouncerVersion{}
	err = c.client.Get().
		Resource("pgbouncerversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PgBouncerVersions that match those selectors.
func (c *pgBouncerVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PgBouncerVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PgBouncerVersionList{}
	err = c.client.Get().
		Resource("pgbouncerversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pgBouncerVersions.
func (c *pgBouncerVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("pgbouncerversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a pgBouncerVersion and creates it.  Returns the server's representation of the pgBouncerVersion, and an error, if there is any.
func (c *pgBouncerVersions) Create(ctx context.Context, pgBouncerVersion *v1alpha1.PgBouncerVersion, opts v1.CreateOptions) (result *v1alpha1.PgBouncerVersion, err error) {
	result = &v1alpha1.PgBouncerVersion{}
	err = c.client.Post().
		Resource("pgbouncerversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pgBouncerVersion).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a pgBouncerVersion and updates it. Returns the server's representation of the pgBouncerVersion, and an error, if there is any.
func (c *pgBouncerVersions) Update(ctx context.Context, pgBouncerVersion *v1alpha1.PgBouncerVersion, opts v1.UpdateOptions) (result *v1alpha1.PgBouncerVersion, err error) {
	result = &v1alpha1.PgBouncerVersion{}
	err = c.client.Put().
		Resource("pgbouncerversions").
		Name(pgBouncerVersion.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pgBouncerVersion).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the pgBouncerVersion and deletes it. Returns an error if one occurs.
func (c *pgBouncerVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pgbouncerversions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pgBouncerVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("pgbouncerversions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched pgBouncerVersion.
func (c *pgBouncerVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PgBouncerVersion, err error) {
	result = &v1alpha1.PgBouncerVersion{}
	err = c.client.Patch(pt).
		Resource("pgbouncerversions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
