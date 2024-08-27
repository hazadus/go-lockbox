package lockbox_test

import (
	"testing"

	lockbox "github.com/hazadus/go-lockbox/internal/lockbox"
)

func TestAdd(t *testing.T) {
	lockbox := lockbox.List{}
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

// TestAddUpdates проверяет, что при добавлении
// записи с существующим Title, запись будет обновлена
func TestAddUpdates(t *testing.T) {
	lockbox := lockbox.List{}
	title := "service"
	password := "12345678"
	lockbox.Add(title, password)
	newPassword := "87654321"

	lockbox.Add(title, newPassword)

	if len(lockbox) != 1 {
		t.Fatalf("Size of lockbox should not increase.")
	}

	if lockbox[0].Title != title {
		t.Errorf("Expected %q, got %q instead.", title, lockbox[0].Title)
	}
	if lockbox[0].Password != newPassword {
		t.Errorf("Expected %q, got %q instead.", newPassword, lockbox[0].Password)
	}
	if lockbox[0].CreatedAt == lockbox[0].UpdatedAt {
		t.Errorf("UpdatedAt should not be equal to CreatedAt after update.")
	}
}

func TestGet(t *testing.T) {
	lockbox := lockbox.List{}
	title := "service"
	password := "12345678"
	lockbox.Add(title, password)

	receivedPassword, _ := lockbox.Get(title)

	if receivedPassword != password {
		t.Errorf("Expected %q, got %q instead.", password, receivedPassword)
	}
}

func TestGetFailsWithUnknownTitle(t *testing.T) {
	lockbox := lockbox.List{}
	title := "service"
	password := "12345678"
	lockbox.Add(title, password)

	receivedPassword, err := lockbox.Get("unknown")

	if err == nil {
		t.Errorf("Get must fail with unknown title.")
	}
	if receivedPassword != "" {
		t.Errorf("Expected '', got %q instead.", receivedPassword)
	}
}

func TestDelete(t *testing.T) {
	lockbox := lockbox.List{}
	title1 := "service1"
	password1 := "12345678"
	title2 := "service2"
	password2 := "87654321"
	lockbox.Add(title1, password1)
	lockbox.Add(title2, password2)

	lockbox.Delete(title1)

	if len(lockbox) != 1 {
		t.Fatalf("Size of lockbox should be 1 after deletion.")
	}
	if lockbox[0].Title != title2 {
		t.Errorf("Expected %q, got %q instead.", title2, lockbox[0].Title)
	}
	if lockbox[0].Password != password2 {
		t.Errorf("Expected %q, got %q instead.", password2, lockbox[0].Password)
	}
}
