package student

import (
	"context"

	"github.com/xxx/testapp/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Student struct {
	DB *gorm.DB
}

func (a *Student) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate()
}

func (a *Student) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Student) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	v1 = v1.Group("student")

	return nil
}

func (a *Student) Release(ctx context.Context) error {
	return nil
}
