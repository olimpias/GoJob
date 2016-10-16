package GoJob

import "testing"

func MockCreateNewAtomicBool() interface{}  {
	return NewAtomicBool();
}

func MockCreateNewAtomicBoolWithPtr() interface{}  {
	return NewAtomicBoolWithPtr();
}

func TestNewAtomicBool(t *testing.T) {
	_,ok := MockCreateNewAtomicBool().(AtomicBool);
	if !ok {
		t.Error("Without pointer expected");
	}
}

func TestNewAtomicBoolWithPtr(t *testing.T) {
	_,ok := MockCreateNewAtomicBoolWithPtr().(*AtomicBool);
	if !ok {
		t.Error("With pointer expected");
	}
}

func TestGetBoolValue(t *testing.T)  {
	boolValue := NewAtomicBool();
	if boolValue.getBoolValue() != false {
		t.Error("Get value is not working properly");
	}
}

func TestSetBoolValue(t *testing.T)  {
	boolValue := NewAtomicBool();
	boolValue.setBoolValue(true);
	if boolValue.getBoolValue() != true {
		t.Error("Set value is not working properly")
	}
}
