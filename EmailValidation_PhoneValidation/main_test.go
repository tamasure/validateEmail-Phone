package main1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidation(t *testing.T) {
	testEmail := EmailValidation("pratamamrandy@gmail.com")
	assert.Equal(t, testEmail, true, "Email Tidak Valid")

	testEmail1 := EmailValidation("PratamaMRandy@gmail.com")
	assert.Equal(t, testEmail1, true, "Email Tidak Valid")

	testEmail2 := EmailValidation("@gmail.com")
	assert.Equal(t, testEmail2, false, "Email Tidak Valid")

	testEmail3 := EmailValidation("pratamamrandy@gmail")
	assert.Equal(t, testEmail3, false, "Email Tidak Valid")

}

var (
	actual    = "6285954648112"
	notActual = ""
)

func TestPhoneValidation(t *testing.T) {
	testPhone1 := PhoneValidation("6285954648112")
	assert.Equal(t, testPhone1, actual, "Error")

	testPhone2 := PhoneValidation("085954648112")
	assert.Equal(t, testPhone2, actual, "Error")

	testPhone3 := PhoneValidation("085")
	assert.Equal(t, testPhone3, notActual, "Error")

	testPhone4 := PhoneValidation("085aassdcvxssdsd")
	assert.Equal(t, testPhone4, notActual, "Error")

	testPhone5 := PhoneValidation("0853324342325254525")
	assert.Equal(t, testPhone5, notActual, "Error")

}
