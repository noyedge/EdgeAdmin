package nodes

import (
	"github.com/iwind/TeaGo"
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/nodes/ipAddresses"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNode)).
			Helper(new(Helper)).
			Prefix("/nodes").
			Post("/delete", new(DeleteAction)).

			// IP地址
			GetPost("/ipAddresses/createPopup", new(ipAddresses.CreatePopupAction)).
			GetPost("/ipAddresses/updatePopup", new(ipAddresses.UpdatePopupAction)).

			// 统计
			Post("/stat", new(StatAction)).
			EndAll()
	})
}
