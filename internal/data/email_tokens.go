package data

import "time"

type EmailTokensQ interface {
	New() EmailTokensQ

	Create(token EmailToken) error
	Verify(email, token string) (bool, error)
	Get(email string) (*EmailToken, error)

	MarkUnsent(id int64) error
	MarkSent(id int64) error
	GetUnsent() ([]EmailToken, error)
}

type EmailToken struct {
	Id         int64      `db:"id" structs:"-"`
	Token      string     `db:"token" structs:"token"`
	LastSentAt *time.Time `db:"last_sent_at" structs:"last_sent_at"`
	Confirmed  bool       `db:"confirmed" structs:"confirmed"`
	Email      string     `db:"email" structs:"email"`
}
