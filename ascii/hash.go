package ascii

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(content []byte, hash string) bool {
	newhash := md5.New()
	newhash.Write([]byte(content))
	return hex.EncodeToString(newhash.Sum(nil)) == hash
}
