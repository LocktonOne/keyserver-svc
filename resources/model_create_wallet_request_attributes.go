/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateWalletRequestAttributes struct {
	// email address provided during the wallet creation
	Email string `json:"email"`
	// arbitrary client-provided string
	KeychainData string `json:"keychain_data"`
	// client-generated salt
	Salt string `json:"salt"`
	// unique identifier of the user account generated during the wallet creation
	WalletId string `json:"wallet_id"`
}
