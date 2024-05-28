package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetRangeTimeData(start, end int64) string {
	yearStr := strconv.Itoa(time.Now().Year()) + "-"
	startStr := time.Unix(start, 0).Format("2006-01-02 15:04")
	endStr := time.Unix(end, 0).Format("2006-01-02 15:04")
	suffix10S := startStr[0:11]
	suffix10E := endStr[0:11]
	if suffix10S == suffix10E {
		// 2006-01-02 15:00~23:00
		return strings.ReplaceAll(fmt.Sprintf("%s %s~%s", suffix10S, startStr[11:], endStr[11:]), yearStr, "")
	}
	// 年份一样 01-02 15 00~23:00
	if startStr[0:4] == endStr[0:4] {
		return strings.ReplaceAll(fmt.Sprintf("%s %s~%s", startStr[0:5], startStr[5:], endStr[5:]), yearStr, "")
	}

	return strings.ReplaceAll(fmt.Sprintf("%s ~ %s", startStr, endStr), yearStr, "")
}

func GetTimeDate(timeInt int64) (res string) {

	res = time.Unix(timeInt, 0).Format("2006-01-02 15:04")
	res = strings.ReplaceAll(res, fmt.Sprintf("%d-", time.Now().Year()), "")
	return
}
