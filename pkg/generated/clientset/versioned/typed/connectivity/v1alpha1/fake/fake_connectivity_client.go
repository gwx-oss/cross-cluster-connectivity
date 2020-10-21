// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vmware-tanzu/cross-cluster-connectivity/pkg/generated/clientset/versioned/typed/connectivity/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConnectivityV1alpha1 struct {
	*testing.Fake
}

func (c *FakeConnectivityV1alpha1) RemoteRegistries(namespace string) v1alpha1.RemoteRegistryInterface {
	return &FakeRemoteRegistries{c, namespace}
}

func (c *FakeConnectivityV1alpha1) ServiceRecords(namespace string) v1alpha1.ServiceRecordInterface {
	return &FakeServiceRecords{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConnectivityV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}