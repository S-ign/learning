package forms

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creats a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		return false
	}

	return true
}

// Valid checks true if no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required generates error messages for fields that are left blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks for min length to enter into a field
func (f *Form) MinLength(length int, fields ...string) map[string]bool {
	fieldStatus := make(map[string]bool)

	for i, field := range fields {
		x := f.Get(field)

		if len(x) < length {
			f.Errors.Add(field, fmt.Sprintf("This field must be at least %d charecters long", length))
			fieldStatus[field] = false

			if i == len(fields) {
				return fieldStatus
			}
		}
	}

	return fieldStatus
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email Address")
	}
}

// IsOnlyNumbers validates phone field, must be all numbers, must be len 10
func (f *Form) IsOnlyNumbers(field string) {
	_, err := strconv.Atoi(strings.TrimSpace(f.Get(field)))
	if err != nil {
		f.Errors.Add(field, "This field must contain only numbers")
	}
}
