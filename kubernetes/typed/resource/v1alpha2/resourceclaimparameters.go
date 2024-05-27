/*
Copyright The Kubernetes Authors.

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

package v1alpha2

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha2 "k8s.io/api/resource/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	resourcev1alpha2 "k8s.io/client-go/applyconfigurations/resource/v1alpha2"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
	consistencydetector "k8s.io/client-go/util/consistencydetector"
)

// ResourceClaimParametersGetter has a method to return a ResourceClaimParametersInterface.
// A group's client should implement this interface.
type ResourceClaimParametersGetter interface {
	ResourceClaimParameters(namespace string) ResourceClaimParametersInterface
}

// ResourceClaimParametersInterface has methods to work with ResourceClaimParameters resources.
type ResourceClaimParametersInterface interface {
	Create(ctx context.Context, resourceClaimParameters *v1alpha2.ResourceClaimParameters, opts v1.CreateOptions) (*v1alpha2.ResourceClaimParameters, error)
	Update(ctx context.Context, resourceClaimParameters *v1alpha2.ResourceClaimParameters, opts v1.UpdateOptions) (*v1alpha2.ResourceClaimParameters, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ResourceClaimParameters, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ResourceClaimParametersList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ResourceClaimParameters, err error)
	Apply(ctx context.Context, resourceClaimParameters *resourcev1alpha2.ResourceClaimParametersApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.ResourceClaimParameters, err error)
	ResourceClaimParametersExpansion
}

// resourceClaimParameters implements ResourceClaimParametersInterface
type resourceClaimParameters struct {
	client rest.Interface
	ns     string
}

// newResourceClaimParameters returns a ResourceClaimParameters
func newResourceClaimParameters(c *ResourceV1alpha2Client, namespace string) *resourceClaimParameters {
	return &resourceClaimParameters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceClaimParameters, and returns the corresponding resourceClaimParameters object, and an error if there is any.
func (c *resourceClaimParameters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ResourceClaimParameters, err error) {
	result = &v1alpha2.ResourceClaimParameters{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceClaimParameters that match those selectors.
func (c *resourceClaimParameters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ResourceClaimParametersList, err error) {
	defer func() {
		if err == nil {
			consistencydetector.CheckListFromCacheDataConsistencyIfRequested(ctx, "list request for resourceclaimparameters", c.list, opts, result)
		}
	}()
	return c.list(ctx, opts)
}

// list takes label and field selectors, and returns the list of ResourceClaimParameters that match those selectors.
func (c *resourceClaimParameters) list(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ResourceClaimParametersList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ResourceClaimParametersList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceClaimParameters.
func (c *resourceClaimParameters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceClaimParameters and creates it.  Returns the server's representation of the resourceClaimParameters, and an error, if there is any.
func (c *resourceClaimParameters) Create(ctx context.Context, resourceClaimParameters *v1alpha2.ResourceClaimParameters, opts v1.CreateOptions) (result *v1alpha2.ResourceClaimParameters, err error) {
	result = &v1alpha2.ResourceClaimParameters{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceClaimParameters).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceClaimParameters and updates it. Returns the server's representation of the resourceClaimParameters, and an error, if there is any.
func (c *resourceClaimParameters) Update(ctx context.Context, resourceClaimParameters *v1alpha2.ResourceClaimParameters, opts v1.UpdateOptions) (result *v1alpha2.ResourceClaimParameters, err error) {
	result = &v1alpha2.ResourceClaimParameters{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		Name(resourceClaimParameters.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceClaimParameters).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceClaimParameters and deletes it. Returns an error if one occurs.
func (c *resourceClaimParameters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceClaimParameters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceClaimParameters.
func (c *resourceClaimParameters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ResourceClaimParameters, err error) {
	result = &v1alpha2.ResourceClaimParameters{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceClaimParameters.
func (c *resourceClaimParameters) Apply(ctx context.Context, resourceClaimParameters *resourcev1alpha2.ResourceClaimParametersApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.ResourceClaimParameters, err error) {
	if resourceClaimParameters == nil {
		return nil, fmt.Errorf("resourceClaimParameters provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(resourceClaimParameters)
	if err != nil {
		return nil, err
	}
	name := resourceClaimParameters.Name
	if name == nil {
		return nil, fmt.Errorf("resourceClaimParameters.Name must be provided to Apply")
	}
	result = &v1alpha2.ResourceClaimParameters{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("resourceclaimparameters").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
