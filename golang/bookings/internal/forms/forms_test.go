package forms

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/testing", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/testing", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/testing", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("name", "")
	form := New(postedData)

	has := form.Has("name")
	if has {
		t.Error(fmt.Sprintf("got %t when shoud have got %t", has, !has))
	}

}

func TestForm_MinLength(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("name", "john")
	form := New(postedData)

	form.MinLength(100, "john", "cat", "elephant")
	if form.Valid() {
		t.Error("shows minlength of 100 met when field is shorter")
	}
}

func TestForm_IsEmail(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("email", "john")
	form := New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form passed validation with invalid email form should be invalid - no @ in string")
	}

	postedData = url.Values{}
	postedData.Add("email", "")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form passed validation with invalid email form should be invalid - field is blank")
	}

	postedData = url.Values{}
	postedData.Add("email", "john@")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form passed validation with invalid email form should be invalid - no domain")
	}

	postedData = url.Values{}
	postedData.Add("email", "john@john")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form passed validation with invalid email form should be invalid - no tld")
	}

	postedData = url.Values{}
	postedData.Add("email", "john@john.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form failed validation with valid email form should be valid")
	}
}

func TestForm_IsOnlyNumbers(t *testing.T) {

	postedData := url.Values{}
	postedData.Add("phone", "john")
	form := New(postedData)
	form.IsOnlyNumbers("phone")
	if form.Valid() {
		t.Error("form passed validation with invalid number - no numbers")
	}

	isError := form.Errors.Get("phone")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("phone", "")
	form = New(postedData)
	form.IsOnlyNumbers("phone")
	if form.Valid() {
		t.Error("form passed validation with invalid number - blank")
	}

	postedData = url.Values{}
	postedData.Add("phone", "j0hn")
	form = New(postedData)
	form.IsOnlyNumbers("phone")
	if form.Valid() {
		t.Error("form passed validation with invalid number - letters and numbers")
	}

	postedData = url.Values{}
	postedData.Add("phone", "1234567890")
	form = New(postedData)
	form.IsOnlyNumbers("phone")
	if !form.Valid() {
		t.Error("form failed validation with valid number - should be valid")
	}

	isError = form.Errors.Get("phone")
	if isError != "" {
		t.Error("should not have an error but got one")
	}

}
