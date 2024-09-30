package models

type Gases struct {
	GAS_ID   int    `db:"gas_id" json:"gas_id"`
	GAS_NAME string `db:"gas_name" json:"gas_name"`
}
