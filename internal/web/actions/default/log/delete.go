package log

import (
	"github.com/noyedge/EdgeAdmin/internal/configloaders"
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	LogId int64
}) {
	// 记录日志
	defer this.CreateLogInfo(codes.Log_LogDeleteLog, params.LogId)

	// 读取配置
	config, err := configloaders.LoadLogConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if !config.CanDelete {
		this.Fail("已设置不能删除")
	}

	// 执行删除
	_, err = this.RPC().LogRPC().DeleteLogPermanently(this.AdminContext(), &pb.DeleteLogPermanentlyRequest{LogId: params.LogId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
