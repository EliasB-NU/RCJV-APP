import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
// ignore: unused_import
import 'package:robocup_junior_app/models/matches.dart'; 

class SpielplanPage extends StatefulWidget {
  @override
  _SpielplanPageState createState() => _SpielplanPageState();
}

class _SpielplanPageState extends State<SpielplanPage> {
  List<String> teamFavorites = [];
  List<String> institutionFavorites = [];
  List<Match> filteredMatches = [];

  @override
  void initState() {
    super.initState();
    _loadFavorites().then((_) => _fetchAndFilterMatches());
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
    _fetchAndFilterMatches(); // Liste aktualisieren
  }

  Future<void> _removeInstitutionFavorite(int index) async {
    final prefs = await SharedPreferences.getInstance();
    institutionFavorites.removeAt(index);
    await prefs.setStringList('institutionFavorites', institutionFavorites);
    setState(() {});
    _fetchAndFilterMatches(); // Liste aktualisieren
  }

  Future<void> _fetchAndFilterMatches() async {
    try {
      final response = await http.get(Uri.parse('http://10.0.2.2:3006')); // Bei iOS evtl. localhost
      if (response.statusCode == 200) {
        final decoded = json.decode(response.body);
        final List<dynamic> allMatchesJson = decoded['other'];
        final List<Match> allMatches = allMatchesJson.map((json) => Match.fromJson(json)).toList();

        final Map<String, List<String>> leagueTeamMap = {};
        for (final fav in teamFavorites) {
          final parts = fav.split(':');
          if (parts.length == 2) {
            final league = parts[0];
            final team = parts[1];
            leagueTeamMap.putIfAbsent(league, () => []).add(team);
          }
        }

        setState(() {
          filteredMatches = allMatches.where((match) {
            final inLeagueTeams = leagueTeamMap[match.league]?.contains(match.teamName) ?? false;
            final inInstitutions = institutionFavorites.contains(match.institutionName);
            return inLeagueTeams || inInstitutions;
          }).toList();
        });
      } else {
        print('Fehler beim Laden: ${response.statusCode}');
      }
    } catch (e) {
      print('Fehler beim Abrufen des Spielplans: $e');
    }
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
      }),

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
      }),

      SizedBox(height: 32),
      Text('Spielplan',
          style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
      SizedBox(height: 8),

      if (filteredMatches.isEmpty)
        Text('Kein Spielplan für die aktuellen Favoriten verfügbar.'),

      // Favoriten-Anzeige
      ...teamFavorites.map((favorite) {
        final split = favorite.split(':');
        if (split.length != 2) return Container();

        final league = split[0];
        final team = split[1];
        final matches = filteredMatches.where((m) => m.league == league && m.teamName == team).toList();
        return _buildFavoritenMatchGroup('$league – $team', matches);
      }),

      ...institutionFavorites.map((institution) {
        final matches = filteredMatches.where((m) => m.institutionName == institution).toList();
        return _buildFavoritenMatchGroup(institution, matches);
      }),
    ],
  );
}

  }


// Modellklasse für ein Match
class Match {
  final String name;
  final String field;
  final DateTime startTime;
  final String teamName;
  final String institutionName;
  final String league;

  Match({
    required this.name,
    required this.field,
    required this.startTime,
    required this.teamName,
    required this.institutionName,
    required this.league,
  });

  factory Match.fromJson(Map<String, dynamic> json) {
    return Match(
      name: json['name'],
      field: json['field'],
      startTime: DateTime.parse(json['startTime']),
      teamName: json['teamName'],
      institutionName: json['institutionName'],
      league: json['league'],
    );
  }
}

Widget _buildFavoritenMatchGroup(String title, List<Match> matches) {
  return ExpansionTile(
    title: Text(title, style: TextStyle(fontWeight: FontWeight.bold)),
    children: matches.map((match) {
      final hasError = match.teamName.isEmpty || match.field.isEmpty || match.startTime == DateTime(2000);
      final textStyle = hasError ? TextStyle(color: Colors.red) : null;

      return ExpansionTile(
        title: Text('${match.name}', style: textStyle),
        subtitle: Text(
          'Zeit: ${match.startTime != DateTime(2000) ? match.startTime.toLocal() : 'Upload error'} – '
          'Feld: ${match.field.isNotEmpty ? match.field : 'Upload error'}',
          style: textStyle,
        ),
        children: [
          ListTile(
            title: Text(
              'Team: ${match.teamName.isNotEmpty ? match.teamName : 'Upload error'}',
              style: match.teamName.isEmpty ? TextStyle(color: Colors.red) : null,
            ),
            subtitle: Text(
              'Institution: ${match.institutionName.isNotEmpty ? match.institutionName : 'Upload error'}',
              style: match.institutionName.isEmpty ? TextStyle(color: Colors.red) : null,
            ),
          ),
        ],
      );
    }).toList(),
  );
}

