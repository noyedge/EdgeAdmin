package regions

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/langs/codes"
	"github.com/noyedge/EdgeCommon/pkg/rpc/pb"
)

type SortAction struct {
	actionutils.ParentAction
}

func (this *SortAction) RunPost(params struct {
	RegionIds []int64
}) {
	defer this.CreateLogInfo(codes.NodeRegion_LogSortNodeRegions)

	_, err := this.RPC().NodeRegionRPC().UpdateNodeRegionOrders(this.AdminContext(), &pb.UpdateNodeRegionOrdersRequest{NodeRegionIds: params.RegionIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
