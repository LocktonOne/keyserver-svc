/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Kdf struct {
	Key
	Attributes KdfAttributes `json:"attributes"`
}
type KdfResponse struct {
	Data     Kdf      `json:"data"`
	Included Included `json:"included"`
}

type KdfListResponse struct {
	Data     []Kdf    `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustKdf - returns Kdf from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustKdf(key Key) *Kdf {
	var kDF Kdf
	if c.tryFindEntry(key, &kDF) {
		return &kDF
	}
	return nil
}
