package examples

import (
	"context"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/http"
)

type ReadInput struct {
}

type ReadOutput struct {
}

func Read(ctx context.Context, client http.Client, readInp ReadInput) (*ReadOutput, error) {

	// TODO

	return nil, nil
}
