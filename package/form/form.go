package form

import (
	"net/http"
	"net/url"
	"strings"
)

type Form struct{
	url.Values
	Error errors
}


// Initialize a form structure
func New(data url.Values) *Form{
 return &Form{
	 data,
	 errors(map[string][]string{}),
 }

}

// check if this field exist or not
func (f *Form) Has(field string, r *http.Request) bool{
	x := r.Form.Get(field)
	if x == "" {
		f.Error.Add(field,"This field can not be empty.");
		return false
	}
	return true

}

// Validate if there's errors or not
func (f *Form) Valid() bool{
	return len(f.Error) == 0
}

// Required checks for required fields
func (f *Form) Required(fields ...string){
	for _, field := range fields{
		value := f.Get(field) 
		if strings.TrimSpace(value) == ""{
			f.Error.Add(field,"This field can not be empty.")
		}
	}
}
