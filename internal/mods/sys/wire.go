package sys

import (
	"github.com/google/wire"
	"github.com/xxx/testapp/internal/mods/sys/api"
	"github.com/xxx/testapp/internal/mods/sys/biz"
	"github.com/xxx/testapp/internal/mods/sys/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(SYS), "*"),
	wire.Struct(new(dal.Logger), "*"),
	wire.Struct(new(biz.Logger), "*"),
	wire.Struct(new(api.Logger), "*"),
	wire.Struct(new(dal.Dictionary), "*"),
	wire.Struct(new(biz.Dictionary), "*"),
	wire.Struct(new(api.Dictionary), "*"),
)
