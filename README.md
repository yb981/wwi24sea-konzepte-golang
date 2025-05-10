# Team Go - WWI24SEA

## Struktur

- `calculator/`: CLI-Taschenrechner mit Stack-basiertem Rechenmodell
- `datastructures/`: Eigene Implementierungen von Stack, Queue, LinkedList sowie Erweiterungen für Funktionale Programmierung
- `edsl/`: Domain-Spezifische Sprache (EDSL) für mathematische Ausdrücke
- `concurrency/`: Demonstration von Goroutines & Channels
- `password/`: Passwortwiederherstellung

## Voraussetzungen

- Go >= 1.20

## Build & Test

```bash
go build ./...
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html