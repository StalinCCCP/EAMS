package define

import "time"

type SMLQresp struct {
	MaintenanceProcessID uint
	SoftwareID           uint
	MaintenanceDate      time.Time
	Cost                 float64
	Status               string
}
