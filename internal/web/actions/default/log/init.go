package log

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeLog)).
			Helper(new(Helper)).
			Prefix("/log").
			Get("", new(IndexAction)).
			Get("/exportExcel", new(ExportExcelAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/clean", new(CleanAction)).
			GetPost("/settings", new(SettingsAction)).

			EndAll()
	})
}
