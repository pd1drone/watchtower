package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type WaterLevelDataRealTime struct {
	WaterLevel string `json:"water_level"`
}

func GetWaterLevelRealTimeData(db sqlx.Ext) (*WaterLevelDataRealTime, error) {

	var ID int64
	var waterlvl float64
	var timestamp int64

	rows, err := db.Queryx(`SELECT ID, WaterLevel,Timestamp FROM HistoryLogs ORDER BY Timestamp DESC LIMIT 1`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &waterlvl, &timestamp)
		if err != nil {
			return nil, err
		}
	}

	return &WaterLevelDataRealTime{
		WaterLevel: fmt.Sprintf("%f", waterlvl),
	}, nil
}
