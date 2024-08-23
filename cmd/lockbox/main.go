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

var lockboxFileName = ".lockbox.json"

func main() {
	addFlag := flag.String("add", "", "Название сервиса")
	passwordFlag := flag.String("pwd", "", "Пароль")
	flag.Parse()

	recordList := &lockbox.List{}

	if err := recordList.Load(lockboxFileName); err != nil {
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
		if err := recordList.Save(lockboxFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Нет флагов или неверные флаги
		fmt.Fprintln(os.Stderr, "Неверные параметры")
		os.Exit(1)
	}
}