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
)
