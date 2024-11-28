package domains

import (
	"encoding/json"
	"time"
)

type SessionUser struct {
	FullName  string    `redis:"FullName"`
	Email     string    `redis:"Email"`
	CreatedAt time.Time `redis:"CreatedAt"`
	PubKey    string    `redis:"PubKey"`
	Role      string    `redis:"Role"`
}

func (s *SessionUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SessionUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
