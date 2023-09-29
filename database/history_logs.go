package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type WaterLevelHistory struct {
	ID         int64   `json:"ID"`
	WaterLevel float64 `json:"WaterLevel"`
	Timestamp  int64   `json:"Timestamp"`
}

type WaterLevelHistoryList struct {
	WaterLevelHistory []*WaterLevelHistory `json:"WaterLevelData"`
}

func AddWaterLevelHistory(db sqlx.Ext, WaterLevel float64) error {

	_, err := db.Exec(`INSERT INTO HistoryLogs(
		WaterLevel,
		Timestamp
	)
	Values(?,?)`,
		WaterLevel,
		time.Now().Unix(),
	)
	if err != nil {
		return err
	}

	return nil
}

func GetWaterLevelHistory(db sqlx.Ext) (*WaterLevelHistoryList, error) {
	waterHistoryArray := make([]*WaterLevelHistory, 0)

	var ID int64
	var waterlvl float64
	var timestamp int64

	rows, err := db.Queryx(`SELECT ID, WaterLevel,Timestamp FROM HistoryLogs ORDER BY Timestamp DESC LIMIT 1800`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &waterlvl, &timestamp)
		if err != nil {
			return nil, err
		}
		waterHistoryArray = append(waterHistoryArray, &WaterLevelHistory{
			ID:         ID,
			WaterLevel: waterlvl,
			Timestamp:  timestamp,
		})
	}

	return &WaterLevelHistoryList{
		WaterLevelHistory: waterHistoryArray,
	}, nil
}
