package escalation

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type CreateEscalationResult struct {
	client.ResultMetadata
	Result string `json:"result"`
	Data   Data   `json:"data"`
}

type Data struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (cer *CreateEscalationResult) ValidateResultMetadata() error {
	if cer.Data == (Data{}) || len(cer.Result) == 0 {
		return errors.New("Could not retrieve create escalation result.")
	}

	return nil
}

type GetEscalationResult struct {
	client.ResultMetadata
	Escalation *Escalation `json:"data"`
}

func (ger *GetEscalationResult) ValidateResultMetadata() error {
	if ger.Escalation == (&Escalation{}) {
		return errors.New("Could not retrieve get escalation result.")
	}

	return nil
}
