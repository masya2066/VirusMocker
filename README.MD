# Virus Mocker API

## Описание

Этот API предоставляет MOCK сервис для сканирования файлов.
Он способен работать в многопоточном режиме, полностью эмулируя поведение реальной системы СЗИ.

### Требования:
- Go 1.21
- PostgreSQL >= 14

### Конфигурация:
Установите конфигурацию в `configs/config.yml`

```yaml
server:
  port: 8080 # порт, по которому будет доступен сервис
database:
  host: localhost # адрес базы данных
  port: 5432 # порт базы данных
  username: user # имя пользователя базы данных
  password: password # пароль пользователя базы данных
  name: db # имя базы данных
```

### Запуск:
```sh
git clone https://github.com/your-repo/virus_mocker.git
cd virus_mocker
go mod tidy
go run main.go
```

#### Kata:
1. Отправка файлов в KATA:
```sh
POST /api_v1/kata/scanner/v1 
```
2. Получение списка файлов в KATA со статусом:
```sh
GET /api_v1/kata/scanner/v1/state
```
3. Удаление файла из KATA:
```sh
DELETE /api_v1/kata/scanner/v1/:scanId
```
