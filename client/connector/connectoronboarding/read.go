package connectoronboarding

import (
	"context"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/http"
)

type ReadInput struct {
}

func NewReadInput() ReadInput {
	return ReadInput{}
}

type ReadOutput struct {
}

func Read(ctx context.Context, client http.Client, readInp ReadInput) (*ReadOutput, error) {

	// empty

	return nil, nil
}
