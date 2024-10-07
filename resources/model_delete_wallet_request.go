/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DeleteWalletRequest struct {
	Key
	Attributes DeleteWalletRequestAttributes `json:"attributes"`
}
type DeleteWalletRequestResponse struct {
	Data     DeleteWalletRequest `json:"data"`
	Included Included            `json:"included"`
}

type DeleteWalletRequestListResponse struct {
	Data     []DeleteWalletRequest `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustDeleteWalletRequest - returns DeleteWalletRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDeleteWalletRequest(key Key) *DeleteWalletRequest {
	var deleteWalletRequest DeleteWalletRequest
	if c.tryFindEntry(key, &deleteWalletRequest) {
		return &deleteWalletRequest
	}
	return nil
}
