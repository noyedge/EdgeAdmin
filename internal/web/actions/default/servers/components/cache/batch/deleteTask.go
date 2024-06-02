// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package cache

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type DeleteTaskAction struct {
	actionutils.ParentAction
}

func (this *DeleteTaskAction) RunPost(params struct {
	TaskId int64
}) {
	defer this.CreateLogInfo(codes.HTTPCacheTask_LogDeleteHTTPCacheTask, params.TaskId)

	_, err := this.RPC().HTTPCacheTaskRPC().DeleteHTTPCacheTask(this.AdminContext(), &pb.DeleteHTTPCacheTaskRequest{
		HttpCacheTaskId: params.TaskId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
