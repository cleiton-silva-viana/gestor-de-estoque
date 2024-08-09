package entitie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Corporation = Contact{
	Phone: 2140028922,
	Email: "corporation@email.com",
}

func Test_CheckIfFieldsAreEmpyt_AllFieldsAreOk(t *testing.T) {
	result := Corporation.CheckIfFieldsAreEmpty()
	assert.Nil(t, result, "The object can't have return error")
}
