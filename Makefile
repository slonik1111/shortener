# Имя бинарника
BINARY = shortenerBinary

# Имя модуля
MODULE = github.com/slonik1111/shortener

# Основная цель — запуск сервера
.PHONY: run
run: build
	./$(BINARY)

# Сборка бинарника
.PHONY: build
build: mod
	go build -o $(BINARY) main.go

# Создание go.mod, если его нет, и подтягивание зависимостей
.PHONY: mod
mod:
	@if [ ! -f go.mod ]; then \
		go mod init $(MODULE); \
	fi
	go mod tidy

# Очистка бинарников и обновление зависимостей
.PHONY: clean
clean:
	rm -f $(BINARY)

# Проверка кода на ошибки
.PHONY: lint
lint:
	gofmt -l -w .
	go vet ./...

# Тесты (если будут)
.PHONY: test
test:
	go test ./...
