package dashboard

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.Prefix("/dashboard").
			Data("teaMenu", "dashboard").
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeCommon)).
			GetPost("", new(IndexAction)).
			Post("/restartLocalAPINode", new(RestartLocalAPINodeAction)).
			EndAll()
	})
}
