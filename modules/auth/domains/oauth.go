package domains

import "encoding/json"

type OauthSession struct {
	URL      string `redis:"-"`
	Provider string `redis:"Provider"`
	State    string `redis:"State"`
	Verifier string `redis:"Verifier"`
}

func (o *OauthSession) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *OauthSession) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, o)
}
