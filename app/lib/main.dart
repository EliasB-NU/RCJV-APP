import 'package:flutter/material.dart';
import 'pages/matches_part.dart';
import 'pages/home_page.dart';
import 'widgets/side_menu.dart';
import 'pages/settings_part.dart';
import 'pages/links_page.dart';
import 'package:http/http.dart' as http;



void main() {
  // Get config

  runApp(RobocupJuniorApp());
}
Future<http.Response> fetchAlbum() {
  return http.get(Uri.parse('https://jsonplaceholder.typicode.com/albums/1'));
}

class RobocupJuniorApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'RoboCup Junior',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: MainPage(),
      debugShowCheckedModeBanner: false,
    );
  }
}

class MainPage extends StatefulWidget {
  @override
  _MainPageState createState() => _MainPageState();
}

class _MainPageState extends State<MainPage> {
  int _selectedIndex = 0;
  final List<Widget> _pages = [
  HomePage(),
  SpielplanPage(), // <- wichtig: hier muss SpielplanPage eingebunden sein
  PlaceholderWidget(title: 'Infos'),
  LinksPage(),
  SettingsPage(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
    Navigator.pop(context);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Robocup Junior Vöhringen')),
      drawer: SideMenu(onItemTapped: _onItemTapped),
      body: _pages[_selectedIndex],
    );
  }
}

class PlaceholderWidget extends StatelessWidget {
  final String title;
  const PlaceholderWidget({required this.title});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Text('$title Seite (noch leer)', style: TextStyle(fontSize: 18)),
    );
  }
}