// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	operatorv1alpha1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageContentSourcePolicies implements ImageContentSourcePolicyInterface
type FakeImageContentSourcePolicies struct {
	Fake *FakeOperatorV1alpha1
}

var imagecontentsourcepoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("imagecontentsourcepolicies")

var imagecontentsourcepoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("ImageContentSourcePolicy")

// Get takes name of the imageContentSourcePolicy, and returns the corresponding imageContentSourcePolicy object, and an error if there is any.
func (c *FakeImageContentSourcePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(imagecontentsourcepoliciesResource, name), &v1alpha1.ImageContentSourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageContentSourcePolicy), err
}

// List takes label and field selectors, and returns the list of ImageContentSourcePolicies that match those selectors.
func (c *FakeImageContentSourcePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageContentSourcePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(imagecontentsourcepoliciesResource, imagecontentsourcepoliciesKind, opts), &v1alpha1.ImageContentSourcePolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageContentSourcePolicyList{ListMeta: obj.(*v1alpha1.ImageContentSourcePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageContentSourcePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageContentSourcePolicies.
func (c *FakeImageContentSourcePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(imagecontentsourcepoliciesResource, opts))
}

// Create takes the representation of a imageContentSourcePolicy and creates it.  Returns the server's representation of the imageContentSourcePolicy, and an error, if there is any.
func (c *FakeImageContentSourcePolicies) Create(ctx context.Context, imageContentSourcePolicy *v1alpha1.ImageContentSourcePolicy, opts v1.CreateOptions) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(imagecontentsourcepoliciesResource, imageContentSourcePolicy), &v1alpha1.ImageContentSourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageContentSourcePolicy), err
}

// Update takes the representation of a imageContentSourcePolicy and updates it. Returns the server's representation of the imageContentSourcePolicy, and an error, if there is any.
func (c *FakeImageContentSourcePolicies) Update(ctx context.Context, imageContentSourcePolicy *v1alpha1.ImageContentSourcePolicy, opts v1.UpdateOptions) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(imagecontentsourcepoliciesResource, imageContentSourcePolicy), &v1alpha1.ImageContentSourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageContentSourcePolicy), err
}

// Delete takes name of the imageContentSourcePolicy and deletes it. Returns an error if one occurs.
func (c *FakeImageContentSourcePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(imagecontentsourcepoliciesResource, name, opts), &v1alpha1.ImageContentSourcePolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageContentSourcePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(imagecontentsourcepoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageContentSourcePolicyList{})
	return err
}

// Patch applies the patch and returns the patched imageContentSourcePolicy.
func (c *FakeImageContentSourcePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imagecontentsourcepoliciesResource, name, pt, data, subresources...), &v1alpha1.ImageContentSourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageContentSourcePolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied imageContentSourcePolicy.
func (c *FakeImageContentSourcePolicies) Apply(ctx context.Context, imageContentSourcePolicy *operatorv1alpha1.ImageContentSourcePolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	if imageContentSourcePolicy == nil {
		return nil, fmt.Errorf("imageContentSourcePolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(imageContentSourcePolicy)
	if err != nil {
		return nil, err
	}
	name := imageContentSourcePolicy.Name
	if name == nil {
		return nil, fmt.Errorf("imageContentSourcePolicy.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imagecontentsourcepoliciesResource, *name, types.ApplyPatchType, data), &v1alpha1.ImageContentSourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageContentSourcePolicy), err
}
