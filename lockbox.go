/*
Package lockbox содержит логику для работы хранимыми записями.
*/
package lockbox

import "time"

// record представляет хранимую запись
type record struct {
	Title     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Lockbox представляет список хранимых записей
type Lockbox []record

// Add добавляет запись в Lockbox.
// Если запись с таким title уже существует, она будет обновлена.
func (l *Lockbox) Add(title string, password string) {
	// Сначала проверим, нет ли записи с таким title
	for i, rec := range *l {
		if rec.Title == title {
			(*l)[i].Password = password
			(*l)[i].UpdatedAt = time.Now()
			return
		}
	}

	timeCreated := time.Now()
	rec := record{
		Title:     title,
		Password:  password,
		CreatedAt: timeCreated,
		UpdatedAt: timeCreated,
	}
	*l = append(*l, rec)
}
