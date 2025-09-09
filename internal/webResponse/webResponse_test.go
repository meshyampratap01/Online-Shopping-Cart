package webResponse

import "testing"

func TestNewErrorResponse(t *testing.T){
	expected:=&WebResponse{
		Code: 200,
		Message: "asdf",
	}

	got:=NewErrorResponse(200,"asdf")

	if *expected!=*got{
		t.Errorf("wanted: %v, got: %v",expected,got)
	}

	
}

func TestNewSuccessResponse(t *testing.T){
	expected:=&WebResponse{
		Code: 200,
		Message: "asdf",
		Data: "asdf",
	}

	got:=NewSuccessResponse(200,"asdf","asdf")

	if *expected!=*got{
		t.Errorf("wanted: %v, got: %v",expected,got)
	}

	
}