// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"github.com/hashicorp/consul/internal/resource"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	GroupName      = "mesh"
	CurrentVersion = "v1alpha1"

	ProxyConfigurationKind = "ProxyConfiguration"
	UpstreamsKind          = "Upstreams"
)

var (
	ProxyConfigurationV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         ProxyConfigurationKind,
	}

	UpstreamsV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         UpstreamsKind,
	}
)

func Register(r resource.Registry) {
	registerProxyConfigurationV1Alpha1Type(r)
	registerUpstreamsV1Alpha1Type(r)
}

func registerProxyConfigurationV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     ProxyConfigurationV1Alpha1Type,
		Proto:    &pbmesh.ProxyConfiguration{},
		Validate: nil,
	})
}

func registerUpstreamsV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     UpstreamsV1Alpha1Type,
		Proto:    &pbmesh.Upstreams{},
		Validate: nil,
	})
}
