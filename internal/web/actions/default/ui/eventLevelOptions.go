package ui

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

type EventLevelOptionsAction struct {
	actionutils.ParentAction
}

func (this *EventLevelOptionsAction) RunPost(params struct{}) {
	this.Data["eventLevels"] = firewallconfigs.FindAllFirewallEventLevels()

	this.Success()
}
