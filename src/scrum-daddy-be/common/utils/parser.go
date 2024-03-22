package utils

import (
	"strconv"
)

func ParseToInt(intString string) (int, error) {
	id, err := strconv.Atoi(intString)
	if err != nil {
		return 0, err
	}
	return id, nil
}
