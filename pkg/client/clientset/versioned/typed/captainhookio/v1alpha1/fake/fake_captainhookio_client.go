/*
 Generated Code
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/garethjevans/captain-hook/pkg/client/clientset/versioned/typed/captainhookio/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCaptainhookV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCaptainhookV1alpha1) Hooks(namespace string) v1alpha1.HookInterface {
	return &FakeHooks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCaptainhookV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
