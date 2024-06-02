package locations

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/servers/serverutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(serverutils.NewServerHelper()).
			Prefix("/servers/server/settings/locations").
			Get("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Post("/delete", new(DeleteAction)).
			Post("/sort", new(SortAction)).
			EndAll()
	})
}
