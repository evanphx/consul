// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	GroupName      = "catalog"
	CurrentVersion = "v1alpha1"

	WorkloadKind         = "Workload"
	ServiceKind          = "Service"
	ServiceEndpointsKind = "ServiceEndpoints"
	VirtualIPsKind       = "VirtualIPs"
	HealthStatusKind     = "HealthStatus"
	HealthChecksKind     = "HealthChecks"
	NodeKind             = "Node"
	DNSPolicyKind        = "DNSPolicy"
)

var (
	WorkloadV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         WorkloadKind,
	}

	ServiceV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         ServiceKind,
	}

	ServiceEndpointsV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         ServiceEndpointsKind,
	}

	VirtualIPsV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         VirtualIPsKind,
	}

	HealthStatusV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         HealthStatusKind,
	}

	HealthChecksV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         HealthChecksKind,
	}

	NodeV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         NodeKind,
	}

	DNSPolicyV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         DNSPolicyKind,
	}
)

func Register(r resource.Registry) {
	RegisterWorkloadV1Alpha1Type(r)
	RegisterServiceV1Alpha1Type(r)
	RegisterServiceEndpointsV1Alpha1Type(r)
	RegisterNodeV1Alpha1Type(r)
	RegisterHealthStatusV1Alpha1Type(r)
	RegisterHealthChecksV1Alpha1Type(r)
	RegisterDNSPolicyV1Alpha1Type(r)
	RegisterVirtualIPsV1Alpha1Type(r)
}

func RegisterWorkloadV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     WorkloadV1Alpha1Type,
		Proto:    &pbcatalog.Workload{},
		Validate: nil,
	})
}

func RegisterServiceV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     ServiceV1Alpha1Type,
		Proto:    &pbcatalog.Service{},
		Validate: nil,
	})
}

func RegisterServiceEndpointsV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     ServiceEndpointsV1Alpha1Type,
		Proto:    &pbcatalog.ServiceEndpoints{},
		Validate: nil,
	})
}

func RegisterVirtualIPsV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     VirtualIPsV1Alpha1Type,
		Proto:    &pbcatalog.VirtualIPs{},
		Validate: nil,
	})
}

func RegisterHealthStatusV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     HealthStatusV1Alpha1Type,
		Proto:    &pbcatalog.HealthStatus{},
		Validate: nil,
	})
}

func RegisterHealthChecksV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     HealthChecksV1Alpha1Type,
		Proto:    &pbcatalog.HealthChecks{},
		Validate: nil,
	})
}

func RegisterNodeV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     NodeV1Alpha1Type,
		Proto:    &pbcatalog.Node{},
		Validate: nil,
	})
}

func RegisterDNSPolicyV1Alpha1Type(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     DNSPolicyV1Alpha1Type,
		Proto:    &pbcatalog.DNSPolicy{},
		Validate: nil,
	})
}
