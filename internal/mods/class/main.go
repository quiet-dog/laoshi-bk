package class

import (
	"context"

	"github.com/xxx/testapp/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/api"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"gorm.io/gorm"
)

type Class struct {
	DB        *gorm.DB
	SignAPI   *api.Sign
	ActiveAPI *api.Active
	PkAPI     *api.Pk
	EmployAPI *api.Employ
	TaoLunAPI *api.TaoLun
	FileAPI   *api.File
}

func (a *Class) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Sign), new(schema.Active), new(schema.Pk), new(schema.Employ), new(schema.TaoLun), new(schema.File))
}

func (a *Class) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Class) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	v1 = v1.Group("class")
	sign := v1.Group("signs")
	{
		sign.GET("", a.SignAPI.Query)
		sign.GET(":id", a.SignAPI.Get)
		sign.POST("", a.SignAPI.Create)
		sign.PUT(":id", a.SignAPI.Update)
		sign.DELETE(":id", a.SignAPI.Delete)
	}
	active := v1.Group("actives")
	{
		active.GET("", a.ActiveAPI.Query)
		active.GET(":id", a.ActiveAPI.Get)
		active.POST("", a.ActiveAPI.Create)
		active.PUT(":id", a.ActiveAPI.Update)
		active.DELETE(":id", a.ActiveAPI.Delete)
	}
	pk := v1.Group("pks")
	{
		pk.GET("", a.PkAPI.Query)
		pk.GET(":id", a.PkAPI.Get)
		pk.POST("", a.PkAPI.Create)
		pk.PUT(":id", a.PkAPI.Update)
		pk.DELETE(":id", a.PkAPI.Delete)
	}
	employ := v1.Group("employs")
	{
		employ.GET("", a.EmployAPI.Query)
		employ.GET(":id", a.EmployAPI.Get)
		employ.POST("", a.EmployAPI.Create)
		employ.PUT(":id", a.EmployAPI.Update)
		employ.DELETE(":id", a.EmployAPI.Delete)
	}
	taoLun := v1.Group("tao-luns")
	{
		taoLun.GET("", a.TaoLunAPI.Query)
		taoLun.GET(":id", a.TaoLunAPI.Get)
		taoLun.POST("", a.TaoLunAPI.Create)
		taoLun.PUT(":id", a.TaoLunAPI.Update)
		taoLun.DELETE(":id", a.TaoLunAPI.Delete)
	}
	file := v1.Group("files")
	{
		file.GET("", a.FileAPI.Query)
		file.GET(":id", a.FileAPI.Get)
		file.POST("", a.FileAPI.Create)
		file.PUT(":id", a.FileAPI.Update)
		file.DELETE(":id", a.FileAPI.Delete)
	}

	return nil
}

func (a *Class) Release(ctx context.Context) error {
	return nil
}
