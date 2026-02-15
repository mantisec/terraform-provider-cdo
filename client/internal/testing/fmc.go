package testing

import (
	"github.com/mantisec/terraform-provider-cdo/go-client/device/cloudfmc"
	"github.com/mantisec/terraform-provider-cdo/go-client/device/cloudfmc/fmcplatform"
	"github.com/mantisec/terraform-provider-cdo/go-client/model/cloudfmc/accesspolicies"
	"github.com/mantisec/terraform-provider-cdo/go-client/model/cloudfmc/fmcdomain"
)

func (m Model) FmcReadOutput() cloudfmc.ReadOutput {
	return cloudfmc.ReadOutput{
		Host: m.FmcHost,
	}
}

func (m Model) FmcDomainInfo() fmcplatform.ReadDomainInfoOutput {
	return fmcplatform.ReadDomainInfoOutput{
		Items: []fmcdomain.Item{
			{
				Uuid: m.FmcDomainUuid.String(),
				Name: "",
				Type: "",
			},
		},
	}
}

func (m Model) ReadAccessPolicies() accesspolicies.AccessPolicies {
	return accesspolicies.AccessPolicies{
		Items: []accesspolicies.Item{
			{
				Id:   m.FtdAccessPolicyUid.String(),
				Name: m.FtdAccessPolicyName,
				Type: "",
			},
		},
	}
}
