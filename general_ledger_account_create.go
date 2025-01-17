// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type GeneralLedgerAccountCreate struct {

	// Unique code to identify the ledger account. Each code must start
	// with a letter or number. The following special characters are
	// allowed: `-_.,:`
	Code *string `json:"code,omitempty"`

	// Optional description.
	Description *string `json:"description,omitempty"`

	AccountType *string `json:"account_type,omitempty"`
}
