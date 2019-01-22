package alert

type Responder interface {
	SetID(id string)
}

type Team struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (t *Team) SetID(id string) {
	t.ID = id
}

func (t *Team) SetName(name string) {
	t.Name = name
}

func (t *Team) getID() string {
	return t.ID
}

func (t *Team) getName() string {
	return t.Name
}

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

func (u *User) SetID(id string) {
	u.ID = id
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

type Escalation struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (e *Escalation) SetID(id string) {
	e.ID = id
}

func (e *Escalation) SetUsername(name string) {
	e.Name = name
}

type Schedule struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (s *Schedule) SetID(id string) {
	s.ID = id
}

func (s *Schedule) SetUsername(name string) {
	s.Name = name
}

type ResponderDTO struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type"`
}

func (r *ResponderDTO) SetID(id string) {
	r.Id = id
}

func (r *ResponderDTO) SetName(id string) {
	r.Id = id
}

func (r *ResponderDTO) getName() string {
	return r.Name
}

func (r *ResponderDTO) getID() string {
	return r.Id
}
