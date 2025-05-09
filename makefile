# Makefile für Go-Projekt

# Standard-Ziel: baue und teste alles
all: build test coverage

# Baue alle Module im Projekt
build:
	go build ./...

# Führe alle Tests aus
test:
	go test ./... -v

# Generiere Coverage-Datei und HTML-Bericht
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Zeige Coverage in der Konsole
coverage-report:
	go tool cover -func=coverage.out

# Entferne temporäre Dateien
clean:
	rm -f coverage.out coverage.html