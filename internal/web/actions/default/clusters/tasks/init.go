package tasks

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/clusters/clusterutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNode)).
			Helper(clusterutils.NewClustersHelper()).
			Prefix("/clusters/tasks").
			GetPost("/listPopup", new(ListPopupAction)).
			Post("/check", new(CheckAction)).
			Post("/delete", new(DeleteAction)).
			Post("/deleteBatch", new(DeleteBatchAction)).
			Post("/deleteAll", new(DeleteAllAction)).

			EndAll()
	})
}
