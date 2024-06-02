package waf

import (	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type UpdateGroupOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateGroupOnAction) RunPost(params struct {
	GroupId int64
	IsOn    bool
}) {
	// 日志
	defer this.CreateLogInfo(codes.WAFRuleGroup_LogUpdateRuleGroupIsOn, params.GroupId)

	_, err := this.RPC().HTTPFirewallRuleGroupRPC().UpdateHTTPFirewallRuleGroupIsOn(this.AdminContext(), &pb.UpdateHTTPFirewallRuleGroupIsOnRequest{
		FirewallRuleGroupId: params.GroupId,
		IsOn:                params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
