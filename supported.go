package main

import (
	"fmt"

	"github.com/prometheus/procfs"
)

type ProcDirectory struct {
	Base string
	Show func(fpath string, fs procfs.FS) (string, error)
}

type PidDirectory struct {
	Base string
	Show func(fpath string, fs procfs.Proc) (string, error)
}

func PidFiles(pid int) []PidDirectory {
	base := "/proc/" + fmt.Sprint(pid)

	return []PidDirectory{
		{
			Base: base + "/cgroup",
			Show: PidCgroup,
		},
		{
			Base: base + "/cmdline",
			Show: PidCmdline,
		},
		{
			Base: base + "/comm",
			Show: PidComm,
		},
		{
			Base: base + "/cwd",
			Show: PidCwd,
		},
		{
			Base: base + "/environ",
			Show: PidEnviron,
		},
		{
			Base: base + "/limits",
			Show: PidLimits,
		},
		{
			Base: base + "/net/dev",
			Show: PidNetDev,
		},
		{
			Base: base + "/maps",
			Show: PidMaps,
		},
		{
			Base: base + "/schedstat",
			Show: PidSchedstat,
		},
		{
			Base: base + "/stat",
			Show: PidStat,
		},
	}
}

func SupportedFiles() []ProcDirectory {
	return []ProcDirectory{
		{
			Base: "/proc/pressure",
			Show: Pressure,
		},
		{
			Base: "/proc/net/dev",
			Show: NetDev,
		},
		{
			Base: "/proc/buddyinfo",
			Show: BuddyInfo,
		},
		{
			Base: "/proc/cpuinfo",
			Show: CPUInfo,
		},
		{
			Base: "/proc/cmdline",
			Show: CmdLine,
		},
		{
			Base: "/proc/loadavg",
			Show: LoadAvg,
		},
		{
			Base: "/proc/mdstat",
			Show: MdStat,
		},
		{
			Base: "/proc/meminfo",
			Show: MemInfo,
		},
		{
			Base: "/proc/net/protocols",
			Show: NetProcotols,
		},
		{
			Base: "/proc/net/sockstat",
			Show: NetSockstat,
		},
		{
			Base: "/proc/net/sockstat6",
			Show: NetSockstat6,
		},
		{
			Base: "/proc/net/stat",
			Show: NetStat,
		},
		{
			Base: "/proc/net/tcp",
			Show: NetTcp,
		},
		{
			Base: "/proc/net/tcp6",
			Show: NetTcp6,
		},
		{
			Base: "/proc/net/udp",
			Show: NetUdp,
		},
		{
			Base: "/proc/net/udp6",
			Show: NetUdp6,
		},
		{
			Base: "/proc/net/unix",
			Show: NetUnix,
		},
		{
			Base: "/proc/schedstat",
			Show: Schedstat,
		},
		{
			Base: "/proc/slabinfo",
			Show: Slabinfo,
		},
		{
			Base: "/proc/stat",
			Show: Stat,
		},
		{
			Base: "/proc/swaps",
			Show: Swaps,
		},
		{
			Base: "/proc/vmstat",
			Show: VmStat,
		},
		{
			Base: "/proc/zoneinfo",
			Show: ZoneInfo,
		},
	}
}
