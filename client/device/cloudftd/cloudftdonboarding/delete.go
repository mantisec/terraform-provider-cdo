package cloudftdonboarding

import (
	"context"
	"github.com/mantisec/terraform-provider-cdo/go-client/internal/http"
)

type DeleteInput struct {
}

func NewDeleteInput() DeleteInput {
	return DeleteInput{}
}

type DeleteOutput struct {
}

func Delete(ctx context.Context, client http.Client, deleteInp DeleteInput) (*DeleteOutput, error) {
	return &DeleteOutput{}, nil
}
