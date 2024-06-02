// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplists

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeAdmin/internal/web/helpers"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/types"
	"strings"
)

type DeleteItemsAction struct {
	actionutils.ParentAction
}

func (this *DeleteItemsAction) RunPost(params struct {
	ItemIds []int64
}) {
	if len(params.ItemIds) == 0 {
		this.Success()
	}

	var itemIdStrings = []string{}
	for _, itemId := range params.ItemIds {
		itemIdStrings = append(itemIdStrings, types.String(itemId))
	}

	defer this.CreateLogInfo(codes.IPList_LogDeleteIPBatch, strings.Join(itemIdStrings, ", "))

	_, err := this.RPC().IPItemRPC().DeleteIPItems(this.AdminContext(), &pb.DeleteIPItemsRequest{IpItemIds: params.ItemIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧菜单Badge更新
	helpers.NotifyIPItemsCountChanges()

	this.Success()
}
