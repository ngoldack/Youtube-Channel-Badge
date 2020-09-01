package internal

import (
	"encoding/json"
)

type EndpointResponse struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
	NamedLogo     string `json:"namedLogo"`
}

func ConvertToJson(l, m string) (s string) {
	er := EndpointResponse{
		SchemaVersion: 1,
		Label:         l,
		Message:       m,
		Color:         "Red",
		NamedLogo:     "Youtube",
	}

	bs, _ := json.Marshal(er)
	s = string(bs)
	return
}
