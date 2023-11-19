package define

import "EAMSbackend/util"

type HLQreq struct {
	HardwareName string
	Category     string
	Status       string
	Location     util.NullString
}
