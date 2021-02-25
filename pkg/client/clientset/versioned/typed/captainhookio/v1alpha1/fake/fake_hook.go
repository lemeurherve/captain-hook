/*
 Generated Code
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/garethjevans/captain-hook/pkg/api/captainhookio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHooks implements HookInterface
type FakeHooks struct {
	Fake *FakeCaptainhookV1alpha1
	ns   string
}

var hooksResource = schema.GroupVersionResource{Group: "captainhook.io", Version: "v1alpha1", Resource: "hooks"}

var hooksKind = schema.GroupVersionKind{Group: "captainhook.io", Version: "v1alpha1", Kind: "Hook"}

// Get takes name of the hook, and returns the corresponding hook object, and an error if there is any.
func (c *FakeHooks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Hook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hooksResource, c.ns, name), &v1alpha1.Hook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hook), err
}

// List takes label and field selectors, and returns the list of Hooks that match those selectors.
func (c *FakeHooks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hooksResource, hooksKind, c.ns, opts), &v1alpha1.HookList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HookList{ListMeta: obj.(*v1alpha1.HookList).ListMeta}
	for _, item := range obj.(*v1alpha1.HookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hooks.
func (c *FakeHooks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hooksResource, c.ns, opts))

}

// Create takes the representation of a hook and creates it.  Returns the server's representation of the hook, and an error, if there is any.
func (c *FakeHooks) Create(ctx context.Context, hook *v1alpha1.Hook, opts v1.CreateOptions) (result *v1alpha1.Hook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hooksResource, c.ns, hook), &v1alpha1.Hook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hook), err
}

// Update takes the representation of a hook and updates it. Returns the server's representation of the hook, and an error, if there is any.
func (c *FakeHooks) Update(ctx context.Context, hook *v1alpha1.Hook, opts v1.UpdateOptions) (result *v1alpha1.Hook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hooksResource, c.ns, hook), &v1alpha1.Hook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hook), err
}

// Delete takes name of the hook and deletes it. Returns an error if one occurs.
func (c *FakeHooks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hooksResource, c.ns, name), &v1alpha1.Hook{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHooks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hooksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HookList{})
	return err
}

// Patch applies the patch and returns the patched hook.
func (c *FakeHooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Hook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hooksResource, c.ns, name, pt, data, subresources...), &v1alpha1.Hook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hook), err
}
