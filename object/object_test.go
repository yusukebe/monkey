package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with differenct content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	bool1 := &Boolean{Value: true}
	bool2 := &Boolean{Value: true}
	bool3 := &Boolean{Value: false}
	bool4 := &Boolean{Value: false}

	if bool1.HashKey() != bool2.HashKey() {
		t.Errorf("boolean with same value have different hash keys")
	}

	if bool3.HashKey() != bool4.HashKey() {
		t.Errorf("boolean with same value have different hash keys")
	}

	if bool1.HashKey() == bool3.HashKey() {
		t.Errorf("boolean with differenct value have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	integer1 := &Integer{Value: 1}
	integer2 := &Integer{Value: 1}
	integer3 := &Integer{Value: 2}

	if integer1.HashKey() != integer2.HashKey() {
		t.Errorf("integer with same value have different hash keys")
	}

	if integer1.HashKey() == integer3.HashKey() {
		t.Errorf("integer with differenct value have same hash keys")
	}
}
