//-----------------------------------------------------------------------------
// Demo API
//-----------------------------------------------------------------------------

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	demov1 "github.com/laforge-tech/k8sgen/pkg/api/demo/v1"
	applyconfigurationsdemov1 "github.com/laforge-tech/k8sgen/pkg/applyconfigurations/demo/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDemos implements DemoInterface
type FakeDemos struct {
	Fake *FakeDemoV1
	ns   string
}

var demosResource = schema.GroupVersionResource{Group: "demo", Version: "v1", Resource: "demos"}

var demosKind = schema.GroupVersionKind{Group: "demo", Version: "v1", Kind: "Demo"}

// Get takes name of the demo, and returns the corresponding demo object, and an error if there is any.
func (c *FakeDemos) Get(ctx context.Context, name string, options v1.GetOptions) (result *demov1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(demosResource, c.ns, name), &demov1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*demov1.Demo), err
}

// List takes label and field selectors, and returns the list of Demos that match those selectors.
func (c *FakeDemos) List(ctx context.Context, opts v1.ListOptions) (result *demov1.DemoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(demosResource, demosKind, c.ns, opts), &demov1.DemoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &demov1.DemoList{ListMeta: obj.(*demov1.DemoList).ListMeta}
	for _, item := range obj.(*demov1.DemoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested demos.
func (c *FakeDemos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(demosResource, c.ns, opts))

}

// Create takes the representation of a demo and creates it.  Returns the server's representation of the demo, and an error, if there is any.
func (c *FakeDemos) Create(ctx context.Context, demo *demov1.Demo, opts v1.CreateOptions) (result *demov1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(demosResource, c.ns, demo), &demov1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*demov1.Demo), err
}

// Update takes the representation of a demo and updates it. Returns the server's representation of the demo, and an error, if there is any.
func (c *FakeDemos) Update(ctx context.Context, demo *demov1.Demo, opts v1.UpdateOptions) (result *demov1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(demosResource, c.ns, demo), &demov1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*demov1.Demo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDemos) UpdateStatus(ctx context.Context, demo *demov1.Demo, opts v1.UpdateOptions) (*demov1.Demo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(demosResource, "status", c.ns, demo), &demov1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*demov1.Demo), err
}

// Delete takes name of the demo and deletes it. Returns an error if one occurs.
func (c *FakeDemos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(demosResource, c.ns, name, opts), &demov1.Demo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDemos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(demosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &demov1.DemoList{})
	return err
}

// Patch applies the patch and returns the patched demo.
func (c *FakeDemos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *demov1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(demosResource, c.ns, name, pt, data, subresources...), &demov1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*demov1.Demo), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied demo.
func (c *FakeDemos) Apply(ctx context.Context, demo *applyconfigurationsdemov1.DemoApplyConfiguration, opts v1.ApplyOptions) (result *demov1.Demo, err error) {
	if demo == nil {
		return nil, fmt.Errorf("demo provided to Apply must not be nil")
	}
	data, err := json.Marshal(demo)
	if err != nil {
		return nil, err
	}
	name := demo.Name
	if name == nil {
		return nil, fmt.Errorf("demo.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(demosResource, c.ns, *name, types.ApplyPatchType, data), &demov1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*demov1.Demo), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDemos) ApplyStatus(ctx context.Context, demo *applyconfigurationsdemov1.DemoApplyConfiguration, opts v1.ApplyOptions) (result *demov1.Demo, err error) {
	if demo == nil {
		return nil, fmt.Errorf("demo provided to Apply must not be nil")
	}
	data, err := json.Marshal(demo)
	if err != nil {
		return nil, err
	}
	name := demo.Name
	if name == nil {
		return nil, fmt.Errorf("demo.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(demosResource, c.ns, *name, types.ApplyPatchType, data, "status"), &demov1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*demov1.Demo), err
}
