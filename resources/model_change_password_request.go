/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChangePasswordRequest struct {
	Key
	Attributes ChangePasswordRequestAttributes `json:"attributes"`
}
type ChangePasswordRequestResponse struct {
	Data     ChangePasswordRequest `json:"data"`
	Included Included              `json:"included"`
}

type ChangePasswordRequestListResponse struct {
	Data     []ChangePasswordRequest `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustChangePasswordRequest - returns ChangePasswordRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustChangePasswordRequest(key Key) *ChangePasswordRequest {
	var changePasswordRequest ChangePasswordRequest
	if c.tryFindEntry(key, &changePasswordRequest) {
		return &changePasswordRequest
	}
	return nil
}
