package utils

import "os"

const (
	DateTimeLayout  string = "2006-01-02 15:04:05"
	DateLayout      string = "2006-01-02"
	TimeLayout      string = "15:04:05"
	TimeShortLayout string = "15:04"
)

func GetEnv(key string, defValue string) string {
	if val := os.Getenv(key); val == "" {
		return defValue
	} else {
		return val
	}
}
