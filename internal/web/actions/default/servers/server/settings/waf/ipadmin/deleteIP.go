package ipadmin

import (	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type DeleteIPAction struct {
	actionutils.ParentAction
}

func (this *DeleteIPAction) RunPost(params struct {
	FirewallPolicyId int64
	ItemId           int64
}) {
	// 日志
	defer this.CreateLogInfo(codes.WAF_LogDeleteIPFromWAFPolicy, params.FirewallPolicyId, params.ItemId)

	// TODO 判断权限

	_, err := this.RPC().IPItemRPC().DeleteIPItem(this.AdminContext(), &pb.DeleteIPItemRequest{IpItemId: params.ItemId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
