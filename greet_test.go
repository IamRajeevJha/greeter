package greeter

import "testing"

func TestFunction(t *testing.T) {
	res := Greet("rajeev")
	expected := "Hello rajeev"
	if res != expected {
		t.Errorf("Got =  %s,  want =  %s", res, expected)
	}

}
