package servers

import "github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdateAction) RunGet(params struct{}) {
	this.Show()
}
