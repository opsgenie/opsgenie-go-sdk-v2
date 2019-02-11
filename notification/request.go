package notification

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
)

type CreateRuleStepRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
	Contact        og.Contact    `json:"contact"`
	SendAfter      *og.SendAfter `json:"sendAfter,omitempty"`
	Enabled        bool          `json:"enabled,omitempty"`
}

func (crs CreateRuleStepRequest) Validate() error {
	err := validateRuleIdentifier(crs.UserIdentifier, crs.RuleId)
	if err != nil {
		return err
	}

	err = validateContact(&crs.Contact)
	if err != nil {
		return err
	}
	return nil
}

func (crs CreateRuleStepRequest) ResourcePath() string {

	return "/v2/users/" + crs.UserIdentifier + "/notification-rules/" + crs.RuleId + "/steps"
}

func (crs CreateRuleStepRequest) Method() string {
	return "POST"
}

type GetRuleStepRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
	RuleStepId     string
}

func (grs GetRuleStepRequest) Validate() error {
	err := validateRuleStepIdentifier(grs.UserIdentifier, grs.RuleId, grs.RuleStepId)
	if err != nil {
		return err
	}
	return nil
}

func (grs GetRuleStepRequest) ResourcePath() string {

	return "/v2/users/" + grs.UserIdentifier + "/notification-rules/" + grs.RuleId + "/steps/" + grs.RuleStepId
}

func (grs GetRuleStepRequest) Method() string {
	return "GET"
}

type UpdateRuleStepRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
	RuleStepId     string
	Contact        *og.Contact   `json:"contact,omitempty"`
	SendAfter      *og.SendAfter `json:"sendAfter,omitempty"`
	Enabled        bool          `json:"enabled,omitempty"`
}

func (urs UpdateRuleStepRequest) Validate() error {
	err := validateRuleStepIdentifier(urs.UserIdentifier, urs.RuleId, urs.RuleStepId)
	if err != nil {
		return err
	}
	if urs.Contact != nil {
		err = validateContact(urs.Contact)
		if err != nil {
			return err
		}
	}

	return nil
}

func (urs UpdateRuleStepRequest) ResourcePath() string {

	return "/v2/users/" + urs.UserIdentifier + "/notification-rules/" + urs.RuleId + "/steps/" + urs.RuleStepId
}

func (urs UpdateRuleStepRequest) Method() string {
	return "PATCH"
}

type DeleteRuleStepRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
	RuleStepId     string
}

func (drs DeleteRuleStepRequest) Validate() error {
	err := validateRuleStepIdentifier(drs.UserIdentifier, drs.RuleId, drs.RuleStepId)
	if err != nil {
		return err
	}
	return nil
}

func (drs DeleteRuleStepRequest) ResourcePath() string {

	return "/v2/users/" + drs.UserIdentifier + "/notification-rules/" + drs.RuleId + "/steps/" + drs.RuleStepId
}

func (drs DeleteRuleStepRequest) Method() string {
	return "DELETE"
}

type ListRuleStepsRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
}

func (lrs ListRuleStepsRequest) Validate() error {
	err := validateRuleIdentifier(lrs.UserIdentifier, lrs.RuleId)
	if err != nil {
		return err
	}
	return nil
}

func (lrs ListRuleStepsRequest) ResourcePath() string {

	return "/v2/users/" + lrs.UserIdentifier + "/notification-rules/" + lrs.RuleId + "/steps"
}

func (lrs ListRuleStepsRequest) Method() string {
	return "GET"
}

type EnableRuleStepRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
	RuleStepId     string
}

func (ers EnableRuleStepRequest) Validate() error {
	err := validateRuleStepIdentifier(ers.UserIdentifier, ers.RuleId, ers.RuleStepId)
	if err != nil {
		return err
	}
	return nil
}

func (ers EnableRuleStepRequest) ResourcePath() string {

	return "/v2/users/" + ers.UserIdentifier + "/notification-rules/" + ers.RuleId + "/steps/" + ers.RuleStepId + "/enable"
}

func (ers EnableRuleStepRequest) Method() string {
	return "POST"
}

type DisableRuleStepRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
	RuleStepId     string
}

func (drs DisableRuleStepRequest) Validate() error {
	err := validateRuleStepIdentifier(drs.UserIdentifier, drs.RuleId, drs.RuleStepId)
	if err != nil {
		return err
	}
	return nil
}

