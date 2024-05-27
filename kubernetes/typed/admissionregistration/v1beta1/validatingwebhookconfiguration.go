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

package v1beta1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1beta1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
	consistencydetector "k8s.io/client-go/util/consistencydetector"
)

// ValidatingWebhookConfigurationsGetter has a method to return a ValidatingWebhookConfigurationInterface.
// A group's client should implement this interface.
type ValidatingWebhookConfigurationsGetter interface {
	ValidatingWebhookConfigurations() ValidatingWebhookConfigurationInterface
}

// ValidatingWebhookConfigurationInterface has methods to work with ValidatingWebhookConfiguration resources.
type ValidatingWebhookConfigurationInterface interface {
	Create(ctx context.Context, validatingWebhookConfiguration *v1beta1.ValidatingWebhookConfiguration, opts v1.CreateOptions) (*v1beta1.ValidatingWebhookConfiguration, error)
	Update(ctx context.Context, validatingWebhookConfiguration *v1beta1.ValidatingWebhookConfiguration, opts v1.UpdateOptions) (*v1beta1.ValidatingWebhookConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ValidatingWebhookConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ValidatingWebhookConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ValidatingWebhookConfiguration, err error)
	Apply(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1beta1.ValidatingWebhookConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.ValidatingWebhookConfiguration, err error)
	ValidatingWebhookConfigurationExpansion
}

// validatingWebhookConfigurations implements ValidatingWebhookConfigurationInterface
type validatingWebhookConfigurations struct {
	client rest.Interface
}

// newValidatingWebhookConfigurations returns a ValidatingWebhookConfigurations
func newValidatingWebhookConfigurations(c *AdmissionregistrationV1beta1Client) *validatingWebhookConfigurations {
	return &validatingWebhookConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the validatingWebhookConfiguration, and returns the corresponding validatingWebhookConfiguration object, and an error if there is any.
func (c *validatingWebhookConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ValidatingWebhookConfiguration, err error) {
	result = &v1beta1.ValidatingWebhookConfiguration{}
	err = c.client.Get().
		Resource("validatingwebhookconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ValidatingWebhookConfigurations that match those selectors.
func (c *validatingWebhookConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ValidatingWebhookConfigurationList, err error) {
	defer func() {
		if err == nil {
			consistencydetector.CheckListFromCacheDataConsistencyIfRequested(ctx, "list request for validatingwebhookconfigurations", c.list, opts, result)
		}
	}()
	return c.list(ctx, opts)
}

// list takes label and field selectors, and returns the list of ValidatingWebhookConfigurations that match those selectors.
func (c *validatingWebhookConfigurations) list(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ValidatingWebhookConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ValidatingWebhookConfigurationList{}
	err = c.client.Get().
		Resource("validatingwebhookconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested validatingWebhookConfigurations.
func (c *validatingWebhookConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("validatingwebhookconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a validatingWebhookConfiguration and creates it.  Returns the server's representation of the validatingWebhookConfiguration, and an error, if there is any.
func (c *validatingWebhookConfigurations) Create(ctx context.Context, validatingWebhookConfiguration *v1beta1.ValidatingWebhookConfiguration, opts v1.CreateOptions) (result *v1beta1.ValidatingWebhookConfiguration, err error) {
	result = &v1beta1.ValidatingWebhookConfiguration{}
	err = c.client.Post().
		Resource("validatingwebhookconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(validatingWebhookConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a validatingWebhookConfiguration and updates it. Returns the server's representation of the validatingWebhookConfiguration, and an error, if there is any.
func (c *validatingWebhookConfigurations) Update(ctx context.Context, validatingWebhookConfiguration *v1beta1.ValidatingWebhookConfiguration, opts v1.UpdateOptions) (result *v1beta1.ValidatingWebhookConfiguration, err error) {
	result = &v1beta1.ValidatingWebhookConfiguration{}
	err = c.client.Put().
		Resource("validatingwebhookconfigurations").
		Name(validatingWebhookConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(validatingWebhookConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the validatingWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *validatingWebhookConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("validatingwebhookconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *validatingWebhookConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("validatingwebhookconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched validatingWebhookConfiguration.
func (c *validatingWebhookConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ValidatingWebhookConfiguration, err error) {
	result = &v1beta1.ValidatingWebhookConfiguration{}
	err = c.client.Patch(pt).
		Resource("validatingwebhookconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied validatingWebhookConfiguration.
func (c *validatingWebhookConfigurations) Apply(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1beta1.ValidatingWebhookConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.ValidatingWebhookConfiguration, err error) {
	if validatingWebhookConfiguration == nil {
		return nil, fmt.Errorf("validatingWebhookConfiguration provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(validatingWebhookConfiguration)
	if err != nil {
		return nil, err
	}
	name := validatingWebhookConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("validatingWebhookConfiguration.Name must be provided to Apply")
	}
	result = &v1beta1.ValidatingWebhookConfiguration{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("validatingwebhookconfigurations").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
