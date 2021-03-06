/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RequirementGapsGetter has a method to return a RequirementGapInterface.
// A group's client should implement this interface.
type RequirementGapsGetter interface {
	RequirementGaps() RequirementGapInterface
}

// RequirementGapInterface has methods to work with RequirementGap resources.
type RequirementGapInterface interface {
	Create(*v1beta1.RequirementGap) (*v1beta1.RequirementGap, error)
	Update(*v1beta1.RequirementGap) (*v1beta1.RequirementGap, error)
	UpdateStatus(*v1beta1.RequirementGap) (*v1beta1.RequirementGap, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.RequirementGap, error)
	List(opts v1.ListOptions) (*v1beta1.RequirementGapList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RequirementGap, err error)
	RequirementGapExpansion
}

// requirementGaps implements RequirementGapInterface
type requirementGaps struct {
	client rest.Interface
}

// newRequirementGaps returns a RequirementGaps
func newRequirementGaps(c *ResourceV1beta1Client) *requirementGaps {
	return &requirementGaps{
		client: c.RESTClient(),
	}
}

// Get takes name of the requirementGap, and returns the corresponding requirementGap object, and an error if there is any.
func (c *requirementGaps) Get(name string, options v1.GetOptions) (result *v1beta1.RequirementGap, err error) {
	result = &v1beta1.RequirementGap{}
	err = c.client.Get().
		Resource("requirementgaps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RequirementGaps that match those selectors.
func (c *requirementGaps) List(opts v1.ListOptions) (result *v1beta1.RequirementGapList, err error) {
	result = &v1beta1.RequirementGapList{}
	err = c.client.Get().
		Resource("requirementgaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested requirementGaps.
func (c *requirementGaps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("requirementgaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a requirementGap and creates it.  Returns the server's representation of the requirementGap, and an error, if there is any.
func (c *requirementGaps) Create(requirementGap *v1beta1.RequirementGap) (result *v1beta1.RequirementGap, err error) {
	result = &v1beta1.RequirementGap{}
	err = c.client.Post().
		Resource("requirementgaps").
		Body(requirementGap).
		Do().
		Into(result)
	return
}

// Update takes the representation of a requirementGap and updates it. Returns the server's representation of the requirementGap, and an error, if there is any.
func (c *requirementGaps) Update(requirementGap *v1beta1.RequirementGap) (result *v1beta1.RequirementGap, err error) {
	result = &v1beta1.RequirementGap{}
	err = c.client.Put().
		Resource("requirementgaps").
		Name(requirementGap.Name).
		Body(requirementGap).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *requirementGaps) UpdateStatus(requirementGap *v1beta1.RequirementGap) (result *v1beta1.RequirementGap, err error) {
	result = &v1beta1.RequirementGap{}
	err = c.client.Put().
		Resource("requirementgaps").
		Name(requirementGap.Name).
		SubResource("status").
		Body(requirementGap).
		Do().
		Into(result)
	return
}

// Delete takes name of the requirementGap and deletes it. Returns an error if one occurs.
func (c *requirementGaps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("requirementgaps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *requirementGaps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("requirementgaps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched requirementGap.
func (c *requirementGaps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RequirementGap, err error) {
	result = &v1beta1.RequirementGap{}
	err = c.client.Patch(pt).
		Resource("requirementgaps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
