// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplists

import (
	"github.com/noyedge/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/noyedge/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

type LevelOptionsAction struct {
	actionutils.ParentAction
}

func (this *LevelOptionsAction) RunPost(params struct{}) {
	this.Data["levels"] = firewallconfigs.FindAllFirewallEventLevels()

	this.Success()
}
