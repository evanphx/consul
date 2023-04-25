package catalog

import (
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/resource/catalog/internal/types"
)

var (
	WorkloadV1Alpha1Type         = types.WorkloadV1Alpha1Type
	ServiceV1Alpha1Type          = types.ServiceV1Alpha1Type
	ServiceEndpointsV1Alpha1Type = types.ServiceEndpointsV1Alpha1Type
	VirtualIPsV1Alpha1Type       = types.VirtualIPsV1Alpha1Type
	NodeV1Alpha1Type             = types.NodeV1Alpha1Type
	HealthStatusV1Alpha1Type     = types.HealthStatusV1Alpha1Type
	HealthChecksV1Alpha1Type     = types.HealthChecksV1Alpha1Type
	DNSPolicyV1Alpha1Type        = types.DNSPolicyV1Alpha1Type
)

// RegisterTypes adds all resource types within the "catalog" API group
// to the given type registry
func RegisterTypes(r resource.Registry) {
	types.Register(r)
}
