package str

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/qamarian-dtp/err"
)

func RandAnStr (length uint8) (string, error) {
	if length == 0 {
		return "", nil
	}

	randByte := make ([]byte, int (length))
	_, errX := rand.Read (randByte)
	if errX != nil {
		return "", err.New ("Unable to source random bytes for string generation.", nil, nil, errX)
	}
	randStr := base32.StdEncoding.EncodeToString (randByte)
	if len (randStr) < int (length) {
		return "", err.New ("Fatal error: string generated is less than length demanded.", nil, nil)
	}

	return randStr [0:length], nil
}
