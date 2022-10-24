package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokene/keyserver-svc/internal/data"
)

const kdfTableName = "kdf"

func NewKDFQ(db *pgdb.DB) data.KDFQ {
	return &KDFQ{
		db:  db.Clone(),
		sql: sq.Select("*").From(kdfTableName),
	}
}

type KDFQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *KDFQ) New() data.KDFQ {
	return NewKDFQ(q.db)
}

func (q *KDFQ) Get() (*data.KDF, error) {
	var result data.KDF
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *KDFQ) Select() ([]data.KDF, error) {
	var result []data.KDF
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *KDFQ) FilterByKDFVersion(version int) data.KDFQ {
	q.sql = q.sql.Where(sq.Eq{"version": version})
	return q
}
