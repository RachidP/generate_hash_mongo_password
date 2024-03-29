package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	for {

		// Enter a password and generate a salted hash
		pwd := getPwd()
		hash := hashAndSalt(pwd)
		fmt.Println("plaintext  Password", string(pwd))
		fmt.Println("Hashed Password", hash)

		//compare password
		//fmt.Printf("\n\ncheck password\n")
		//fmt.Println("Enter the same password again and compare it with the first password entered")
		//pwd2 := getPwd()
		//pwdMatch := comparePasswords(hash, pwd2)
		//fmt.Println("Passwords Match?", pwdMatch)
	}
}

//getPwd get plaintext password from command line
func getPwd() []byte {
	// Prompt the user to enter a password
	fmt.Println("Enter a password")
	// Variable to store the users input
	var pwd string
	// Read the users input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)
}

//hashAndSalt hash the password
func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

//CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns nil on success, or an error on failure
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
