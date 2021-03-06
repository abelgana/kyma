// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/idppreset/pkg/client/clientset/versioned/typed/ui/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeUiV1alpha1 struct {
	*testing.Fake
}

func (c *FakeUiV1alpha1) IDPPresets() v1alpha1.IDPPresetInterface {
	return &FakeIDPPresets{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeUiV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
