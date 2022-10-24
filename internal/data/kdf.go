package data

type KDFQ interface {
	New() KDFQ

	Get() (*KDF, error)
	Select() ([]KDF, error)
	FilterByKDFVersion(version int) KDFQ
}

type KDF struct {
	Version   int     `db:"version"`
	Algorithm string  `db:"algorithm"`
	Bits      uint    `db:"bits"`
	N         uint    `db:"n"`
	R         uint    `db:"r"`
	P         uint    `db:"p"`
	Salt      *string `db:"salt"`
}
