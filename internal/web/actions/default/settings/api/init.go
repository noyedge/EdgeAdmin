package api

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/settings/api/node"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeSetting)).
			Helper(NewHelper()).
			Helper(settingutils.NewAdvancedHelper("apiNodes")).
			Prefix("/settings/api").
			Get("", new(IndexAction)).
			Get("/methodStats", new(MethodStatsAction)).
			GetPost("/node/createPopup", new(node.CreatePopupAction)).
			Post("/delete", new(DeleteAction)).
			EndAll()
	})
}
