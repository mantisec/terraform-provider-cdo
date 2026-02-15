package connectoronboarding

import (
	"context"
	"github.com/mantisec/terraform-provider-cdo/go-client/connector"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/http"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/retry"
	"github.com/mantisec/terraform-provider-cdo/go-client/model/device/status"
)

func UntilConnectorStatusIsActive(ctx context.Context, client http.Client, readInp connector.ReadByNameInput, readOutp *connector.ReadOutput) retry.Func {
	return func() (bool, error) {
		readRes, err := connector.ReadByName(ctx, client, readInp)
		*readOutp = *readRes
		if err != nil {
			return false, err
		}
		client.Logger.Printf("connector status: %v\n", readRes.ConnectorStatus)
		if readRes.ConnectorStatus == status.Active {
			return true, nil
		}
		return false, nil
	}
}
