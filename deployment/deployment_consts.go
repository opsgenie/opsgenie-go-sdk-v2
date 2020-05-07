package deployment

type State string

const (
	Initial    State = "initial"
	Started          = "started"
	Successful       = "successful"
	Failed           = "failed"
)

type EnvironmentType string

const (
	Test       EnvironmentType = "Test"
	Staging                    = "Staging"
	Production                 = "Production"
)
