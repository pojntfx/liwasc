package models

type Service struct {
	ServiceName             string
	PortNumber              int
	TransportProtocol       string
	Description             string
	Assignee                string
	Contact                 string
	RegistrationDate        string
	ModificationDate        string
	Reference               string
	ServiceCode             string
	UnauthorizedUseReported string
	AssignmentNotes         string
}
