package class

import (
	"github.com/google/wire"
	"github.com/xxx/testapp/internal/mods/class/api"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(Class), "*"),
	wire.Struct(new(dal.Sign), "*"),
	wire.Struct(new(biz.Sign), "*"),
	wire.Struct(new(api.Sign), "*"),
	wire.Struct(new(dal.Active), "*"),
	wire.Struct(new(biz.Active), "*"),
	wire.Struct(new(api.Active), "*"),
	wire.Struct(new(dal.Pk), "*"),
	wire.Struct(new(biz.Pk), "*"),
	wire.Struct(new(api.Pk), "*"),
	wire.Struct(new(dal.Employ), "*"),
	wire.Struct(new(biz.Employ), "*"),
	wire.Struct(new(api.Employ), "*"),
	wire.Struct(new(dal.TaoLun), "*"),
	wire.Struct(new(biz.TaoLun), "*"),
	wire.Struct(new(api.TaoLun), "*"),
	wire.Struct(new(dal.File), "*"),
	wire.Struct(new(biz.File), "*"),
	wire.Struct(new(api.File), "*"),
	wire.Struct(new(dal.SignLog), "*"),
	wire.Struct(new(biz.SignLog), "*"),
	wire.Struct(new(api.SignLog), "*"),
	wire.Struct(new(dal.Comment), "*"),
	wire.Struct(new(biz.Comment), "*"),
	wire.Struct(new(api.Comment), "*"),
	wire.Struct(new(dal.PkLog), "*"),
	wire.Struct(new(biz.PkLog), "*"),
	wire.Struct(new(api.PkLog), "*"),
	wire.Struct(new(dal.PkScore), "*"),
	wire.Struct(new(biz.PkScore), "*"),
	wire.Struct(new(api.PkScore), "*"),
	wire.Struct(new(dal.Class), "*"),
	wire.Struct(new(biz.Class), "*"),
	wire.Struct(new(api.Class), "*"),
	wire.Struct(new(dal.PaiOne), "*"),
	wire.Struct(new(biz.PaiOne), "*"),
	wire.Struct(new(api.PaiOne), "*"),
	wire.Struct(new(dal.PaiTwo), "*"),
	wire.Struct(new(biz.PaiTwo), "*"),
	wire.Struct(new(api.PaiTwo), "*"),
	wire.Struct(new(dal.PaiThree), "*"),
	wire.Struct(new(biz.PaiThree), "*"),
	wire.Struct(new(api.PaiThree), "*"),
	wire.Struct(new(dal.PaiFour), "*"),
	wire.Struct(new(biz.PaiFour), "*"),
	wire.Struct(new(api.PaiFour), "*"),
)
