// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplists

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type ReadAllAction struct {
	actionutils.ParentAction
}

func (this *ReadAllAction) RunPost(params struct{}) {
	defer this.CreateLogInfo(codes.IPItem_LogReadAllIPItems)

	_, err := this.RPC().IPItemRPC().UpdateIPItemsRead(this.AdminContext(), &pb.UpdateIPItemsReadRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧菜单Badge更新
	helpers.NotifyIPItemsCountChanges()

	this.Success()
}
