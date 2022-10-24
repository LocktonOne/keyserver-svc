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
	WalletId     string `db:"wallet_id"`
	Email        string `db:"email"`
	KeychainData string `db:"keychain_data"`
	Salt         string `db:"salt"`
	Verified     bool   `db:"verified"`
	//VerificationToken string  `db:"verification_token"`
}
