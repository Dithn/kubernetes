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

package v1

import (
	context "context"
	fmt "fmt"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	apply "k8s.io/client-go/util/apply"
	extensionsv1 "k8s.io/code-generator/examples/crd/apis/extensions/v1"
	applyconfigurationextensionsv1 "k8s.io/code-generator/examples/crd/applyconfiguration/extensions/v1"
	scheme "k8s.io/code-generator/examples/crd/clientset/versioned/scheme"
)

// TestTypesGetter has a method to return a TestTypeInterface.
// A group's client should implement this interface.
type TestTypesGetter interface {
	TestTypes(namespace string) TestTypeInterface
}

// TestTypeInterface has methods to work with TestType resources.
type TestTypeInterface interface {
	Create(ctx context.Context, testType *extensionsv1.TestType, opts metav1.CreateOptions) (*extensionsv1.TestType, error)
	Update(ctx context.Context, testType *extensionsv1.TestType, opts metav1.UpdateOptions) (*extensionsv1.TestType, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, testType *extensionsv1.TestType, opts metav1.UpdateOptions) (*extensionsv1.TestType, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*extensionsv1.TestType, error)
	List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1.TestTypeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *extensionsv1.TestType, err error)
	Apply(ctx context.Context, testType *applyconfigurationextensionsv1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *extensionsv1.TestType, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, testType *applyconfigurationextensionsv1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *extensionsv1.TestType, err error)
	GetExtended(ctx context.Context, name string, opts metav1.GetOptions) (*extensionsv1.TestType, error)
	ListExtended(ctx context.Context, opts metav1.ListOptions) (*extensionsv1.TestTypeList, error)
	CreateExtended(ctx context.Context, testType *extensionsv1.TestType, opts metav1.CreateOptions) (*extensionsv1.TestType, error)
	UpdateExtended(ctx context.Context, testType *extensionsv1.TestType, opts metav1.UpdateOptions) (*extensionsv1.TestType, error)
	PatchExtended(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *extensionsv1.TestType, err error)
	ApplyExtended(ctx context.Context, testType *applyconfigurationextensionsv1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *extensionsv1.TestType, err error)
	GetSubresource(ctx context.Context, testTypeName string, options metav1.GetOptions) (*extensionsv1.TestSubresource, error)
	CreateSubresource(ctx context.Context, testTypeName string, testSubresource *extensionsv1.TestSubresource, opts metav1.CreateOptions) (*extensionsv1.TestSubresource, error)
	UpdateSubresource(ctx context.Context, testTypeName string, testSubresource *extensionsv1.TestSubresource, opts metav1.UpdateOptions) (*extensionsv1.TestSubresource, error)
	ApplySubresource(ctx context.Context, testTypeName string, testSubresource *applyconfigurationextensionsv1.TestSubresourceApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1.TestSubresource, error)

	TestTypeExpansion
}

// testTypes implements TestTypeInterface
type testTypes struct {
	*gentype.ClientWithListAndApply[*extensionsv1.TestType, *extensionsv1.TestTypeList, *applyconfigurationextensionsv1.TestTypeApplyConfiguration]
}

// newTestTypes returns a TestTypes
func newTestTypes(c *ExtensionsExampleV1Client, namespace string) *testTypes {
	return &testTypes{
		gentype.NewClientWithListAndApply[*extensionsv1.TestType, *extensionsv1.TestTypeList, *applyconfigurationextensionsv1.TestTypeApplyConfiguration](
			"testtypes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *extensionsv1.TestType { return &extensionsv1.TestType{} },
			func() *extensionsv1.TestTypeList { return &extensionsv1.TestTypeList{} },
		),
	}
}

