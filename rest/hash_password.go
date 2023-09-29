package rest

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5HashPassword(pass string) string {

	// Create an MD5 hash object
	hash := md5.New()

	// Write the input string to the hash object
	hash.Write([]byte(pass))

	// Get the final hash sum as a byte slice
	hashSum := hash.Sum(nil)

	// Convert the hash sum to a hexadecimal string
	hashString := hex.EncodeToString(hashSum)

	return hashString

}
