package savedsearches

import "fmt"

type ResponderMeta struct {
	meta string
}

func (rm *ResponderMeta) validate() {
	fmt.Println("Responder")
}

type CreateAlertInput struct {
	Message       string `json:"message,omitempty"`
	Alias         string `json:"alias,omitempty"`
	Description   string `json:"description,omitempty"`
	ResponderMeta `json:"responders"`
}

type CreateAlertRequest struct {
	Uri              string
	CreateAlertInput *CreateAlertInput
}

func (rm *CreateAlertInput) validate() {
	rm.ResponderMeta.validate()
	fmt.Println("CreateAlertInput")
}

func NewCreateAlert(input *CreateAlertInput) {
	input.validate()
}
