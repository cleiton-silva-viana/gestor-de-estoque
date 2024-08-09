package entitie

type Contact struct {
	Phone int
	Email string
}

func (c *Contact) CheckIfFieldsAreEmpty() *[]error {
	var errs []error
	// verify email address
	// verify phone number
	if len(errs) != 0 {
		return &errs
	}
	return nil
}
