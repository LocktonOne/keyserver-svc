/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// KDF parameters which should be used for derivation
type KdfAttributes struct {
	// kdf algorithm
	Algorithm string  `json:"algorithm"`
	Bits      int32   `json:"bits"`
	N         int32   `json:"n"`
	P         int32   `json:"p"`
	R         int32   `json:"r"`
	Salt      *string `json:"salt,omitempty"`
}
