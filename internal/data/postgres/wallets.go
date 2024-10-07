package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokene/keyserver-svc/internal/data"
)

const walletsTableName = "wallets"

func NewWalletsQ(db *pgdb.DB) data.WalletsQ {
	return &WalletsQ{
		db:  db.Clone(),
		sql: sq.Select("*").From(walletsTableName),
	}
}

type WalletsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *WalletsQ) New() data.WalletsQ {
	return NewWalletsQ(q.db.Clone())
}

func (q *WalletsQ) Get() (*data.Wallet, error) {
	var result data.Wallet
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *WalletsQ) Select() ([]data.Wallet, error) {
	var result []data.Wallet
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *WalletsQ) Update(wallet data.Wallet) error {
	clauses := structs.Map(wallet)
	stmt := sq.Update(walletsTableName).SetMap(clauses).Where(sq.Eq{"id": wallet.Id})

	return q.db.Exec(stmt)
}

func (q *WalletsQ) Create(wallet data.Wallet) (int64, error) {
	clauses := structs.Map(wallet)
	var id int64

	stmt := sq.Insert(walletsTableName).SetMap(clauses).Suffix("returning id")
	err := q.db.Get(&id, stmt)

	return id, err
}

func (q *WalletsQ) FilterByEmail(email string) data.WalletsQ {
	q.sql = q.sql.Where(sq.Eq{"email": email})
	return q
}

func (q *WalletsQ) FilterByWalletID(walletID string) data.WalletsQ {
	q.sql = q.sql.Where(sq.Eq{"wallet_id": walletID})
	return q
}

func (q *WalletsQ) Delete(walletID string) error {
	stmt := sq.Delete(walletsTableName).Where(sq.Eq{"id": walletID})

	return q.db.Exec(stmt)
}
