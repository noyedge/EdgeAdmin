package database

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
			Helper(settingutils.NewAdvancedHelper("database")).
			Prefix("/settings/database").
			Get("", new(IndexAction)).
			GetPost("/update", new(UpdateAction)).
			GetPost("/clean", new(CleanAction)).
			GetPost("/cleanSetting", new(CleanSettingAction)).
			GetPost("/truncateTable", new(TruncateTableAction)).
			GetPost("/deleteTable", new(DeleteTableAction)).
			EndAll()
	})
}
