package accesskeys

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type UpdateIsOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateIsOnAction) RunPost(params struct {
	AccessKeyId int64
	IsOn        bool
}) {
	defer this.CreateLogInfo(codes.UserAccessKey_LogUpdateUserAccessKeyIsOn, params.AccessKeyId)

	_, err := this.RPC().UserAccessKeyRPC().UpdateUserAccessKeyIsOn(this.AdminContext(), &pb.UpdateUserAccessKeyIsOnRequest{
		UserAccessKeyId: params.AccessKeyId,
		IsOn:            params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
