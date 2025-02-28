# Calc Service

## Описание  
Веб-сервис для распределённого вычисления арифметических выражений с использованием оркестратора и агентов. 

Проект состоит из трёх частей:
- **Оркестратор** — управляет задачами и распределяет их между агентами.
- **Агенты** — вычисляют арифметические выражения.
- **Веб-интерфейс** — позволяет пользователям отправлять запросы на вычисление выражений.

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



Запуск

Способ 1: Docker Compose (рекомендуется)

Для упрощённого запуска с помощью Docker Compose:

1. Запустите проект
   ```bash
   docker-compose up --build


2. После этого:

Оркестратор будет доступен по адресу: http://localhost:8080

Веб-интерфейс будет доступен по адресу: http://localhost:8081




Способ 2: Запуск вручную

Если у вас установлен Go, можно запустить сервер вручную:

1. Запустите оркестратор:
   ```bash
   go run orchestrator/main.go &


2. Запустите агента:
   ```bash
   go run agent/main.go &


3. Запустите веб-интерфейс:
   ```bash
   go run web/main.go &



Сервер будет доступен по адресу: http://localhost:8080/api/v1/calculate.

Примеры запросов

1. Вычисление выражения

Для вычисления выражения отправьте POST-запрос:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

Ответ:

{
  "result": "6.00"
}

2. Некорректное выражение (Ошибка 422)

Если выражение содержит недопустимые символы, вы получите ошибку:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+abc"
}'

Ответ:

{
  "error": "Expression is not valid"
}

3. Внутренняя ошибка сервера (Ошибка 500)

Если произойдёт неожиданная ошибка, сервер вернёт:

{
  "error": "Internal server error"
}

Документация

1. Оркестратор принимает задачи, передает их агентам и получает результаты.


2. Агенты вычисляют арифметические выражения и отправляют результаты обратно в оркестратор.


3. Веб-интерфейс позволяет пользователю отправлять выражения и получать ответы через HTTP.



Схема работы

[Web] → [Orchestrator] → [Agents] → [Orchestrator] → [Web]

---

### 2. **`docker-compose.yml`**

```yaml
version: '3.8'

services:
  orchestrator:
    build:
      context: ./orchestrator
    ports:
      - "8080:8080"
    networks:
      - calc-network
    depends_on:
      - agent

  agent:
    build:
      context: ./agent
    networks:
      - calc-network

  web:
    build:
      context: ./web
    ports:
      - "8081:8081"
    networks:
      - calc-network

networks:
  calc-network:
    driver: bridge
