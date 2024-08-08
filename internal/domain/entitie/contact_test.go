package entitie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Corporation = Contact{
	Phone: 2140028922,
	Email: "corporation@email.com",
}

func Test_CreateContact_AllInputsAreOk(t *testing.T) {
	expected := nil

	result := Corporation.CreateContact()

	assert.Nil(t, expected, "The object can't have return error")
}
