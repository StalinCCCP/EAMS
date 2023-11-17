package define

import "EAMSbackend/util"

type HDLQreq struct {
	HardwareName string
	Category     string
	Status       string
	Location     util.NullString
}
