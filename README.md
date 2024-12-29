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
## Примеры запросов
1. Для вычисления выражения отправьте POST-запрос:
   ```bash
   curl --location 'http://localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
   "expression": "2+2*2"
   }'
2. Ответ:
   ```Json
   {
   "result": "6.00"
   }
## Некорректное выражение (Ошибка 422
1. Если выражение содержит недопустимые символы, вы получите ответ:
   ```bash
   curl --location 'http://localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
   "expression": "2+abc"
   }'
2. Ответ:
   ```json
   {
   "error": "Expression is not valid"
    }
## Внутренняя ошибка сервера (Ошибка 500)
1. Если произошла неожиданная ошибка, сервер вернёт:
   ```json
   {
   "error": "Internal server error"
   }
