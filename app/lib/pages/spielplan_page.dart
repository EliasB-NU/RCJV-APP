import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class SpielplanPage extends StatefulWidget {
  @override
  _SpielplanPageState createState() => _SpielplanPageState();
}

class _SpielplanPageState extends State<SpielplanPage> {
  String? favoriteLeague;
  String? favoriteTeam;

  @override
  void initState() {
    super.initState();
    loadFavorites();
  }

  Future<void> loadFavorites() async {
    final prefs = await SharedPreferences.getInstance();
    setState(() {
      favoriteLeague = prefs.getString('favoriteLeague');
      favoriteTeam = prefs.getString('favoriteTeam');
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            'Spielplan',
            style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
          ),
          SizedBox(height: 16),
          if (favoriteLeague != null && favoriteTeam != null)
            Text('Gespeicherter Favorit: $favoriteTeam in $favoriteLeague',
                style: TextStyle(fontSize: 18)),
          if (favoriteLeague == null || favoriteTeam == null)
            Text('Kein Favorit gespeichert.', style: TextStyle(fontSize: 18)),
        ],
      ),
    );
  }
}
