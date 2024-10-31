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
	DB         *gorm.DB
	SignAPI    *api.Sign
	ActiveAPI  *api.Active
	PkAPI      *api.Pk
	EmployAPI  *api.Employ
	TaoLunAPI  *api.TaoLun
	FileAPI    *api.File
	SignLogAPI *api.SignLog
	CommentAPI *api.Comment
	PkLogAPI   *api.PkLog
	PkScoreAPI *api.PkScore
	ClassAPI   *api.Class
}

func (a *Class) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Sign), new(schema.Active), new(schema.Pk), new(schema.Employ), new(schema.TaoLun), new(schema.File), new(schema.SignLog), new(schema.Comment), new(schema.PkLog), new(schema.PkScore), new(schema.Class))
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
		// employ.GET("preview", a.EmployAPI.Preview)
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
	signLog := v1.Group("sign-logs")
	{
		signLog.GET("", a.SignLogAPI.Query)
		signLog.GET(":id", a.SignLogAPI.Get)
		signLog.POST("", a.SignLogAPI.Create)
		signLog.PUT(":id", a.SignLogAPI.Update)
		signLog.DELETE(":id", a.SignLogAPI.Delete)
	}
	comment := v1.Group("comments")
	{
		comment.GET("", a.CommentAPI.Query)
		comment.GET(":id", a.CommentAPI.Get)
		comment.POST("", a.CommentAPI.Create)
		comment.PUT(":id", a.CommentAPI.Update)
		comment.DELETE(":id", a.CommentAPI.Delete)
	}
	pkLog := v1.Group("pk-logs")
	{
		pkLog.GET("", a.PkLogAPI.Query)
		pkLog.GET(":id", a.PkLogAPI.Get)
		pkLog.POST("", a.PkLogAPI.Create)
		pkLog.PUT(":id", a.PkLogAPI.Update)
		pkLog.DELETE(":id", a.PkLogAPI.Delete)
	}
	pkScore := v1.Group("pk-scores")
	{
		pkScore.GET("", a.PkScoreAPI.Query)
		pkScore.GET(":id", a.PkScoreAPI.Get)
		pkScore.POST("", a.PkScoreAPI.Create)
		pkScore.PUT(":id", a.PkScoreAPI.Update)
		pkScore.DELETE(":id", a.PkScoreAPI.Delete)
	}
	class := v1.Group("classes")
	{
		class.GET("", a.ClassAPI.Query)
		class.GET(":id", a.ClassAPI.Get)
		class.POST("", a.ClassAPI.Create)
		class.PUT(":id", a.ClassAPI.Update)
		class.DELETE(":id", a.ClassAPI.Delete)
		class.GET("tree", a.ClassAPI.Tree)
	}

	return nil
}

func (a *Class) Release(ctx context.Context) error {
	return nil
}
