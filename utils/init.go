package utils

import "context"

type UtilsModule struct{}

func NewUtilsModule(ctx context.Context) *UtilsModule {
	return &UtilsModule{}
}
