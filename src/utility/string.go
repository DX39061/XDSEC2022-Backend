package utility

import (
	"fmt"
	"strconv"
)

// GetSerialNumber format uint into a serial number string.
func GetSerialNumber(id uint) string {
	return fmt.Sprintf("CER%08X", id)
}

// AtoU convert string to uint.
func AtoU(str string) (uint, error) {
	i, err := strconv.Atoi(str)
	return uint(i), err
}
