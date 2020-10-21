// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/vmware-tanzu/cross-cluster-connectivity/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// RemoteRegistries returns a RemoteRegistryInformer.
	RemoteRegistries() RemoteRegistryInformer
	// ServiceRecords returns a ServiceRecordInformer.
	ServiceRecords() ServiceRecordInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// RemoteRegistries returns a RemoteRegistryInformer.
func (v *version) RemoteRegistries() RemoteRegistryInformer {
	return &remoteRegistryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceRecords returns a ServiceRecordInformer.
func (v *version) ServiceRecords() ServiceRecordInformer {
	return &serviceRecordInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}