package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

func GenerateWalletID(email, password, salt []byte, n, r, p, hashLength int) (string, error) {
	derivationToken := []byte("WALLET_ID")
	firstByte, _ := hex.DecodeString("01")

	walletSalt := append(firstByte, salt...)
	walletSalt = append(walletSalt, email...)
	walletSaltHashed := sha256.Sum256(walletSalt)

	masterKey, err := scrypt.Key(password, walletSaltHashed[:], n, r, p, hashLength)
	if err != nil {
		return "", err
	}

	walletID := hmac.New(sha256.New, masterKey)
	walletID.Write(derivationToken)

	return hex.EncodeToString(walletID.Sum(nil)), nil
}
