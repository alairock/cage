package cage

import (
	"strings"
	"testing"
)

func TestGenUid(t *testing.T) {
	prefix := "bob"
	uid := GenUID(prefix)
	if !strings.Contains(uid, prefix) {
		t.Fatal("uid does not contain prefix")
	}
}
