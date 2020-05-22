package repo


import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"github.com/joho/godotenv"
	"os"
	"fmt"
)
func GetSaltForCustomer(customerId string) string {
	err := godotenv.Load(".env")
     fmt.Println(err)
	if err != nil {
		fmt.Println("Cannot find .env file")
		return "Cannot find .env file"
	}
	serverHash := os.Getenv("SERVER_HASH")
	fmt.Println("serverhash",serverHash,customerId)
    hasher := md5.New()
	hasher.Write([]byte(customerId))
	hexHasher := hex.EncodeToString(hasher.Sum(nil))
	data := strings.ToUpper(hexHasher)+ serverHash
	newSha256 := sha256.New()
	newSha256.Write([]byte(data))
    finalData := base64.StdEncoding.EncodeToString(newSha256.Sum(nil))
	fmt.Println("invalid salt",finalData)
	return finalData

}




