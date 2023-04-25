package mesh

import (
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/resource/mesh/internal/types"
)

var (
	ProxyConfigurationV1Alpha1Type = types.ProxyConfigurationV1Alpha1Type
	UpstreamsV1Alpha1Type          = types.UpstreamsV1Alpha1Type
)

// RegisterTypes adds all resource types within the "catalog" API group
// to the given type registry
func RegisterTypes(r resource.Registry) {
	types.Register(r)
}
