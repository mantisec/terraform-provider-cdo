package tenants

import (
	"context"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/http"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/url"
)

func DeleteByUid(ctx context.Context, client http.Client, deleteInp DeleteByUidInput) (interface{}, error) {
	client.Logger.Println("Removing tenant by UID from the MSP portal " + deleteInp.Uid)
	deleteUrl := url.MspManagedTenantByUid(client.BaseUrl(), deleteInp.Uid)
	if err := client.NewDelete(ctx, deleteUrl).Send(&DeleteOutput{}); err != nil {
		return nil, err
	}

	return nil, nil
}
