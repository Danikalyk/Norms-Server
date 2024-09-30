package models

type Machine struct {
	MACHINE_ID   int    `db:"machine_id" json:"machine_id"`
	MACHINE_NAME string `db:"machine_name" json:"machine_name"`
}
