package validators

import (
	"strings"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	err := ValidateEmail("shyam@gmail.com")
	if err != nil {
		t.Error("no error wanted got error: ", err)
	}

	err = ValidateEmail("")
	if err == nil {
		t.Error("wanted error got no error")
	}
}

func TestValidatePassword(t *testing.T) {
	err := ValidatePassword("123")
	if err == nil {
		t.Error("wanted error got no error")
	}

	err = ValidatePassword("password@123")
	if err != nil {
		t.Error("wanted no error got error")
	}
}

func TestValidateName(t *testing.T) {
	err := ValidateName("")
	if err == nil {
		t.Error("wanted error got no error")
	}

	name := strings.Repeat("-", 1011)
	err = ValidateName(name)
	if err == nil {
		t.Error("wanted error got no error")
	}

	err=ValidateName("shyam")
	if err!=nil{
		t.Error("wanted no error got error")
	}
}

// func TestValidateJWT(t *testing.T){
// 	_,err:=ValidateJWT("")
// 	if err!=nil{
// 		t.Error("wanted no error , got error")
// 	}
// }

func TestValidateCoupon(t *testing.T){
	err:=ValidateCoupon("",20)
	if err==nil{
		t.Error("wanted error got no error")
	}

	err=ValidateCoupon("shyam",0)
	if err==nil{
		t.Error("wanted error got no error")
	}

	err=ValidateCoupon("shyam",78)
	if err!=nil{
		t.Error("wanted no error go error")
	}
}