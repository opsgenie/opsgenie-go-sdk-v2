package deployment

type Environment struct {
	Type        EnvironmentType `json:"type"`
	Id          string          `json:"id"`
	DisplayName string          `json:"displayName, omitempty"`
}
