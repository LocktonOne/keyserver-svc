/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Wallet struct {
	Key
	Attributes WalletAttributes `json:"attributes"`
}
type WalletResponse struct {
	Data     Wallet   `json:"data"`
	Included Included `json:"included"`
}

type WalletListResponse struct {
	Data     []Wallet `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustWallet - returns Wallet from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustWallet(key Key) *Wallet {
	var wallet Wallet
	if c.tryFindEntry(key, &wallet) {
		return &wallet
	}
	return nil
}
