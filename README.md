# Team Go - WWI24SEA

Das Repository von der Gruppe GO der Modernen Programmierkonzepte WWI24SEA - DHBW. Das Repository besteht insgesamt aus fünf verschiedenen Teilaufgaben, welche wir alle im Unterricht thematisiert haben.

## Struktur

- `calculator/`: CLI-Taschenrechner mit Stack-basiertem Rechenmodell
- `datastructures/`: Eigene Implementierungen von Stack, Queue, LinkedList sowie Erweiterungen für Funktionale Programmierung
- `edsl/`: Domain-Spezifische Sprache (EDSL) für mathematische Ausdrücke und Vektorengrafiken
- `concurrency/`: Demonstration von Goroutines & Channels
- `password/`: Passwortwiederherstellung

## Tests

Für alle im Repository enthaltenden Projekte wurden Tests geschrieben und die Testabdeckung beträgt jeweils zwischen 70% bis 100%. Der Hauptgrund wieso die Testabdeckung nicht 100% liegt daran, dass wir die main Methoden nicht getestet haben, da sie keine logische Funtionalität enthalten. Weiter wurden nicht alle einfachen Consolen Ausgaben getest. Die Demos haben 0% Testcoverage da sie nur Anwendungsbeispiele sind. Ansonsten wurden alle Zeilen des Codes getest.
Es gibt jedoch immer noch Optimierungspotential bei den Tests. Die Testdateien sind teilweise etwas unstrukturiert und ungeordnet, da jeder etwas zu den Tests beigetragen hat. Dadurch enstehen auch Redunanz bei den Tests und man könne bei allen Testdateien noch Refactoring betreiben.

## Voraussetzungen

Voraussetzung zur Nutzung der im Repository aufgeführten Aufgaben ist eine GO Version von 1.20 oder höher. 

## Build & Test

```bash
go build ./...
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
