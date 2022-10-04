//-----------------------------------------------------------------------------
// Demo API
//-----------------------------------------------------------------------------

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/laforge-tech/k8sgen/pkg/api/demo/v1"
	demov1 "github.com/laforge-tech/k8sgen/pkg/applyconfigurations/demo/v1"
	scheme "github.com/laforge-tech/k8sgen/pkg/client/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DemosGetter has a method to return a DemoInterface.
// A group's client should implement this interface.
type DemosGetter interface {
	Demos(namespace string) DemoInterface
}

// DemoInterface has methods to work with Demo resources.
type DemoInterface interface {
	Create(ctx context.Context, demo *v1.Demo, opts metav1.CreateOptions) (*v1.Demo, error)
	Update(ctx context.Context, demo *v1.Demo, opts metav1.UpdateOptions) (*v1.Demo, error)
	UpdateStatus(ctx context.Context, demo *v1.Demo, opts metav1.UpdateOptions) (*v1.Demo, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Demo, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DemoList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Demo, err error)
	Apply(ctx context.Context, demo *demov1.DemoApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Demo, err error)
	ApplyStatus(ctx context.Context, demo *demov1.DemoApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Demo, err error)
	DemoExpansion
}

// demos implements DemoInterface
type demos struct {
	client rest.Interface
	ns     string
}

// newDemos returns a Demos
func newDemos(c *DemoV1Client, namespace string) *demos {
	return &demos{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the demo, and returns the corresponding demo object, and an error if there is any.
func (c *demos) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Demo, err error) {
	result = &v1.Demo{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("demos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Demos that match those selectors.
func (c *demos) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DemoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DemoList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("demos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested demos.
func (c *demos) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("demos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a demo and creates it.  Returns the server's representation of the demo, and an error, if there is any.
func (c *demos) Create(ctx context.Context, demo *v1.Demo, opts metav1.CreateOptions) (result *v1.Demo, err error) {
	result = &v1.Demo{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("demos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(demo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a demo and updates it. Returns the server's representation of the demo, and an error, if there is any.
func (c *demos) Update(ctx context.Context, demo *v1.Demo, opts metav1.UpdateOptions) (result *v1.Demo, err error) {
	result = &v1.Demo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("demos").
		Name(demo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(demo).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *demos) UpdateStatus(ctx context.Context, demo *v1.Demo, opts metav1.UpdateOptions) (result *v1.Demo, err error) {
	result = &v1.Demo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("demos").
		Name(demo.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(demo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the demo and deletes it. Returns an error if one occurs.
func (c *demos) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("demos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *demos) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("demos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched demo.
func (c *demos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Demo, err error) {
	result = &v1.Demo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("demos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied demo.
func (c *demos) Apply(ctx context.Context, demo *demov1.DemoApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Demo, err error) {
	if demo == nil {
		return nil, fmt.Errorf("demo provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(demo)
	if err != nil {
		return nil, err
	}
	name := demo.Name
	if name == nil {
		return nil, fmt.Errorf("demo.Name must be provided to Apply")
	}
	result = &v1.Demo{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("demos").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *demos) ApplyStatus(ctx context.Context, demo *demov1.DemoApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Demo, err error) {
	if demo == nil {
		return nil, fmt.Errorf("demo provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(demo)
	if err != nil {
		return nil, err
	}

	name := demo.Name
	if name == nil {
		return nil, fmt.Errorf("demo.Name must be provided to Apply")
	}

	result = &v1.Demo{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("demos").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
