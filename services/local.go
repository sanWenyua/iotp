package services

import (
	"github.com/juju/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Local struct {
	*ServicesStat
}

// Stat get all system resources usage
func (l *Local) Stat() error {
	l.Name = "localhost"

	// Get all cpu usage
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return errors.Annotate(err, "Get cpu usage fail")
	}
	l.CPUUsage = cpuUsage[0]

	// Get all disk usage
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return errors.Annotate(err, "Get disk usage fail")
	}
	l.DiskUsage = diskUsage.UsedPercent

	// Get all memory usage
	memUsage, err := mem.VirtualMemory()
	if err != nil {
		return errors.Annotate(err, "Get memory usage fail")
	}
	l.MemoryUsage = memUsage.UsedPercent

	return nil
}

func NewLocal() *Local {

	return new(Local)
}
