package ui

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeCommon)).
			Helper(settingutils.NewHelper("ui")).
			Prefix("/settings/ui").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
