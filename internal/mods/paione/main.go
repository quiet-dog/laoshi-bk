package paione

import (
	"context"

	"github.com/xxx/testapp/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/paione/api"
	"github.com/xxx/testapp/internal/mods/paione/schema"
	"gorm.io/gorm"
)

type PaiOne struct {
	DB        *gorm.DB
	PaiOneAPI *api.PaiOne
}

func (a *PaiOne) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.PaiOne))
}

func (a *PaiOne) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *PaiOne) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	v1 = v1.Group("paione")
	paiOne := v1.Group("pai-ones")
	{
		paiOne.GET("", a.PaiOneAPI.Query)
		paiOne.GET(":id", a.PaiOneAPI.Get)
		paiOne.POST("", a.PaiOneAPI.Create)
		paiOne.PUT(":id", a.PaiOneAPI.Update)
		paiOne.DELETE(":id", a.PaiOneAPI.Delete)
	}
	return nil
}

func (a *PaiOne) Release(ctx context.Context) error {
	return nil
}
