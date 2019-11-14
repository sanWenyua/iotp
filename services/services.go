package services

type ServicesStat struct {
	Name        string
	MemoryUsage float64
	CPUUsage    float64
	DiskUsage   float64
}

type Services interface {
	Stat() ServicesStat
	Start() error
	Stop() error
}
