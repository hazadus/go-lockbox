# 🔑 go-lockbox

CLI tool для удобного хранения и использования паролей.

Утилита создана мной в ходе [изучения языка Go](https://hazadus.github.io/knowledge/Languages/Go/Go) для личного использования. 

----

## Фичи проекта

- ✅ Хранение списка в файле в формате JSON.
- ✅ Настройка файла для хранения списка через env var.
- ✅ Тесты CLI.
- Флаги команды:
	- ✅ `-add <service> -pwd <password>`: добавить запись `service` с паролем `password`.
	- ✅ `-get <service>`: получить пароль от записи `service`.
	- ✅ `-delete <service>`: удалить запись `service` из списка.
	- ✅ `-list`: вывести перечень названий сервисов в списке.
      - ✅ `-v`: используется совместно с `-list` и выводить более подробную информацию.
- [ ] Шифровать данные в файле с использованием PIN-кода.

----

## Компиляция и настройка

```bash
make build
sudo ln -s ~/Projects/go-lockbox/lockbox /usr/local/bin/lockbox
```

По умолчанию, список хранится в файле `~/.lockbox.json`.

Для хранения информации в другом файле, установите переменную окружения `GO_LOCKBOX_FILENAME`, например:

```bash
export GO_LOCKBOX_FILENAME=~/.my_lockbox.json
```

## Как пользоваться

После приведённой выше настройки, можно пользоваться `lockbox` из любой директории в системе. Примеры:

```bash
# Добавить пароль от сервиса
lockbox -add amgold.ru -pwd 12345678

# Получить пароль и скопировать его в буфер обмена (MacOS)
lockbox -get amgold.ru | pbcopy

# Вывести перечень названий сервисов в списке
lockbox -list
lockbox -list -v

# Удалить запись
lockbox -del amgold.ru
```
