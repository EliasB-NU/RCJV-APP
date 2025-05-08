// pages/spielplan_page.dart
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class SpielplanPage extends StatefulWidget {
  @override
  _SpielplanPageState createState() => _SpielplanPageState();
}

class _SpielplanPageState extends State<SpielplanPage> {
  List<String> teamFavorites = [];
  List<String> institutionFavorites = [];

  @override
  void initState() {
    super.initState();
    _loadFavorites();
  }

  Future<void> _loadFavorites() async {
    final prefs = await SharedPreferences.getInstance();
    setState(() {
      teamFavorites = prefs.getStringList('teamFavorites') ?? [];
      institutionFavorites = prefs.getStringList('institutionFavorites') ?? [];
    });
  }

  Future<void> _removeTeamFavorite(int index) async {
    final prefs = await SharedPreferences.getInstance();
    teamFavorites.removeAt(index);
    await prefs.setStringList('teamFavorites', teamFavorites);
    setState(() {});
  }

  Future<void> _removeInstitutionFavorite(int index) async {
    final prefs = await SharedPreferences.getInstance();
    institutionFavorites.removeAt(index);
    await prefs.setStringList('institutionFavorites', institutionFavorites);
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: const EdgeInsets.all(16.0),
      children: [
        Text('Favoriten – Teams / Liga',
            style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
        SizedBox(height: 8),
        if (teamFavorites.isEmpty)
          Text('Keine Favoriten gespeichert.'),
        ...teamFavorites.asMap().entries.map((entry) {
          int index = entry.key;
          String favorite = entry.value;
          return Card(
            child: ListTile(
              title: Text(favorite),
              trailing: IconButton(
                icon: Icon(Icons.delete),
                onPressed: () => _removeTeamFavorite(index),
              ),
            ),
          );
        }).toList(),

        SizedBox(height: 24),
        Text('Favoriten – Institution',
            style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
        SizedBox(height: 8),
        if (institutionFavorites.isEmpty)
          Text('Keine Favoriten gespeichert.'),
        ...institutionFavorites.asMap().entries.map((entry) {
          int index = entry.key;
          String favorite = entry.value;
          return Card(
            child: ListTile(
              title: Text(favorite),
              trailing: IconButton(
                icon: Icon(Icons.delete),
                onPressed: () => _removeInstitutionFavorite(index),
              ),
            ),
          );
        }).toList(),
      ],
    );
  }
}
