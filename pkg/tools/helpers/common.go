package helpers

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StrToBool force convert string to boolean
func StrToBool(s string) bool {
	if s == "да" {
		return true
	}

	if v, err := strconv.ParseBool(s); err == nil {
		return v
	}

	return false
}

// StrToInt force convert string to int
func StrToInt(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return 0
}

// StrToInt64 force convert string to int64
func StrToInt64(s string) int64 {
	return int64(StrToInt(s))
}

// StrToUInt force convert string to uint
func StrToUInt64(s string) uint64 {
	return uint64(StrToInt(s))
}

// StrToFloat force convert string to float64
func StrToFloat(s string) float64 {
	if n, err := strconv.ParseFloat(s, 64); err == nil {
		return n
	}
	return float64(0.0)
}

// Abs return abs int value
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Md5Hash calculate md5 hash of string
func Md5Hash(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func TrimAllSpace(str string) string {
	withoutSpace := strings.Split(strings.ToLower(strings.TrimSpace(str)), " ")
	return strings.Join(withoutSpace, "")
}

// IsPathExists check is path exists
func IsPathExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
