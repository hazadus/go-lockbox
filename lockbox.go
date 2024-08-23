/*
Package lockbox содержит логику для работы хранимыми записями.
*/
package lockbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

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

// Get возвращает пароль записи с указанным title.
func (l *Lockbox) Get(title string) (string, error) {
	for _, rec := range *l {
		if rec.Title == title {
			return rec.Password, nil
		}
	}
	return "", fmt.Errorf("Item '%s' does not exist", title)
}

// Save method сохраняет Lockbox в формате JSON в
// указанном файле.
func (l *Lockbox) Save(filename string) error {
	jsonList, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonList, 0644)
}

// Load загружает содержимое файла в формате JSON
// в Lockbox.
func (l *Lockbox) Load(filename string) error {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(fileContent) == 0 {
		return nil
	}

	return json.Unmarshal(fileContent, l)
}
