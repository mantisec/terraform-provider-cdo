package duoadminpanel

import (
	"context"
	"github.com/mantisec/terraform-provider-cdo/go-client/device"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/http"
)

type DeleteInput struct {
	Uid string
}

type DeleteOutput = device.DeleteOutput

func Delete(ctx context.Context, client http.Client, deleteInp DeleteInput) (*DeleteOutput, error) {

	client.Logger.Println("deleting duo admin panel")

	return device.Delete(ctx, client, *device.NewDeleteInput(deleteInp.Uid))
}
