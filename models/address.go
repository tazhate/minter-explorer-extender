package models

import "time"

type Address struct {
	ID                  uint64                `json:"id"`
	Address             string                `json:"address" sql:",unique; type:varchar(64)"`
	UpdatedAt           time.Time             `json:"updated_at"`
	Balances            []*Balance            `json:"balances"`             //relation has many to Balances
	Rewards             []*Reward             `json:"rewards"`              //relation has many to Rewards
	Slashes             []*Slash              `json:"slashes"`              //relation has many to Slashes
	Transactions        []*Transaction        `json:"transactions"`         //relation has many to Transactions
	InvalidTransactions []*InvalidTransaction `json:"invalid_transactions"` //relation has many to InvalidTransactions
}

// Return address with prefix
func (a *Address) GetAddress() string {
	return `Mx` + a.Address
}
