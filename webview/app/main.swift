// SwiftUI App mit seitlichem Menü (Side Menu)
import SwiftUI

@main
struct RobocupJuniorApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}

struct ContentView: View {
    @State private var showMenu = false

    var body: some View {
        NavigationView {
            ZStack {
                MainView()
                    .navigationBarTitle("Robocup Junior Vöhringen", displayMode: .inline)
                    .navigationBarItems(leading: Button(action: {
                        withAnimation {
                            showMenu.toggle()
                        }
                    }) {
                        Image(systemName: "line.horizontal.3")
                            .imageScale(.large)
                    })

                if showMenu {
                    SideMenu(showMenu: $showMenu)
                        .transition(.move(edge: .leading))
                }
            }
        }
    }
}

struct MainView: View {
    var body: some View {
        VStack {
            Spacer()
            Text("Willkommen beim Robocup Junior Vöhringen")
                .font(.title2)
                .padding()
                .background(Color.gray.opacity(0.2))
                .cornerRadius(10)
            Spacer()
        }
    }
}

struct SideMenu: View {
    @Binding var showMenu: Bool

    var body: some View {
        VStack(alignment: .leading) {
            NavigationLink(destination: MainView()) {
                Text("Home")
                    .padding()
            }
            NavigationLink(destination: Text("Spielplan Page")) {
                Text("Spielplan")
                    .padding()
            }
            NavigationLink(destination: Text("Infos Page")) {
                Text("Infos")
                    .padding()
            }
            NavigationLink(destination: Text("Links Page")) {
                Text("Links")
                    .padding()
            }
            NavigationLink(destination: Text("Einstellungen Page")) {
                Text("Einstellungen")
                    .padding()
            }
            Spacer()
        }
        .frame(maxWidth: 250, alignment: .leading)
        .background(Color(UIColor.systemGray6))
        .edgesIgnoringSafeArea(.all)
    }
}

// Hinweis: Dies ist eine Grundstruktur. Die einzelnen Pages können später mit eigenen Views ausgefüllt werden.
