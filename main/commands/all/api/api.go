package api

import (
	"github.com/xtls/xray-core/main/commands/base"
)

// CmdAPI calls an API in an Auray process.
var CmdAPI = &base.Command{
	UsageLine: "{{.Exec}} api",
	Short:     "Call an API in an Auray process",
	Long: `{{.Exec}} {{.LongName}} provides tools to manipulate Auray via its API.
`,
	Commands: []*base.Command{
		cmdRestartLogger,
		cmdGetStats,
		cmdQueryStats,
		cmdSysStats,
		cmdBalancerInfo,
		cmdBalancerOverride,
		cmdAddInbounds,
		cmdAddOutbounds,
		cmdRemoveInbounds,
		cmdRemoveOutbounds,
		cmdListInbounds,
		cmdListOutbounds,
		cmdAddInboundUsers,
		cmdRemoveInboundUsers,
		cmdInboundUser,
		cmdInboundUserCount,
		cmdAddRules,
		cmdRemoveRules,
		cmdListRules,
		cmdSourceIpBlock,
		cmdOnlineStats,
		cmdOnlineStatsIpList,
		cmdGetAllOnlineUsers,
	},
}
