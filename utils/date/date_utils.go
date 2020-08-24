package date

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"

func GetNowString() string {
	return time.Now().Format(apiDateLayout)
}
