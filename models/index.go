package models

type index struct {
	publicKey string "json:'public_key'"
	manifest  string "json:'manifest'"
	blob      string "json:'blob'"
	signature string "json:'signature'"
	version   int    "json:'version'"
}
