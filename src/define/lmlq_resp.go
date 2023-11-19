package define

import "time"

type LMLQresp struct {
	MaintenanceProcessID uint
	LabID                uint
	MaintenanceDate      time.Time
	Cost                 float64
	Status               string
}
