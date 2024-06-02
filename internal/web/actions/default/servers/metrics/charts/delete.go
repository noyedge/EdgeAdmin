// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package charts

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ChartId int64
}) {
	defer this.CreateLogInfo(codes.MetricChart_LogDeleteMetricChart, params.ChartId)

	_, err := this.RPC().MetricChartRPC().DeleteMetricChart(this.AdminContext(), &pb.DeleteMetricChartRequest{MetricChartId: params.ChartId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
