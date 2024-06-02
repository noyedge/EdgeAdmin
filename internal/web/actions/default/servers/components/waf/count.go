package waf

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type CountAction struct {
	actionutils.ParentAction
}

func (this *CountAction) RunPost(params struct{}) {
	countResp, err := this.RPC().HTTPFirewallPolicyRPC().CountAllEnabledHTTPFirewallPolicies(this.AdminContext(), &pb.CountAllEnabledHTTPFirewallPoliciesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["count"] = countResp.Count

	this.Success()
}
