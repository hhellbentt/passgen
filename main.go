package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	password := ""
	fmt.Println("Введите пароль, чтобы получить hash")
	fmt.Scan(&password)
	hash := md5.Sum([]byte(password))
	fmt.Println(hex.EncodeToString(hash[:]))

}
