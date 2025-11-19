package domain_models

import domain_enums "github.com/alaa-aqeel/looply-app/src/core/Domain/enums"

type Commands struct {
	ID          string
	CliendId    string
	Value       string // value of command mybe image or text
	Type        string // image, text or file
	Results     string // response of commands
	ProcessedBy string
	Status      domain_enums.CommandStatus // pending, process, done, failed

	Client *Client
}

func (m *Commands) SetStatus(val int) (*Commands, error) {
	status, err := domain_enums.NewCommandStatus(val)
	if err != nil {
		return nil, err
	}

	m.Status = status

	return m, nil
}
