package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func getPwd(txt string) []byte {
	fmt.Print(txt)
	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func main() {

	pwd := getPwd("Inserire la password: ")
	hash := hashAndSalt(pwd)

	pwd2 := getPwd("Reinserire la password: ")
	pwdMatch := comparePasswords(hash, pwd2)
	fmt.Println("Passwords Match?", pwdMatch)

}
