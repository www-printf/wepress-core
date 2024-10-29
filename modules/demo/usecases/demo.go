package usecases

import (
	"context"

	"github.com/www-printf/wepress-core/modules/demo/domains"
)

type DemoUsecase interface {
	Get(ctx context.Context) (*domains.Demo, error)
}

type demoUsecase struct{}

func NewDemoUsecase() DemoUsecase {
	return &demoUsecase{}
}

func (u *demoUsecase) Get(ctx context.Context) (*domains.Demo, error) {
	demo := domains.Demo{
		Message: "Hello, World!",
	}
	return &demo, nil
}
