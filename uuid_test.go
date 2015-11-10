package uuid

import (
	"testing"
)

func Test_New(t *testing.T) {
	New().String()
}

func Test_Parse(t *testing.T) {
	uuidString := New().String()
	if otherUUID, err := Parse(uuidString); err != nil || uuidString != otherUUID.String() {
		t.Fail()
	}
}

func Test_IsEmpty(t *testing.T) {
	uuid := New()
	uuid2 := UUID{}
	if IsEmpty(uuid) {
		t.Fail()
	}
	if !IsEmpty(uuid2) {
		t.Fail()
	}
}
