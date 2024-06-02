package security

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeSetting)).
			Helper(settingutils.NewHelper("security")).
			Prefix("/settings/security").
			GetPost("", new(IndexAction)).
			Post("/dismissXFFPrompt", new(DismissXFFPromptAction)).
			EndAll()
	})
}
