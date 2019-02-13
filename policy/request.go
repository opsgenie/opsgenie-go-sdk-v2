package policy

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
)

type CreateAlertPolicyRequest struct {
	client.BaseRequest
	MainFields
	Message                  string             `json:"message,omitempty"`
	Continue                 bool               `json:"continue,omitempty"`
	Alias                    string             `json:"alias,omitempty"`
	AlertDescription         string             `json:"alertDescription,omitempty"`
	Entity                   string             `json:"entity,omitempty"`
	Source                   string             `json:"source,omitempty"`
	IgnoreOriginalDetails    bool               `json:"ignoreOriginalDetails,omitempty"`
	Actions                  []string           `json:"actions,omitempty"`
	IgnoreOriginalActions    bool               `json:"ignoreOriginalActions,omitempty"`
	Details                  []string           `json:"details,omitempty"`
	IgnoreOriginalResponders bool               `json:"ignoreOriginalResponders,omitempty"`
	Responders               *[]alert.Responder `json:"responders,omitempty"`
	IgnoreOriginalTags       bool               `json:"ignoreOriginalTags,omitempty"`
	Tags                     []string           `json:"tags,omitempty"`
	Priority                 alert.Priority     `json:"priority,omitempty"`
}

type CreateNotificationPolicyRequest struct {
	client.BaseRequest
	MainFields
	AutoRestartAction   *AutoRestartAction   `json:"autoRestartAction,omitempty"`
	AutoCloseAction     *AutoCloseAction     `json:"autoCloseAction,omitempty"`
	DeDuplicationAction *DeDuplicationAction `json:"deduplicationActionAction,omitempty"`
	DelayAction         *DelayAction         `json:"delayAction,omitempty"`
	Suppress            bool                 `json:"suppress,omitempty"`
}

type MainFields struct {
	PolicyType        string              `json:"type,omitempty"`
	Name              string              `json:"name,omitempty"`
	Enabled           bool                `json:"enabled"`
	PolicyDescription string              `json:"policyDescription"`
	Filter            *og.Filter          `json:"filter,omitempty"`
	TimeRestriction   *og.TimeRestriction `json:"timeRestrictions,omitempty"`
	TeamId            string
}

