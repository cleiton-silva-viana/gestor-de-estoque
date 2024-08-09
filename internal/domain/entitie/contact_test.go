package entitie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var corporation = Contact{
	Phone: 2140028922,
	Email: "corporation@email.com",
}

func Test_CheckIfFieldsAreEmpyt_AllFieldsAreOk(t *testing.T) {
	result := corporation.CheckIfFieldsAreEmpty()
	assert.Nil(t, result, "The object can't have return error")
}

func Test_CheckIfFieldsAreEmpty_WithEmailEmpty(t *testing.T) {
	corporation.Email = "                 "

	result := corporation.CheckIfFieldsAreEmpty()

	assert.NotNil(t, result, "Error expected - email address is empty")
}