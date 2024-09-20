package util

import "strconv"

func ConvertToNumber(value string) (int64, error) {
	valueNumeric, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return int64(valueNumeric), nil
}
