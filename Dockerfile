# 1. Указываем базовый образ с Go
FROM golang:1.21-alpine

# 2. Создаём рабочую директорию в контейнере
WORKDIR /app

# 3. Копируем файлы зависимостей и устанавливаем их
COPY go.mod go.sum ./
RUN go mod download

# 4. Копируем исходный код проекта
COPY . .

# 5. Собираем бинарник
RUN go build -o shortenerBinary main.go

# 6. Команда запуска контейнера
CMD ["./shortenerBinary"]
