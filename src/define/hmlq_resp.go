package define

import "time"

type HMLQresp struct {
	MaintenanceProcessID uint
	HardwareID           uint
	MaintenanceDate      time.Time
	Cost                 float64
	Status               string
}
