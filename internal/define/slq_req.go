package define

import "EAMSbackend/util"

type SLQreq struct {
	SoftwareName string
	Version      string
	Status       string
	Location     util.NullString
}
