// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package accounts

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	AccountId int64
}) {
	defer this.CreateLogInfo(codes.ACMEProviderAccount_LogDeleteACMEProviderAccount, params.AccountId)

	_, err := this.RPC().ACMEProviderAccountRPC().DeleteACMEProviderAccount(this.AdminContext(), &pb.DeleteACMEProviderAccountRequest{AcmeProviderAccountId: params.AccountId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
