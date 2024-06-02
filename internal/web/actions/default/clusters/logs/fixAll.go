// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package logs

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type FixAllAction struct {
	actionutils.ParentAction
}

func (this *FixAllAction) RunPost(params struct {
}) {
	defer this.CreateLogInfo(codes.NodeLog_LogFixAllLogs)

	_, err := this.RPC().NodeLogRPC().FixAllNodeLogs(this.AdminContext(), &pb.FixAllNodeLogsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧数字Badge更新
	helpers.NotifyNodeLogsCountChange()

	this.Success()
}
