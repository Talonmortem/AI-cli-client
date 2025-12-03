package host

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type Host struct {
	HostInfo *host.InfoStat
}

type Metrics struct {
	CPU    []cpu.InfoStat
	Memory *mem.VirtualMemoryStat
	Disk   *disk.UsageStat
}

func NewHost() (*Host, error) {

	// Host / OS
	hostInfo, _ := host.Info()

	return &Host{
		HostInfo: hostInfo,
	}, nil

}

func (h *Host) GetHostMetrics() *Metrics {
	// CPU info
	cpuInfo, _ := cpu.Info()

	// Memory
	vm, _ := mem.VirtualMemory()

	// Disk
	dk, _ := disk.Usage("/")

	return &Metrics{
		CPU:    cpuInfo,
		Memory: vm,
		Disk:   dk,
	}
}
