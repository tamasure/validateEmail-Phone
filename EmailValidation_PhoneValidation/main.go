package main1

import (
	"regexp"

	"github.com/dongri/phonenumber"
)

func PhoneValidation(phone string) string {
	number := phonenumber.Parse(phone, "ID")
	return number
}

func EmailValidation(email string) bool {
	emailRegexp := regexp.MustCompile(`^[a-z0-9A-Z._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegexp.MatchString(email)
}

// func main() {

// 	ph := PhoneValidation("628121")
// 	fmt.Println(ph)

// 	fmt.Println(EmailValidation("xxxxx"))
// 	fmt.Println(EmailValidation("@gmail.com"))
// 	fmt.Println(EmailValidation("pratamamrandy@gmail"))
// 	fmt.Println(EmailValidation("prata"))
// 	fmt.Println(PhoneValidation("6281286163232"))
// 	fmt.Println(PhoneValidation("081286163232"))
// 	fmt.Println(PhoneValidation("+628128616323222222"))

// }