func (cap CreateAlertPolicyRequest) Validate() error {
	err := ValidateMainFields(&cap.MainFields)
	if err != nil {
		return err
	}
	if cap.Message == "" {
		return errors.New("alert message cannot be empty")
	}
	if cap.Responders != nil {
		err = ValidateResponders(cap.Responders)
		if err != nil {
			return err
		}
	}
	if cap.Priority != "" {
		err = alert.ValidatePriority(cap.Priority)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cap CreateAlertPolicyRequest) ResourcePath() string {
	return "/v2/policies"
}

func (cap CreateAlertPolicyRequest) Method() string {
	return "POST"
}

func (cnp CreateNotificationPolicyRequest) Validate() error {
	err := ValidateMainFields(&cnp.MainFields)
	if err != nil {
		return err
	}
	if cnp.TeamId == "" {
		return errors.New("policy team id should be provided")
	}
	if cnp.AutoRestartAction != nil {
		err = ValidateAutoRestartAction(*cnp.AutoRestartAction)
		if err != nil {
			return err
		}
	}
	if cnp.AutoCloseAction != nil {
		err = ValidateAutoCloseAction(*cnp.AutoCloseAction)
		if err != nil {
			return err
		}
	}
	if cnp.DeDuplicationAction != nil {
		err = ValidateDeDuplicationAction(*cnp.DeDuplicationAction)
		if err != nil {
			return err
		}
	}
	if cnp.DelayAction != nil {
		err = ValidateDelayAction(*cnp.DelayAction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cnp CreateNotificationPolicyRequest) ResourcePath() string {
	return "/v2/policies"
}

func (cnp CreateNotificationPolicyRequest) RequestParams() map[string]string {
	if cnp.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = cnp.TeamId
	return params
}

func (cnp CreateNotificationPolicyRequest) Method() string {
	return "POST"
}

type GetAlertPolicyRequest struct {
	client.BaseRequest
	Id     string
	TeamId string
}

func (gap GetAlertPolicyRequest) Validate() error {
	if gap.Id == "" {
		return errors.New("policy id should be provided")
	}
	return nil
}

func (gap GetAlertPolicyRequest) ResourcePath() string {
	return "/v2/policies/" + gap.Id
}

func (gap GetAlertPolicyRequest) RequestParams() map[string]string {
	if gap.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = gap.TeamId
	return params
}

func (gap GetAlertPolicyRequest) Method() string {
	return "GET"
}

type GetNotificationPolicyRequest struct {
	client.BaseRequest
	Id     string
	TeamId string
}

func (gnp GetNotificationPolicyRequest) Validate() error {
	if gnp.Id == "" {
		return errors.New("policy id should be provided")
	}
	if gnp.TeamId == "" {
		return errors.New("policy team id should be provided")
	}
	return nil
}

func (gnp GetNotificationPolicyRequest) ResourcePath() string {
	return "/v2/policies/" + gnp.Id
}

func (gnp GetNotificationPolicyRequest) RequestParams() map[string]string {
	if gnp.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = gnp.TeamId
	return params
}

func (gnp GetNotificationPolicyRequest) Method() string {
	return "GET"
}

type UpdateAlertPolicyRequest struct {
	client.BaseRequest
	MainFields
	Message                  string                 `json:"message,omitempty"`
	Continue                 bool                   `json:"continue,omitempty"`
	Alias                    string                 `json:"alias,omitempty"`
	AlertDescription         string                 `json:"alertDescription,omitempty"`
	Entity                   string                 `json:"entity,omitempty"`
	Source                   string                 `json:"source,omitempty"`
	IgnoreOriginalDetails    bool                   `json:"ignoreOriginalDetails,omitempty"`
	Actions                  []string               `json:"actions,omitempty"`
	IgnoreOriginalActions    bool                   `json:"ignoreOriginalActions,omitempty"`
	Details                  map[string]interface{} `json:"details,omitempty"`
	IgnoreOriginalResponders bool                   `json:"ignoreOriginalResponders,omitempty"`
	Responders               *[]alert.Responder     `json:"responders,omitempty"`
	IgnoreOriginalTags       bool                   `json:"ignoreOriginalTags,omitempty"`
	Tags                     []string               `json:"tags,omitempty"`
	Priority                 alert.Priority         `json:"priority,omitempty"`
	Id                       string
}

func (uap UpdateAlertPolicyRequest) Validate() error {
	err := ValidatePolicyIdentifier("alert", uap.Id, uap.TeamId)
	if err != nil {
		return err
	}
	err = ValidateMainFields(&uap.MainFields)
	if err != nil {
		return err
	}
	if uap.Message == "" {
		return errors.New("alert message cannot be empty")
	}
	if uap.Responders != nil {
		err = ValidateResponders(uap.Responders)
		if err != nil {
			return err
		}
	}
	if uap.Priority != "" {
		err = alert.ValidatePriority(uap.Priority)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uap UpdateAlertPolicyRequest) ResourcePath() string {
	return "/v2/policies/" + uap.Id
}

func (uap UpdateAlertPolicyRequest) Method() string {
	return "PUT"
}

type UpdateNotificationPolicyRequest struct {
	client.BaseRequest
	MainFields
	AutoRestartAction   *AutoRestartAction   `json:"autoRestartAction,omitempty"`
	AutoCloseAction     *AutoCloseAction     `json:"autoCloseAction,omitempty"`
	DeDuplicationAction *DeDuplicationAction `json:"deduplicationActionAction,omitempty"`
	DelayAction         *DelayAction         `json:"delayAction,omitempty"`
	Suppress            bool                 `json:"suppress,omitempty"`
	Id                  string
}

func (unp UpdateNotificationPolicyRequest) Validate() error {
	err := ValidatePolicyIdentifier("notification", unp.Id, unp.TeamId)
	if err != nil {
		return err
	}
	err = ValidateMainFields(&unp.MainFields)
	if err != nil {
		return err
	}
	if unp.TeamId == "" {
		return errors.New("policy team id should be provided")
	}
	if unp.AutoRestartAction != nil {
		err = ValidateAutoRestartAction(*unp.AutoRestartAction)
		if err != nil {
			return err
		}
	}
	if unp.AutoCloseAction != nil {
		err = ValidateAutoCloseAction(*unp.AutoCloseAction)
		if err != nil {
			return err
		}
	}
	if unp.DeDuplicationAction != nil {
		err = ValidateDeDuplicationAction(*unp.DeDuplicationAction)
		if err != nil {
			return err
		}
	}
	if unp.DelayAction != nil {
		err = ValidateDelayAction(*unp.DelayAction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (unp UpdateNotificationPolicyRequest) ResourcePath() string {
	return "/v2/policies/" + unp.Id
}

func (unp UpdateNotificationPolicyRequest) RequestParams() map[string]string {
	params := make(map[string]string)
	params["teamId"] = unp.TeamId
	return params
}

func (unp UpdateNotificationPolicyRequest) Method() string {
	return "PUT"
}

type DeletePolicyRequest struct {
	client.BaseRequest
	Id     string `json:"id,omitempty"`
	TeamId string
	Type   PolicyType
}

func (dpr DeletePolicyRequest) Validate() error {
	if dpr.Type != AlertPolicy && dpr.Type != NotificationPolicy {
		return errors.New("policy type should be one of alert or notification")
	}
	err := ValidatePolicyIdentifier(string(dpr.Type), dpr.Id, dpr.TeamId)
	if err != nil {
		return err
	}
	return nil
}

func (dpr DeletePolicyRequest) ResourcePath() string {
	return "/v2/policies/" + dpr.Id
}

func (dpr DeletePolicyRequest) Method() string {
	return "DELETE"
}

func (dpr DeletePolicyRequest) RequestParams() map[string]string {
	if dpr.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = dpr.TeamId
	return params
}

type DisablePolicyRequest struct {
	client.BaseRequest
	Id     string `json:"id,omitempty"`
	TeamId string
	Type   PolicyType
}

func (dpr DisablePolicyRequest) Validate() error {
	if dpr.Type != AlertPolicy && dpr.Type != NotificationPolicy {
		return errors.New("policy type should be one of alert or notification")
	}
	err := ValidatePolicyIdentifier(string(dpr.Type), dpr.Id, dpr.TeamId)
	if err != nil {
		return err
	}
	return nil
}

func (dpr DisablePolicyRequest) ResourcePath() string {
	return "/v2/policies/" + dpr.Id + "/disable"
}

func (dpr DisablePolicyRequest) Method() string {
	return "POST"
}

func (dpr DisablePolicyRequest) RequestParams() map[string]string {
	if dpr.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = dpr.TeamId
	return params
}

type EnablePolicyRequest struct {
	client.BaseRequest
	Id     string `json:"id,omitempty"`
	TeamId string
	Type   PolicyType
}

func (dpr EnablePolicyRequest) Validate() error {
	if dpr.Type != AlertPolicy && dpr.Type != NotificationPolicy {
		return errors.New("policy type should be one of alert or notification")
	}
	err := ValidatePolicyIdentifier(string(dpr.Type), dpr.Id, dpr.TeamId)
	if err != nil {
		return err
	}
	return nil
}

func (dpr EnablePolicyRequest) ResourcePath() string {
	return "/v2/policies/" + dpr.Id + "/enable"
}

func (dpr EnablePolicyRequest) Method() string {
	return "POST"
}

func (dpr EnablePolicyRequest) RequestParams() map[string]string {
	if dpr.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = dpr.TeamId
	return params
}

type ChangeOrderRequest struct {
	client.BaseRequest
	Id          string `json:"id,omitempty"`
	TeamId      string
	Type        PolicyType
	TargetIndex int `json:"targetIndex,omitempty"`
}

func (dpr ChangeOrderRequest) Validate() error {
	if dpr.Type != AlertPolicy && dpr.Type != NotificationPolicy {
		return errors.New("policy type should be one of alert or notification")
	}
	err := ValidatePolicyIdentifier(string(dpr.Type), dpr.Id, dpr.TeamId)
	if err != nil {
		return err
	}
	if dpr.TargetIndex < 0 {
		return errors.New("target index should be at least 0")
	}
	return nil
}

func (dpr ChangeOrderRequest) ResourcePath() string {
	return "/v2/policies/" + dpr.Id + "/change-order"
}

func (dpr ChangeOrderRequest) Method() string {
	return "POST"
}

func (dpr ChangeOrderRequest) RequestParams() map[string]string {
	if dpr.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = dpr.TeamId
	return params
}

type ListAlertPoliciesRequest struct {
	client.BaseRequest
	TeamId string
}

func (dpr ListAlertPoliciesRequest) Validate() error {
	return nil
}

func (dpr ListAlertPoliciesRequest) ResourcePath() string {
	return "/v2/policies/alert"
}

func (dpr ListAlertPoliciesRequest) Method() string {
	return "GET"
}

func (dpr ListAlertPoliciesRequest) RequestParams() map[string]string {
	if dpr.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = dpr.TeamId
	return params
}

type ListNotificationPoliciesRequest struct {
	client.BaseRequest
	TeamId string
}

func (dpr ListNotificationPoliciesRequest) Validate() error {
	if dpr.TeamId == "" {
		return errors.New("team id should be provided")
	}
	return nil
}

func (dpr ListNotificationPoliciesRequest) ResourcePath() string {
	return "/v2/policies/notification"
}

func (dpr ListNotificationPoliciesRequest) Method() string {
	return "GET"
}

func (dpr ListNotificationPoliciesRequest) RequestParams() map[string]string {
	if dpr.TeamId == "" {
		return nil
	}
	params := make(map[string]string)
	params["teamId"] = dpr.TeamId
	return params
}

type PolicyType string

type Duration struct {
	TimeAmount int         `json:"timeAmount,omitempty"`
	TimeUnit   og.TimeUnit `json:"timeUnit,omitempty"`
}

type AutoRestartAction struct {
	Duration       *Duration `json:"duration,omitempty"`
	MaxRepeatCount int       `json:"maxRepeatCount,omitempty"`
}

type AutoCloseAction struct {
	Duration *Duration `json:"duration,omitempty"`
}

type DeDuplicationAction struct {
	DeDuplicationActionType DeDuplicationActionType `json:"deduplicationType,omitempty"`
	Duration                *Duration               `json:"duration,omitempty"`
	Count                   int                     `json:"count,omitempty"`
}

type DelayAction struct {
	DelayOption DelayType `json:"delayOption,omitempty"`
	UntilMinute int       `json:"untilMinute,omitempty"`
	UntilHour   int       `json:"untilHour,omitempty"`
	Duration    *Duration `json:"duration,omitempty"`
}

type DeDuplicationActionType string
type DelayType string

const (
	ValueBased     DeDuplicationActionType = "value-based"
	FrequencyBased DeDuplicationActionType = "frequency-based"

	ForDuration        DelayType  = "for-duration"
	NextTime           DelayType  = "next-time"
	NextWeekday        DelayType  = "next-weekday"
	NextMonday         DelayType  = "next-monday"
	NextTuesday        DelayType  = "next-tuesday"
	NextWednesday      DelayType  = "next-wednesday"
	NextThursday       DelayType  = "next-thursday"
	NextFriday         DelayType  = "next-friday"
	NextSaturday       DelayType  = "next-saturday"
	NextSunday         DelayType  = "next-sunday"
	AlertPolicy        PolicyType = "alert"
	NotificationPolicy PolicyType = "notification"
)

func ValidateMainFields(fields *MainFields) error {
	if fields == nil {
		return errors.New("policy main fields should be provided")
	}
	if fields.PolicyType != "alert" && fields.PolicyType != "notification" {
		return errors.New("policy type should be alert or notification")
	}
	if fields.Name == "" {
		return errors.New("policy name cannot be empty")
	}
	if fields.Filter != nil {
		err := og.ValidateFilter(*fields.Filter)
		if err != nil {
			return err
		}
	}
	if fields.TimeRestriction != nil {
		err := og.ValidateRestrictions(*fields.TimeRestriction)
		if err != nil {
			return err
		}
	}
	return nil
}

func ValidateDuration(duration *Duration) error {
	if duration != nil && duration.TimeUnit != "" && duration.TimeUnit != og.Days && duration.TimeUnit != og.Hours && duration.TimeUnit != og.Minutes {
		return errors.New("timeUnit provided for duration is not valid")
	}
	if duration != nil && duration.TimeAmount <= 0 {
		return errors.New("duration timeAmount should be greater than zero")
	}
	if duration != nil && duration.TimeUnit == "" {
		duration.TimeUnit = og.Minutes
	}
	return nil
}

func ValidateDeDuplicationAction(action DeDuplicationAction) error {
	if action.DeDuplicationActionType != ValueBased && action.DeDuplicationActionType != FrequencyBased {
		return errors.New("deDuplication action type should be one of value-based or frequency-based")
	}
	if action.Duration != nil {
		err := ValidateDuration(action.Duration)
		if err != nil {
			return err
		}
	}
	if action.Count < 0 {
		return errors.New("deDuplication count is not valid")
	}
	return nil
}

func ValidateAutoRestartAction(action AutoRestartAction) error {
	if action.Duration == nil {
		return errors.New("autoRestart action duration cannot be empty")
	}
	err := ValidateDuration(action.Duration)
	if err != nil {
		return err
	}
	if action.MaxRepeatCount < 0 {
		return errors.New("autoRestart maxRepeatCount is not valid")
	}
	return nil
}

func ValidateAutoCloseAction(action AutoCloseAction) error {
	if action.Duration == nil {
		return errors.New("autoClose action duration cannot be empty")
	}
	err := ValidateDuration(action.Duration)
	if err != nil {
		return err
	}
	return nil
}

func ValidateDelayAction(action DelayAction) error {
	if action.DelayOption != ForDuration && action.DelayOption != NextTime && action.DelayOption != NextWeekday &&
		action.DelayOption != NextMonday && action.DelayOption != NextTuesday && action.DelayOption != NextWednesday &&
		action.DelayOption != NextThursday && action.DelayOption != NextFriday && action.DelayOption != NextSaturday && action.DelayOption != NextSunday {
		return errors.New("delay option should be one of for-duration, next-time, next-weekday, next-monday, next-tuesday, next-wednesday, next-thursday, next-friday, next-saturday, next-sunday")
	}
	if action.DelayOption == ForDuration {
		if action.Duration == nil {
			return errors.New("delayAction duration cannot be empty")
		}
		err := ValidateDuration(action.Duration)
		if err != nil {
			return err
		}
	}
	if action.DelayOption != ForDuration && ((action.UntilHour < 0 || action.UntilHour > 23) || (action.UntilMinute < 0 || action.UntilMinute > 59)) {
		return errors.New("delayAction's UntilHour or UntilMinute is not valid")
	}
	return nil
}

func ValidateResponders(responders *[]alert.Responder) error {
	for _, responder := range *responders {
		if responder.Type != alert.UserResponder && responder.Type != alert.TeamResponder {
			return errors.New("responder type for alert policy should be one of team or user")
		}
		if responder.Id == "" {
			return errors.New("responder id should be provided")
		}
	}
	return nil
}

func ValidatePolicyIdentifier(policyType string, id string, teamId string) error {
	if id == "" {
		return errors.New("policy id should be provided")
	}
	if "notification" == policyType && teamId == "" {
		return errors.New("policy team id should be provided")
	}
	return nil
}
