Wilkommen zu den Go Übungen!  
----
[**Link** zu VM Codespaces by GitHub](URhttps://github.com/codespaces/new/wannas234/GoUebungenL)

**Einleitung:**

>═══ 🚀 Basics 🚀 ═══

Das fmt-Paket wird genutzt, um etwas auf dern Konsole auszugeben.

Das time-Paket wird benutzt, um z.B. Pausen einzubauen oder mit Zeiten zu arbeiten.

time.Sleep sagt dem Programm, dass es für eine gewisse Zeit warten soll. (Ausführung wird pausiert)

func main() ist die Hauptfunktion, ohne diese Funktion läuft nichts.

>═══ 🚀 GoRoutinen 🚀 ═══

*func (Funktions Name)* ist eine eigene Funktion die definiert werden muss.

Ruft man die bereits definierte Funktion in der main() mit dem Schlüsselwort **go** auf, wird sie als Go Routine gestartet, (Die Funktion wird im Hintergrund ausgeführt, während main() weiterläuft.)

>═══ 🚀 Channels 🚀 ═══

Ein Channel ist ein spezieller Typ in Go, über den Goroutinen Daten austauschen können. 

->Channel erstellen: **ch := make(chan int)**, 

->Wert in einen Channel legen: **ch <- 42**, 

->Wert aus dem Channel empfangen: **Empfangen := <-ch**)
                                    
------------------------------------------------------------------------------

**Erste Übungseinheit - Aufgabenstellungen:**

Aufgabe 1.1 Erstelle einen "Hallo, Welt!" Ausgabe im Terminal.

Aufgabe 1.2: Kombiniere time.Sleep mit Println um nach einer Konsolen Ausgabe, die nächste Ausgabe erst nach 3 Sekunden starten zu lassen.

Aufgabe 1.3 Erstelle eine Go Routine, welche in der main Methode aufgerufen, und im Hintergrund ausgeführt wird.

------------------------------------------------------------------------------

**Zweite Übungseinheit - Aufgabenstellungen:**

Aufgabe 2.1: Erstelle ein Programm, das zwei Zahlen multipliziert und das Ergebnis durch eine dritte Zahl teilt.

Aufgabe 2.2: Führe Addition und Subtraktion mit Variablen a und b durch

Aufgabe 2.3: Erstelle ein Programm, welches die ersten 6 Fibonacci-Zahlen ausgibt.

------------------------------------------------------------------------------

**Dritte Übungseinheit - Aufgabenstellungen:**

Aufgabe 3.1: Erstelle einen Channel

Aufgabe 3.2: Erstelle mithilfe von Go Routinen und Channels einen Dialog zwischen Lisa und Manfred

Aufgabe 3.3: GO Channels Aufgabe für Mega Ultra Profis 😎.

Aufgabenstellung: Erstelle ein ausführbares GO-Projekt mit einer Funktion namens "boring", welche eine Nachricht als Funktionsargument annimmt und diese permanent
mit einem Delay zwischen 1 - 1000ms an einen Kanal sendet. Die entsandten Nachrichten sollen in der Main-Funktion abgefangen und in der Konsole ausgegeben werden.
Nach der fünften Nachricht wird in der Konsole "Mir reichts. Du bist langweilig." ausgegeben und das Programm beendet. **Tipp: Benutzt die Pakete "fmt", "time" und "math/rand"**

![Rpixel](https://github.com/user-attachments/assets/934acc6e-1eff-4df1-bbbc-ff4c40c7ed49)


