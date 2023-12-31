package crud

import (
	"time"
)

type StreamlogRow struct {
	Title      string `json:"title,omitempty"`
	CTs        uint32 `json:"c_ts,omitempty"`
	LastUpdate uint32 `json:"last_update,omitempty"`
	WatchedNum uint64 `json:"watched_num,omitempty"`
	AreaName   string `json:"area_name,omitempty"`
	Cover      string `json:"cover,omitempty"`
}

func FetchStreamLog(table string, uid int64) []StreamlogRow {
	var results []StreamlogRow
	Sqldb.Table(table).Where("uid = ?", uid).Order("_id DESC").Find(&results)
	if len(results) != 0 {
		return results
	}
	return []StreamlogRow{}
}

func StreamLogByTstamp(ts uint32, uid int64) []StreamlogRow {
	return FetchStreamLog(time.Unix(int64(ts+8*3600), 0).Format("2006_01_02"), uid)
}
