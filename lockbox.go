/*
Package lockbox содержит логику для работы хранимыми записями.
*/
package lockbox

import "time"

// record представляет хранимую запись
type record struct {
	Title string
	Password string
	CreatedAt   time.Time
}

// Lockbox представляет список хранимых записей
type Lockbox []record

// Add добавляет запись в Lockbox
func (l *Lockbox) Add(title string, password string) {
	rec := record{
		Title: title,
		Password: password,
		CreatedAt: time.Now(),
	}
	*l = append(*l, rec)
}