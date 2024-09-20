package utils

import (
	"errors"
	"strconv"
	"time"
)

func TimeFromStrUnix(str string) (time.Time, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	if i < 1 {
		return time.Time{}, errors.New("wrong time")
	}

	return time.Unix(i, 0), nil
}

func ParseFloat32(str string) (float32, error) {
	i, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}

	return float32(i), nil
}
