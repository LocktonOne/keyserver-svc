/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VerifyWalletRequest struct {
	Key
	Attributes VerifyWalletRequestAttributes `json:"attributes"`
}
type VerifyWalletRequestResponse struct {
	Data     VerifyWalletRequest `json:"data"`
	Included Included            `json:"included"`
}

type VerifyWalletRequestListResponse struct {
	Data     []VerifyWalletRequest `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustVerifyWalletRequest - returns VerifyWalletRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVerifyWalletRequest(key Key) *VerifyWalletRequest {
	var verifyWalletRequest VerifyWalletRequest
	if c.tryFindEntry(key, &verifyWalletRequest) {
		return &verifyWalletRequest
	}
	return nil
}
