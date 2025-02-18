// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	ciliumiov2alpha1 "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/typed/cilium.io/v2alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCiliumNodeConfigs implements CiliumNodeConfigInterface
type fakeCiliumNodeConfigs struct {
	*gentype.FakeClientWithList[*v2alpha1.CiliumNodeConfig, *v2alpha1.CiliumNodeConfigList]
	Fake *FakeCiliumV2alpha1
}

func newFakeCiliumNodeConfigs(fake *FakeCiliumV2alpha1, namespace string) ciliumiov2alpha1.CiliumNodeConfigInterface {
	return &fakeCiliumNodeConfigs{
		gentype.NewFakeClientWithList[*v2alpha1.CiliumNodeConfig, *v2alpha1.CiliumNodeConfigList](
			fake.Fake,
			namespace,
			v2alpha1.SchemeGroupVersion.WithResource("ciliumnodeconfigs"),
			v2alpha1.SchemeGroupVersion.WithKind("CiliumNodeConfig"),
			func() *v2alpha1.CiliumNodeConfig { return &v2alpha1.CiliumNodeConfig{} },
			func() *v2alpha1.CiliumNodeConfigList { return &v2alpha1.CiliumNodeConfigList{} },
			func(dst, src *v2alpha1.CiliumNodeConfigList) { dst.ListMeta = src.ListMeta },
			func(list *v2alpha1.CiliumNodeConfigList) []*v2alpha1.CiliumNodeConfig {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v2alpha1.CiliumNodeConfigList, items []*v2alpha1.CiliumNodeConfig) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
