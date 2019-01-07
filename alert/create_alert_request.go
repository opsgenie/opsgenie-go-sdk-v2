package alert

type CreateAlertInput struct {
	Message     string            `json:"message,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Description string            `json:"description,omitempty"`
	Responders  []ResponderMeta   `json:"responders"`
	VisibleTo   []ResponderMeta   `json:"visibleTo,omitempty"`
	Actions     []string          `json:"actions,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Details     map[string]string `json:"details,omitempty"`
	Entity      string            `json:"entity,omitempty"`
	Source      string            `json:"source,omitempty"`
	Priority    Priority          `json:"priority,omitempty"`
	User        string            `json:"user,omitempty"`
	Note        string            `json:"note,omitempty"`
}

type CreateAlertRequest struct {
	Uri              string
	CreateAlertInput *CreateAlertInput
}

func NewCreateAlertRequest(input *CreateAlertInput) (CreateAlertRequest, error) {

	//init
	uri := generateFullPathWithParams("/v2/alerts", nil)

	return CreateAlertRequest{Uri: uri, CreateAlertInput: input}, nil

}

/*func (r *CreateAlertRequest) Init() {
	if r.Teams != nil {
		var convertedTeams []TeamRecipient
		for _, t := range r.Teams {
			recipient := &RecipientDTO{
				Id:   t.getID(),
				Name: t.getName(),
				Type: "team",
			}

			convertedTeams = append(convertedTeams, recipient)
		}
		r.Teams = convertedTeams
	}

	if r.VisibleTo != nil {
		var convertedVisibleTo []Recipient
		for _, r := range r.VisibleTo {
			switch r.(type) {
			case *Team:
				{
					team := r.(*Team)
					recipient := &RecipientDTO{
						Id:   team.ID,
						Name: team.Name,
						Type: "team",
					}
					convertedVisibleTo = append(convertedVisibleTo, recipient)
				}
			case *User:
				{
					user := r.(*User)
					recipient := &RecipientDTO{
						Id:       user.ID,
						Username: user.Username,
						Type:     "user",
					}
					convertedVisibleTo = append(convertedVisibleTo, recipient)
				}
			}
		}
		r.VisibleTo = convertedVisibleTo
	}
}*/
