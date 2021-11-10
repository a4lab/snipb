package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MaxLenght(max int, field string) {

	value := f.Get(field)

	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > max {
		f.Errors.Add(field, fmt.Sprintf("Maximum Lenght allowed is %d", max))
	}

}

func (f *Form) AcceptableValues(field string, opts ...string) {

	value := f.Get(field)

	if value == "" {
		return

	}
	for _, opt := range opts {

		if value == opt {
			return
		}

	}
	f.Errors.Add(field, "This field is invalid")
}

func (f *Form) MinLenght(min int, field string) {

	value := f.Get(field)

	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < min {
		f.Errors.Add(field, fmt.Sprintf("Minimum Lenght allowed is %d", min))
	}

}

func (f *Form) MatchPattern(field string, regex *regexp.Regexp) {

	value := f.Get(field)

	if value == "" {
		return
	}
	if !regex.MatchString(value) {
		f.Errors.Add(field, "This Field Is Invalid")
	}

}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
