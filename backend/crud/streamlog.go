package crud

import (
	"time"

	"github.com/Grvzard/biligank-v2/backend/config"
)

type StreamlogRow struct {
	Title      string `json:"title,omitempty"`
	CTs        uint32 `json:"c_ts,omitempty"`
	LastUpdate uint32 `json:"last_update,omitempty"`
	WatchedNum uint64 `json:"watched_num,omitempty"`
	AreaName   string `json:"area_name,omitempty"`
	Cover      string `json:"cover,omitempty"`
}

func FetchStreamLog(table string, uid int64, stop *time.Time) []StreamlogRow {
	results := []StreamlogRow{}
	if stop == nil {
		Sqldb.Table(table).Where("uid = ?", uid).Order("_id DESC").Find(&results)
	} else {
		stopTstamp := uint32(stop.Unix())
		Sqldb.Table(table).Where("uid = ?", uid).Where("last_update >= ?", stopTstamp).Order("_id DESC").Find(&results)
	}
	return results
}

var beijing = time.FixedZone("UTC+8", 8*60*60)

// `from` as near, `to` as far
func StreamLogByTstampSlice(uid int64, from uint32, to uint32) []StreamlogRow {
	results := []StreamlogRow{}
	from_ := time.Unix(int64(from), 0).In(beijing)
	to_ := time.Unix(int64(to), 0).In(beijing)
	step := 1
	// recent data are divided into tables by day (in tables like 2024_02_01)
	// data prior to Stage1 are archived monthly (in tables like 2023_10)
	for to_.Before(from_) {
		if from_.Before(config.GlobalConfig.Stage1) {
			step = 2
		}
		if step == 1 {
			tablename := from_.Format("2006_01_02")
			from_ = from_.AddDate(0, 0, -1)
			results = append(results, FetchStreamLog(tablename, uid, nil)...)
		} else if step == 2 {
			tablename := from_.Format("2006_01")
			from_ = from_.AddDate(0, -1, 0)
			stop := to_
			if from_.After(to_) {
				stop = from_
			}
			results = append(results, FetchStreamLog(tablename, uid, &stop)...)
		}
	}
	return results
}
