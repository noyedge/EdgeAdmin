// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ddosProtection

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/default/nodes/nodeutils"
	"github.com/noyedge/EdgeCommon/pkg/messageconfigs"
)

type StatusAction struct {
	actionutils.ParentAction
}

func (this *StatusAction) RunPost(params struct {
	NodeId int64
}) {
	results, err := nodeutils.SendMessageToNodeIds(this.AdminContext(), []int64{params.NodeId}, messageconfigs.MessageCodeCheckLocalFirewall, &messageconfigs.CheckLocalFirewallMessage{
		Name: "nftables",
	}, 10)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["results"] = results
	this.Success()
}
