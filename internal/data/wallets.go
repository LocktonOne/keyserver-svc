package data

type WalletsQ interface {
	New() WalletsQ

	Get() (*Wallet, error)
	Select() ([]Wallet, error)

	Create(wallet Wallet) (int64, error)
	Update(wallet Wallet) error

	FilterByEmail(email string) WalletsQ
	FilterByWalletID(walletID string) WalletsQ
}

type Wallet struct {
	Id           int64  `db:"id" structs:"-"`
	WalletId     string `db:"wallet_id" structs:"wallet_id"`
	Email        string `db:"email" structs:"email"`
	KeychainData string `db:"keychain_data" structs:"keychain_data"`
	Salt         string `db:"salt" structs:"salt"`
	Verified     bool   `db:"verified" structs:"verified"`
}
