package entitie

import (
	"errors"
	"strings"
)

type Contact struct {
	Phone int
	Email string
}

func (c *Contact) CheckIfFieldsAreEmpty() *[]error {
	var errs []error
	err := c.checkIfEmailIsEmpyt()
	if err != nil {
		errs = append(errs, err)
	}
	// verify email address
	// verify phone number
	if len(errs) != 0 {
		return &errs
	}
	return nil
}

func (c *Contact) checkIfEmailIsEmpyt() error {
	if strings.Trim(c.Email, " ") == "" {
		return errors.New("email field cannot is empty")
	}
	return nil
}
