package utils

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// ErrorCheck quick function to check for an error and, optionally terminate the program
func ErrorCheck(err error, where string, kill bool) {
	if err != nil {
		if kill {
			log.WithError(err).Fatalln("Script Terminated")
		} else {
			log.WithError(err).Warnf("@ %s\n", where)
		}
	}
}

// HexToUInt convert hex to int
func HexToUInt(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

// GetSha1Hash generate sha1 hash from interface{}
func GetSha1Hash(payload interface{}) string {

	out, err := json.Marshal(payload)
	if err != nil {
		log.Error(err)
		return ""
	}

	algorithm := sha1.New()
	algorithm.Write(out)
	return fmt.Sprintf("%x", algorithm.Sum(nil))
}
