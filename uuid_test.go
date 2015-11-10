package uuid

import (
	"testing"
)

func Test_NewUUID(t *testing.T) {
	NewUUID().String()
}

func Test_ParseUUID(t *testing.T) {
	uuidString := NewUUID().String()
	if otherUUID, err := ParseUUID(uuidString); err != nil || uuidString != otherUUID.String() {
		t.Fail()
	}
}
