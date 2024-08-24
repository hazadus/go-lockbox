# 🔑 go-lockbox

CLI tool для удобного хранения и использования паролей.

Утилита создана мной в ходе [изучения языка Go](https://hazadus.github.io/knowledge/Languages/Go/Go) для личного использования. 

----

## Фичи проекта

- ✅ Хранение списка в файле в формате JSON.
- ✅ Настройка файла для хранения списка через env var.
- ✅ Тесты CLI.
- Флаги команды:
	- ✅ `-add service -pwd password`: добавить запись `service` с паролем `password`.
	- ✅ `-get service`: получить пароль от записи `service`.
	- [ ] `-delete service`
	- [ ] `-list`
	- [ ] `-pin` (string)
- [ ] Шифровать данные в файле с использованием PIN-кода.

----

## Компиляция и настройка

```bash
make build
sudo ln -s ~/Projects/go-lockbox/lockbox /usr/local/bin/lockbox
```

## Как пользоваться

После приведённой выше настройки, можно пользоваться `lockbox` из любой директории в системе. Примеры:

```bash
# Добавить пароль от сервиса
lockbox -add amgold.ru -pwd 12345678

# Получить пароль и скопировать его в буфер обмена (MacOS)
lockbox -get amgold.ru | pbcopy
```