// GetExtended takes name of the testType, and returns the corresponding testType object, and an error if there is any.
func (c *testTypes) GetExtended(ctx context.Context, name string, options metav1.GetOptions) (result *extensionsv1.TestType, err error) {
	result = &extensionsv1.TestType{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// ListExtended takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *testTypes) ListExtended(ctx context.Context, opts metav1.ListOptions) (result *extensionsv1.TestTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &extensionsv1.TestTypeList{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// CreateExtended takes the representation of a testType and creates it.  Returns the server's representation of the testType, and an error, if there is any.
func (c *testTypes) CreateExtended(ctx context.Context, testType *extensionsv1.TestType, opts metav1.CreateOptions) (result *extensionsv1.TestType, err error) {
	result = &extensionsv1.TestType{}
	err = c.GetClient().Post().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(testType).
		Do(ctx).
		Into(result)
	return
}

// UpdateExtended takes the representation of a testType and updates it. Returns the server's representation of the testType, and an error, if there is any.
func (c *testTypes) UpdateExtended(ctx context.Context, testType *extensionsv1.TestType, opts metav1.UpdateOptions) (result *extensionsv1.TestType, err error) {
	result = &extensionsv1.TestType{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testType.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(testType).
		Do(ctx).
		Into(result)
	return
}

// PatchExtended applies the patch and returns the patched testType.
func (c *testTypes) PatchExtended(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *extensionsv1.TestType, err error) {
	result = &extensionsv1.TestType{}
	err = c.GetClient().Patch(pt).
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyExtended takes the given apply declarative configuration, applies it and returns the applied testType.
func (c *testTypes) ApplyExtended(ctx context.Context, testType *applyconfigurationextensionsv1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *extensionsv1.TestType, err error) {
	if testType == nil {
		return nil, fmt.Errorf("testType provided to ApplyExtended must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	name := testType.Name
	if name == nil {
		return nil, fmt.Errorf("testType.Name must be provided to ApplyExtended")
	}
	request, err := apply.NewRequest(c.GetClient(), testType)
	if err != nil {
		return nil, err
	}
	result = &extensionsv1.TestType{}
	err = request.
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// GetSubresource takes name of the testType, and returns the corresponding extensionsv1.TestSubresource object, and an error if there is any.
func (c *testTypes) GetSubresource(ctx context.Context, testTypeName string, options metav1.GetOptions) (result *extensionsv1.TestSubresource, err error) {
	result = &extensionsv1.TestSubresource{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testTypeName).
		SubResource("testsubresource").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// CreateSubresource takes the representation of a testSubresource and creates it.  Returns the server's representation of the testSubresource, and an error, if there is any.
func (c *testTypes) CreateSubresource(ctx context.Context, testTypeName string, testSubresource *extensionsv1.TestSubresource, opts metav1.CreateOptions) (result *extensionsv1.TestSubresource, err error) {
	result = &extensionsv1.TestSubresource{}
	err = c.GetClient().Post().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testTypeName).
		SubResource("testsubresource").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(testSubresource).
		Do(ctx).
		Into(result)
	return
}

// UpdateSubresource takes the top resource name and the representation of a testSubresource and updates it. Returns the server's representation of the testSubresource, and an error, if there is any.
func (c *testTypes) UpdateSubresource(ctx context.Context, testTypeName string, testSubresource *extensionsv1.TestSubresource, opts metav1.UpdateOptions) (result *extensionsv1.TestSubresource, err error) {
	result = &extensionsv1.TestSubresource{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testTypeName).
		SubResource("subresource").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(testSubresource).
		Do(ctx).
		Into(result)
	return
}

// ApplySubresource takes top resource name and the apply declarative configuration for subresource,
// applies it and returns the applied testSubresource, and an error, if there is any.
func (c *testTypes) ApplySubresource(ctx context.Context, testTypeName string, testSubresource *applyconfigurationextensionsv1.TestSubresourceApplyConfiguration, opts metav1.ApplyOptions) (result *extensionsv1.TestSubresource, err error) {
	if testSubresource == nil {
		return nil, fmt.Errorf("testSubresource provided to ApplySubresource must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	request, err := apply.NewRequest(c.GetClient(), testSubresource)
	if err != nil {
		return nil, err
	}

	result = &extensionsv1.TestSubresource{}
	err = request.
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testTypeName).
		SubResource("subresource").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
