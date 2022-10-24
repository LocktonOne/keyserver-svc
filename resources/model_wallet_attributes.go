/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type WalletAttributes struct {
	// email address provided during the wallet creation
	Email string `json:"email"`
	// client-provided string derived from wallet keys
	KeychainData string `json:"keychain_data"`
	// client-generated salt
	Salt string `json:"salt"`
	// shows whether or not the wallet is verified (whether the user of a wallet has been verified via email link)
	Verified bool `json:"verified"`
	// unique identifier of the user account generated during the wallet creation
	WalletId string `json:"wallet_id"`
}
