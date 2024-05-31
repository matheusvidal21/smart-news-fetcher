package utils

import "time"

type date []byte

func ParseTime(d date) (time.Time, error) {
	data, err := time.Parse("2006-01-02 15:04:05", string(d))
	if err != nil {
		return time.Time{}, nil
	}
	return data, nil
}
