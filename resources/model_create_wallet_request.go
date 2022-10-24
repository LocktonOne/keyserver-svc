/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateWalletRequest struct {
	Key
	Attributes CreateWalletRequestAttributes `json:"attributes"`
}
type CreateWalletRequestResponse struct {
	Data     CreateWalletRequest `json:"data"`
	Included Included            `json:"included"`
}

type CreateWalletRequestListResponse struct {
	Data     []CreateWalletRequest `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustCreateWalletRequest - returns CreateWalletRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateWalletRequest(key Key) *CreateWalletRequest {
	var createWalletRequest CreateWalletRequest
	if c.tryFindEntry(key, &createWalletRequest) {
		return &createWalletRequest
	}
	return nil
}
