package lockbox_test

import (
	"testing"

	lockbox "github.com/hazadus/go-lockbox"
)

func TestAdd(t *testing.T) {
	lockbox := lockbox.Lockbox{}
	title := "service"
	password := "12345678"

	lockbox.Add(title, password)

	if lockbox[0].Title != title {
		t.Errorf("Expected %q, got %q instead.", title, lockbox[0].Title)
	}
	if lockbox[0].Password != password {
		t.Errorf("Expected %q, got %q instead.", password, lockbox[0].Password)
	}
}