package testing

import (
	"fmt"

	"github.com/mantisec/terraform-provider-cdo/go-client/device/ios"
	"github.com/mantisec/terraform-provider-cdo/go-client/model/device/publicapilabels"
	"github.com/mantisec/terraform-provider-cdo/go-client/model/device/tags"
	"github.com/mantisec/terraform-provider-cdo/go-client/model/devicetype"
)

func (m Model) CreateIosInput() ios.CreateInput {
	return ios.CreateInput{
		Name:              m.IosName,
		ConnectorUid:      m.CdgUid.String(),
		ConnectorType:     "CDG",
		SocketAddress:     m.IosHost,
		Labels:            publicapilabels.Type{},
		Username:          m.IosUsername,
		Password:          m.IosPassword,
		IgnoreCertificate: false,
	}
}

func (m Model) ReadIosOutput() ios.ReadOutput {
	return ios.ReadOutput{
		Uid:               m.IosUid.String(),
		Name:              m.IosName,
		CreatedDate:       0,
		LastUpdatedDate:   0,
		DeviceType:        devicetype.Ios,
		ConnectorUid:      m.SdcUid.String(),
		ConnectorType:     "SDC",
		SocketAddress:     fmt.Sprintf("%s:%s", m.IosHost, m.IosPort),
		Port:              m.IosPort,
		Host:              m.IosHost,
		Tags:              tags.Type{},
		IgnoreCertificate: false,
	}
}
