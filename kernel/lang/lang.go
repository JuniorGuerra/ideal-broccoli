package lang

type Message struct {
	ID      string `json:"id,omitempty" lang:"langId"`
	Message string `json:"message,omitempty" lang:"langId"`
}

type HttpErrors struct {
	GeneralInternalError Message
	JsonStructureError   Message
}

var Errors HttpErrors

func init() {
	Errors = HttpErrors{
		GeneralInternalError: Message{
			ID:      "GENERAL_INTERNAL_ERROR",
			Message: "General internal error",
		},
		JsonStructureError: Message{
			ID:      "VALIDATION_ERROR",
			Message: "Validation required tag",
		},
	}
}
