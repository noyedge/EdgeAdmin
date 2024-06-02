package acme

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type RunAction struct {
	actionutils.ParentAction
}

func (this *RunAction) RunPost(params struct {
	TaskId int64
}) {
	defer this.CreateLogInfo(codes.ACMETask_LogRunACMETask, params.TaskId)

	runResp, err := this.RPC().ACMETaskRPC().RunACMETask(this.AdminContext(), &pb.RunACMETaskRequest{AcmeTaskId: params.TaskId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if runResp.IsOk {
		this.Data["certId"] = runResp.SslCertId
		this.Success()
	} else {
		this.Fail(runResp.Error)
	}
}
