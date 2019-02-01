package contact

import (
	"github.com/pkg/errors"
)

type CreateRequest struct {
	UserIdentifier  string
	To              string     `json:"to"`
	MethodOfContact MethodType `json:"method"`
}

func (cr CreateRequest) Validate() error {
	if cr.UserIdentifier == "" {
		return errors.New("User identifier cannot be empty.")
	}
	if cr.To == "" {
		return errors.New("to cannot be empty.")
	}
	if cr.MethodOfContact == "" {
		return errors.New("Method cannot be empty.")
	}

	return nil
}

func (cr CreateRequest) Endpoint() string {
	return "/v2/users/" + cr.UserIdentifier + "/contacts"
}

func (cr CreateRequest) Method() string {
	return "POST"
}

type GetRequest struct {
	UserIdentifier    string
	ContactIdentifier string
}

func (gr GetRequest) Validate() error {
	err := validateIdentifier(gr.UserIdentifier, gr.ContactIdentifier)
	if err != nil {
		return err
	}
	return nil
}

func (gr GetRequest) Endpoint() string {
	return "/v2/users/" + gr.UserIdentifier + "/contacts/" + gr.ContactIdentifier
}

func (gr GetRequest) Method() string {
	return "GET"
}

type UpdateRequest struct {
	UserIdentifier    string
	ContactIdentifier string
	To                string `json:"to"`
}

func (ur UpdateRequest) Validate() error {

	err := validateIdentifier(ur.UserIdentifier, ur.ContactIdentifier)
	if err != nil {
		return err
	}

	if ur.To == "" {
		return errors.New("to cannot be empty.")
	}

	return nil
}

func (ur UpdateRequest) Endpoint() string {
	return "/v2/users/" + ur.UserIdentifier + "/contacts/" + ur.ContactIdentifier
}

func (ur UpdateRequest) Method() string {
	return "PATCH"
}

type DeleteRequest struct {
	UserIdentifier    string
	ContactIdentifier string
}

func (dr DeleteRequest) Validate() error {
	err := validateIdentifier(dr.UserIdentifier, dr.ContactIdentifier)
	if err != nil {
		return err
	}
	return nil
}
func (dr DeleteRequest) Endpoint() string {
	return "/v2/users/" + dr.UserIdentifier + "/contacts/" + dr.ContactIdentifier
}

func (dr DeleteRequest) Method() string {
	return "DELETE"
}

type ListRequest struct {
	UserIdentifier string
}

func (lr ListRequest) Validate() error {
	if lr.UserIdentifier == "" {
		return errors.New("User identifier cannot be empty.")
	}
	return nil
}
func (lr ListRequest) Endpoint() string {
	return "/v2/users/" + lr.UserIdentifier + "/contacts"
}

func (lr ListRequest) Method() string {
	return "GET"
}

type EnableRequest struct {
	UserIdentifier    string
	ContactIdentifier string
}

func (er EnableRequest) Validate() error {
	err := validateIdentifier(er.UserIdentifier, er.ContactIdentifier)
	if err != nil {
		return err
	}
	return nil
}
func (er EnableRequest) Endpoint() string {
	return "/v2/users/" + er.UserIdentifier + "/contacts/" + er.ContactIdentifier + "/enable"
}

func (er EnableRequest) Method() string {
	return "POST"
}

type DisableRequest struct {
	UserIdentifier    string
	ContactIdentifier string
}

func (dr DisableRequest) Validate() error {
	err := validateIdentifier(dr.UserIdentifier, dr.ContactIdentifier)
	if err != nil {
		return err
	}
	return nil
}
func (dr DisableRequest) Endpoint() string {
	return "/v2/users/" + dr.UserIdentifier + "/contacts/" + dr.ContactIdentifier + "/disable"
}

func (dr DisableRequest) Method() string {
	return "POST"
}

func validateIdentifier(userIdentifier string, contactIdentifier string) error {
	if userIdentifier == "" {
		return errors.New("User identifier cannot be empty.")
	}
	if contactIdentifier == "" {
		return errors.New("Contact identifier cannot be empty.")

	}
	return nil
}

type MethodType string

const (
	Sms   MethodType = "sms"
	Email MethodType = "email"
	Voice MethodType = "voice"
)
