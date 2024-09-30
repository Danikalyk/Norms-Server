package models

type CUTTING_INFO struct {
	RECORD_ID       int     `db:"record_id" json:"record_id"`
	MINCUTTINGSPEED float64 `db:"mincuttingspeed" json:"mincuttingspeed"`
	AVECUTTINGSPEED float64 `db:"avecuttingspeed" json:"avecuttingspeed"`
	MAXCUTTINGSPEED float64 `db:"maxcuttingspeed" json:"maxcuttingspeed"`
	INSERTIONTIME   float64 `db:"insertiontime" json:"insertiontime"`
	DATA_UPDATED    bool    `db:"data_updated" json:"data_updated"`
}
