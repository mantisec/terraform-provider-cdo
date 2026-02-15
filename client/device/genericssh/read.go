package genericssh

import (
	"context"
	"github.com/mantisec/terraform-provider-cdo/go-client/device"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/http"
)

type ReadInput struct {
	Uid string
}

type ReadOutput = device.ReadOutput

func NewReadInput(uid string) *ReadInput {
	return &ReadInput{
		Uid: uid,
	}
}

func Read(ctx context.Context, client http.Client, readInp ReadInput) (*ReadOutput, error) {

	client.Logger.Println("reading generic ssh")

	readOutp, err := device.ReadByUid(ctx, client, *device.NewReadByUidInput(readInp.Uid))
	if err != nil {
		return nil, err
	}

	return readOutp, nil
}
