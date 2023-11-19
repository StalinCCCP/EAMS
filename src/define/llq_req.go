package define

import "EAMSbackend/util"

type LLQreq struct {
	LabName  string
	Status   string
	Location util.NullString
}
