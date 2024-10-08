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
- ✅ Шифровать файл с паролями.

----

## Компиляция и настройка

```bash
make build
sudo ln -s ~/Projects/go-lockbox/bin/lockbox /usr/local/bin/lockbox
```

Установите значение переменной-секрета, применямого для шифрования и дешифрования файла с паролями:

```bash
# Значение должно иметь длину 16, 24 или 32 байта:
export GO_LOCKBOX_SECRET="1234567890123456"
```

Не теряйте это значение, без него вы не сможете расшифровать файл с паролями!

По умолчанию, зашифрованный список паролей хранится в файле `~/.lockbox`.

Для хранения информации в другом файле, установите переменную окружения `GO_LOCKBOX_FILENAME`, например:

```bash
export GO_LOCKBOX_FILENAME=~/.my_lockbox
```

Рекомендуется сохранить значения переменных окружения в файлах настройки оболочки (`~/.bashrc`, `~/.zshrc`, и т.п.).

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
