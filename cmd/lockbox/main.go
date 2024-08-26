/*
CLI tool для работы со списком паролей.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hazadus/go-lockbox"
)

var lockboxFileName = "~/.lockbox"

func main() {
	secret, exists := os.LookupEnv("GO_LOCKBOX_SECRET")
	if !exists {
		fmt.Fprintln(os.Stderr, "Не установлено значение переменной окружения GO_LOCKBOX_SECRET.")
		os.Exit(1)
	}
	if secretLen := len(secret); (secretLen != 16) && (secretLen != 24) && (secretLen != 32) {
		fmt.Fprintln(os.Stderr, "Значение переменной окружения GO_LOCKBOX_SECRET должно иметь длину 16, 24 или 32 байта.")
		os.Exit(1)
	}

	// Получаем имя файла из переменной окружения (при наличии)
	if envLockboxFileName := os.Getenv("GO_LOCKBOX_FILENAME"); envLockboxFileName != "" {
		lockboxFileName = envLockboxFileName
	}

	// Опеделяем флаги и получаем их значения
	addFlag := flag.String("add", "", "Добавить сервис <string>. Необходимо указать пароль -pwd <password>.")
	passwordFlag := flag.String("pwd", "", "Пароль <string>.")
	getFlag := flag.String("get", "", "Получить пароль от сервиса <string> в stdout.")
	deleteFlag := flag.String("del", "", "Удалить сервис <string> из списка.")
	listFlag := flag.Bool("list", false, "Вывести названия сохранённых в списке сервисов.")
	verboseFlag := flag.Bool("v", false, "Используется совместно с -list: вывод более подробной информации.")
	flag.Parse()

	recordList := &lockbox.List{}

	if err := recordList.Load(lockboxFileName, secret); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *addFlag != "":
		title := *addFlag
		password := *passwordFlag
		recordList.Add(title, password)

		// Должен быть обязательно указан флаг -pwd с паролем
		if *passwordFlag == "" {
			fmt.Fprintln(os.Stderr, "Не указан параметр -pwd <password>")
			os.Exit(1)
		}

		// Сохранить обновленный список в файл
		if err := recordList.Save(lockboxFileName, secret); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *getFlag != "":
		title := *getFlag
		password, _ := recordList.Get(title)
		fmt.Print(password)

	case *deleteFlag != "":
		title := *deleteFlag
		if err := recordList.Delete(title); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Сохранить обновленный список в файл
		if err := recordList.Save(lockboxFileName, secret); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *listFlag:
		for _, rec := range *recordList {
			if *verboseFlag {
				dateFormat := "Mon Jan 2 2006 15:04:05 MST"
				fmt.Println(fmt.Sprintf("%-16s Создан: %s, Изменён: %s", rec.Title, rec.CreatedAt.Format(dateFormat), rec.UpdatedAt.Format(dateFormat)))
			} else {
				fmt.Println(rec.Title)
			}
		}

	default:
		// Нет флагов или неверные флаги
		fmt.Fprintln(os.Stderr, "Неверные параметры")
		os.Exit(1)
	}
}
