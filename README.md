# Team Go - WWI24SEA

Das Repository von der Gruppe GO der Modernen Programmierkonzepte WWI24SEA - DHBW. Das Repository besteht insgesamt aus fünf verschiedenen Teilaufgaben, welche wir alle im Unterricht thematisiert haben.

## Struktur

- `calculator/`: CLI-Taschenrechner mit Stack-basiertem Rechenmodell
- `datastructures/`: Eigene Implementierungen von Stack, Queue, LinkedList sowie Erweiterungen für Funktionale Programmierung
- `edsl/`: Domain-Spezifische Sprache (EDSL) für mathematische Ausdrücke und Vektorengrafiken
- `concurrency/`: Demonstration von Goroutines & Channels
- `password/`: Passwortwiederherstellung

## Voraussetzungen

Voraussetzung zur Nutzung der im Repository aufgeführten Aufgaben ist eine GO Version von 1.20 oder höher. 

## Build & Test

```bash
go build ./...
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
