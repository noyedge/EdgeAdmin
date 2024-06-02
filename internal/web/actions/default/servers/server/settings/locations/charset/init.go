package charset

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/locationutils"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/servers/serverutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(locationutils.NewLocationHelper()).
			Helper(serverutils.NewServerHelper()).
			Data("tinyMenuItem", "charset").
			Prefix("/servers/server/settings/locations/charset").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
