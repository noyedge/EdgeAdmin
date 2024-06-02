package nodes

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type StatAction struct {
	actionutils.ParentAction
}

func (this *StatAction) RunPost(params struct {
	NodeId int64
}) {
	_, err := this.RPC().NodeRPC().ComposeServerStatNodeBoard(this.AdminContext(), &pb.ComposeServerStatNodeBoardRequest{
		NodeId: params.NodeId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
