package helper

import "time"

func DataFormatada(data time.Time) string {
	return data.Format("02/01/2006 15:04:05")
}
