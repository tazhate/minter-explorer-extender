package models

import "math/big"

type Slash struct {
	ID          uint64     `json:"id"`
	CoinID      uint64     `json:"coin_id"`
	BlockID     uint64     `json:"block_id"`
	AddressID   uint64     `json:"address_id"`
	ValidatorID uint64     `json:"validator_id"`
	Amount      big.Int    `json:"amount" sql:"type:numeric(70)"`
	Coin        *Coin      //Relation has one to Coins
	Block       *Block     //Relation has one to Blocks
	Address     *Address   //Relation has one to Addresses
	Validator   *Validator //Relation has one to Validators
}
