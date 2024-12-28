# Calc Service

## Описание  
Веб-сервис для вычисления арифметических выражений.  

## Установка  
1. Клонируйте репозиторий:  
   ```bash
   git clone https://github.com/qwix7/calc-service.git
2. Перейдите в директорию проекта:
   ```bash
   cd calc-service
3. Убедитесь, что проект собран корректно:
   ```bash
   go mod tidy
## Запуск
1. Запустите сервер:
   ```bash
   go run ./cmd/calc_service/...
2. Сервер будет доступен в локальном адресе:
   ```bash
   http://localhost:8080/api/v1/calculate
