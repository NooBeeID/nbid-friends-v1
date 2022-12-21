package token

import "time"

type Token struct {
	AuthId  int
	Expired time.Time
}