func (drs DisableRuleStepRequest) ResourcePath() string {

	return "/v2/users/" + drs.UserIdentifier + "/notification-rules/" + drs.RuleId + "/steps/" + drs.RuleStepId + "/disable"
}

func (drs DisableRuleStepRequest) Method() string {
	return "POST"
}

type CreateRuleRequest struct {
	client.BaseRequest
	UserIdentifier   string
	Name             string                 `json:"name"`
	ActionType       ActionType             `json:"actionType"`
	Criteria         *og.Criteria           `json:"criteria,omitempty"`
	NotificationTime []NotificationTimeType `json:"notificationTime,omitempty"`
	TimeRestriction  *og.TimeRestriction    `json:"timeRestriction,omitempty"`
	Schedules        []Schedule             `json:"schedules,omitempty"`
	Steps            []*og.Step             `json:"step,omitempty"`
	Order            uint32                 `json:"order,omitempty"`
	Repeat           *Repeat                `json:"repeat,omitempty"`
	Enabled          bool                   `json:"enabled,omitempty"`
}

func (crr CreateRuleRequest) Validate() error {
	if crr.UserIdentifier == "" {
		return errors.New("User identifier cannot be empty.")
	}
	if crr.Name == "" {
		return errors.New("Name cannot be empty.")
	}
	if crr.ActionType == "" {
		return errors.New("Action type cannot be empty.")
	}
	if (crr.ActionType == ScheduleStart || crr.ActionType == ScheduleEnd) && len(crr.NotificationTime) == 0 {
		return errors.New("Notification time cannot be empty.")
	}
	if len(crr.Schedules) != 0 {
		for _, schedule := range crr.Schedules {
			err := validateSchedule(schedule)
			if err != nil {
				return err
			}
		}
	}

	if len(crr.Steps) != 0 {
		for _, step := range crr.Steps {
			err := validateStep(step, crr.ActionType)
			if err != nil {
				return err
			}
		}
	}
	if crr.Criteria != nil {
		err := validateCriteria(crr.Criteria)
		if err != nil {
			return err
		}
	}

	if crr.TimeRestriction != nil {
		err := validateTimeRestriction(crr.TimeRestriction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (crr CreateRuleRequest) ResourcePath() string {

	return "/v2/users/" + crr.UserIdentifier + "/notification-rules"
}

func (crr CreateRuleRequest) Method() string {
	return "POST"
}

type GetRuleRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
}

func (grr GetRuleRequest) Validate() error {
	err := validateRuleIdentifier(grr.UserIdentifier, grr.RuleId)
	if err != nil {
		return err
	}
	return nil
}

func (grr GetRuleRequest) ResourcePath() string {

	return "/v2/users/" + grr.UserIdentifier + "/notification-rules/" + grr.RuleId
}

func (grr GetRuleRequest) Method() string {
	return "GET"
}

type UpdateRuleRequest struct {
	client.BaseRequest
	UserIdentifier   string
	RuleId           string
	Criteria         *og.Criteria           `json:"criteria,omitempty"`
	NotificationTime []NotificationTimeType `json:"notificationTime,omitempty"`
	TimeRestriction  *og.TimeRestriction    `json:"timeRestriction,omitempty"`
	Schedules        []Schedule             `json:"schedules,omitempty"`
	Steps            []*og.Step             `json:"step,omitempty"`
	Order            uint32                 `json:"order,omitempty"`
	Repeat           *Repeat                `json:"repeat,omitempty"`
	Enabled          bool                   `json:"enabled,omitempty"`
}

func (urr UpdateRuleRequest) Validate() error {
	err := validateRuleIdentifier(urr.UserIdentifier, urr.RuleId)
	if err != nil {
		return err
	}
	if len(urr.Schedules) != 0 {
		for _, schedule := range urr.Schedules {
			err := validateSchedule(schedule)
			if err != nil {
				return err
			}
		}
	}

	if len(urr.Steps) != 0 {
		for _, step := range urr.Steps {
			err := validateStepWithoutActionTypeInfo(step)
			if err != nil {
				return err
			}
		}
	}
	if urr.Criteria != nil {
		err := validateCriteria(urr.Criteria)
		if err != nil {
			return err
		}
	}

	if urr.TimeRestriction != nil {
		err := validateTimeRestriction(urr.TimeRestriction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (urr UpdateRuleRequest) ResourcePath() string {

	return "/v2/users/" + urr.UserIdentifier + "/notification-rules/" + urr.RuleId
}

func (urr UpdateRuleRequest) Method() string {
	return "PATCH"
}

type DeleteRuleRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
}

func (drr DeleteRuleRequest) Validate() error {
	err := validateRuleIdentifier(drr.UserIdentifier, drr.RuleId)
	if err != nil {
		return err
	}
	return nil
}

func (drr DeleteRuleRequest) ResourcePath() string {

	return "/v2/users/" + drr.UserIdentifier + "/notification-rules/" + drr.RuleId
}

func (drr DeleteRuleRequest) Method() string {
	return "DELETE"
}

type ListRuleRequest struct {
	client.BaseRequest
	UserIdentifier string
}

func (lrr ListRuleRequest) Validate() error {
	if lrr.UserIdentifier == "" {
		return errors.New("User identifier cannot be empty.")
	}
	return nil
}

func (lrr ListRuleRequest) ResourcePath() string {

	return "/v2/users/" + lrr.UserIdentifier + "/notification-rules"
}

func (lrr ListRuleRequest) Method() string {
	return "GET"
}

type EnableRuleRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
}

func (enr EnableRuleRequest) Validate() error {
	err := validateRuleIdentifier(enr.UserIdentifier, enr.RuleId)
	if err != nil {
		return err
	}
	return nil
}

func (enr EnableRuleRequest) ResourcePath() string {

	return "/v2/users/" + enr.UserIdentifier + "/notification-rules/" + enr.RuleId + "/enable"
}

func (enr EnableRuleRequest) Method() string {
	return "POST"
}

type DisableRuleRequest struct {
	client.BaseRequest
	UserIdentifier string
	RuleId         string
}

func (drr DisableRuleRequest) Validate() error {
	err := validateRuleIdentifier(drr.UserIdentifier, drr.RuleId)
	if err != nil {
		return err
	}
	return nil
}

func (drr DisableRuleRequest) ResourcePath() string {

	return "/v2/users/" + drr.UserIdentifier + "/notification-rules/" + drr.RuleId + "/disable"
}

func (drr DisableRuleRequest) Method() string {
	return "POST"
}

type CopyNotificationRulesRequest struct {
	client.BaseRequest
	UserIdentifier string
	ToUsers        []string    `json:"toUsers"`
	RuleTypes      []RuleTypes `json:"ruleTypes"`
}

func (cnr CopyNotificationRulesRequest) Validate() error {
	if cnr.UserIdentifier == "" {
		return errors.New("User identifier cannot be empty.")
	}
	if len(cnr.ToUsers) == 0 {
		return errors.New("You must specify a list of the users which you want to copy the rules to.")
	}
	if len(cnr.RuleTypes) == 0 {
		return errors.New("Specify a list of the action types you want to copy the rules of.")
	}
	return nil
}

func (cnr CopyNotificationRulesRequest) ResourcePath() string {

	return "/v2/users/" + cnr.UserIdentifier + "/notification-rules/copy-to"
}

func (cnr CopyNotificationRulesRequest) Method() string {
	return "POST"
}

func validateRuleIdentifier(userIdentifier string, ruleIdentifier string) error {
	if userIdentifier == "" {
		return errors.New("User identifier cannot be empty.")
	}
	if ruleIdentifier == "" {
		return errors.New("Rule identifier cannot be empty.")

	}
	return nil
}

func validateRuleStepIdentifier(userIdentifier string, ruleIdentifier string, ruleStepId string) error {
	err := validateRuleIdentifier(userIdentifier, ruleIdentifier)
	if err != nil {
		return err
	}
	if ruleStepId == "" {
		return errors.New("Rule Step identifier cannot be empty.")

	}
	return nil
}

func validateContact(contact *og.Contact) error {
	if contact == nil {
		return errors.New("Contact cannot be empty.")

	}
	if contact.To == "" {
		return errors.New("To cannot be empty.")
	}
	if contact.MethodOfContact == "" {
		return errors.New("Method cannot be empty.")

	}
	return nil
}

type ActionType string

const (
	CreateAlert         ActionType = "create-alert"
	AcknowledgedAlert   ActionType = "acknowledged-alert"
	ClosedAlert         ActionType = "closed-alert"
	AssignedAlert       ActionType = "assigned-alert"
	AddNote             ActionType = "add-note"
	ScheduleStart       ActionType = "schedule-start"
	ScheduleEnd         ActionType = "schedule-end"
	IncomingCallRouting ActionType = "incoming-call-routing"
)

type NotificationTimeType string

const (
	JustBefore        NotificationTimeType = "just-before"
	FifteenMinutesAgo NotificationTimeType = "15-minutes-ago"
	OneHourAgo        NotificationTimeType = "1-hour-ago"
	OneDayAgo         NotificationTimeType = "1-day-ago"
)

type Schedule struct {
	TypeOfSchedule string `json:"type"`
	Name           string `json:"name,omitempty"`
	Id             string `json:"id,omitempty"`
}

func validateSchedule(schedule Schedule) error {
	if schedule.TypeOfSchedule != "schedule" {
		return errors.New("Type of schedule cannot be empty.")
	}
	return nil
}

type Repeat struct {
	LoopAfter uint32 `json:"loopAfter,omitempty"`
	Enabled   bool   `json:"loopAfter,omitempty"`
}

func validateStep(step *og.Step, actionType ActionType) error {
	if step.Contact.To == "" {
		return errors.New("To cannot be empty.")
	}
	if step.Contact.MethodOfContact == "" {
		return errors.New("Method cannot be empty.")
	}
	if (actionType == CreateAlert || actionType == AssignedAlert) && step.SendAfter == nil {
		return errors.New("SendAfter cannot be empty.")
	}

	return nil
}

func validateStepWithoutActionTypeInfo(step *og.Step) error {
	if step.Contact.To == "" {
		return errors.New("To cannot be empty.")
	}
	if step.Contact.MethodOfContact == "" {
		return errors.New("Method cannot be empty.")
	}

	return nil
}

func validateCriteria(criteria *og.Criteria) error {
	if criteria.CriteriaType == "" {
		return errors.New("Type of criteria cannot be empty.")
	}
	if (criteria.CriteriaType == og.MatchAnyCondition || criteria.CriteriaType == og.MatchAllConditions) &&
		criteria.Conditions == nil {
		return errors.New("Conditions cannot be empty.")
	}
	if criteria.Conditions != nil {
		err := og.ValidateConditions(criteria.Conditions)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateTimeRestriction(timeRestriction *og.TimeRestriction) error {
	if timeRestriction.Type == "" {
		return errors.New("Type of time restriction must be time-of-day or weekday-and-time-of-day.")
	}
	if timeRestriction.Type == og.WeekdayAndTimeOfDay && timeRestriction.RestrictionList == nil {
		return errors.New("Restrictions cannot be empty.")
	}
	if len(timeRestriction.RestrictionList) != 0 {
		for _, restriction := range timeRestriction.RestrictionList {
			if timeRestriction.Type == "weekday-and-time-of-day" &&
				(restriction.EndMin == 0 ||
					restriction.EndHour == 0 ||
					restriction.EndDay == "" ||
					restriction.StartDay == "" ||
					restriction.StartHour == 0 ||
					restriction.StartMin == 0) {
				return errors.New("startDay, startHour, startMin, endDay, endHour, endMin cannot be empty.")
			}
		}
	}
	if timeRestriction.Type == "time-of-day" &&
		(timeRestriction.Restriction.EndMin == 0 ||
			timeRestriction.Restriction.EndHour == 0 ||
			timeRestriction.Restriction.StartHour == 0 ||
			timeRestriction.Restriction.StartMin == 0) {
		return errors.New("startHour, startMin, endHour, endMin cannot be empty.")
	}

	return nil
}

type RuleTypes string

const (
	All                   RuleTypes = "all"
	AcknowledgedAlertRule RuleTypes = "acknowledged-alert"
	RenotifiedAlertRule   RuleTypes = "renotified-alert"
	ClosedAlertRule       RuleTypes = "closed-alert"
	ScheduleStartRule     RuleTypes = "schedule-start"
	AssignedAlertRule     RuleTypes = "assigned-alert"
	AddNoteRule           RuleTypes = "add-note"
	NewAlertRule          RuleTypes = "new-alert"
)
