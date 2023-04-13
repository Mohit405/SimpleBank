package token

import "time"

type Maker interface {
	CreateToken(username string, durationt time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

