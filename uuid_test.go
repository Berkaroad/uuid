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

func Test_PutToBytes(t *testing.T) {
	uuid := New()
	buffer := make([]byte, 16)
	PutToBytes(buffer[0:16], uuid)
	uuid2 := LoadFromBytes(buffer)
	if uuid != uuid2 {
		t.Fail()
	}
}
