package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/internal/data"
	"time"
)

const emailTokensTableName = "email_tokens"

func NewEmailTokensQ(db *pgdb.DB) data.EmailTokensQ {
	return &EmailTokensQ{
		db:  db.Clone(),
		sql: sq.Select("*").From(emailTokensTableName),
	}
}

type EmailTokensQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *EmailTokensQ) New() data.EmailTokensQ {
	return NewEmailTokensQ(q.db.Clone())
}

func (q *EmailTokensQ) Create(token data.EmailToken) error {
	clauses := structs.Map(token)

	stmt := sq.Insert(emailTokensTableName).SetMap(clauses)

	return q.db.Exec(stmt)
}

func (q *EmailTokensQ) Verify(email, token string) (bool, error) {
	stmt := sq.
		Update(emailTokensTableName).
		Set("confirmed", true).
		Where(sq.Eq{"email": email, "token": token})

	result, err := q.db.ExecWithResult(stmt)
	if err != nil {
		return false, errors.Wrap(err, "update failed")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "failed to get rows affected")
	}

	return rows > 0, nil
}

func (q *EmailTokensQ) Get(email string) (*data.EmailToken, error) {
	var token data.EmailToken

	q.sql = q.sql.Where(sq.Eq{"email": email})
	err := q.db.Get(&token, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &token, err
}

func (q *EmailTokensQ) MarkUnsent(id int64) error {
	stmt := sq.
		Update(emailTokensTableName).
		Set("last_sent_at", nil).
		Where(sq.Eq{"id": id})

	err := q.db.Exec(stmt)

	return err
}

func (q *EmailTokensQ) MarkSent(id int64) error {
	stmt := sq.
		Update(emailTokensTableName).
		Set("last_sent_at", time.Now().UTC()).
		Where(sq.Eq{"id": id})

	err := q.db.Exec(stmt)

	return err
}

func (q *EmailTokensQ) GetUnsent() ([]data.EmailToken, error) {
	stmt := q.sql.Where(sq.Eq{"confirmed": false, "last_sent_at": nil})

	var result []data.EmailToken

	err := q.db.Select(&result, stmt)

	return result, err
}
