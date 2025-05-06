import Foundation

func showMainScreen() {
    print("\n--- Robocup Junior Vöhringen ---")
    print("Willkommen beim Robocup Junior Vöhringen!")
    print("Wähle einen Menüpunkt:")
    print("1. Home")
    print("2. Spielplan")
    print("3. Infos")
    print("4. Links")
    print("5. Einstellungen")
    print("6. Beenden")
}

var running = true
while running {
    showMainScreen()
    if let choice = readLine() {
        switch choice {
        case "1":
            print("\n[Home] Dies ist der Startscreen.")
        case "2":
            print("\n[Spielplan] (noch leer)")
        case "3":
            print("\n[Infos] (noch leer)")
        case "4":
            print("\n[Links] (noch leer)")
        case "5":
            print("\n[Einstellungen] (noch leer)")
        case "6":
            print("Programm wird beendet. Auf Wiedersehen!")
            running = false
        default:
            print("Ungültige Eingabe. Bitte wähle eine Zahl von 1 bis 6.")
        }
    }
}