package paione

import (
	"github.com/google/wire"
	"github.com/xxx/testapp/internal/mods/paione/api"
	"github.com/xxx/testapp/internal/mods/paione/biz"
	"github.com/xxx/testapp/internal/mods/paione/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(PaiOne), "*"),
	wire.Struct(new(dal.PaiOne), "*"),
	wire.Struct(new(biz.PaiOne), "*"),
	wire.Struct(new(api.PaiOne), "*"),
)
