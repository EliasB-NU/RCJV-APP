import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String? selectedLeague;
  String? selectedTeam;
  String? selectedInstitution;

  List<String> leagues = [];
  List<String> teams = [];
  List<String> institutions = [];

  bool loading = true;

  @override
  void initState() {
    super.initState();
    fetchAllDropdownData();
  }

  Future<void> fetchAllDropdownData() async {
  try {
    // Leagues
    final leaguesResponse = await http.get(Uri.parse('http://localhost:3006/api/v1/leagues'));
    if (leaguesResponse.statusCode == 200) {
      final data = jsonDecode(leaguesResponse.body);
      final extractedLeagues = <String>[];
      data.forEach((key, value) {
        if (value == true) {
          final label = key.replaceAllMapped(RegExp(r'([a-z])([A-Z])'), (m) => '${m[1]} ${m[2]}');
          extractedLeagues.add(label);
        }
      });
      leagues = extractedLeagues;
    }

//Teams
    final teamsResponse = await http.get(Uri.parse('http://localhost:3006/api/v1/teams'));
    if (teamsResponse.statusCode == 200) {
    final data = jsonDecode(teamsResponse.body);

    final List<dynamic> dataList = data['data'];
    final List<String> extractedTeams =
      dataList.map<String>((team) => team['name'] as String).toList();

    setState(() {
    teams = extractedTeams;
    });
  } else {
   print('Fehler beim Laden der Teams: ${teamsResponse.statusCode}');
  }

//Institution
    final institutionResponse = await http.get(Uri.parse('http://localhost:3006/api/v1/teams'));
    if (institutionResponse.statusCode == 200) {
    final data = jsonDecode(institutionResponse.body);

    final List<dynamic> dataList = data['data'];
    final List<String> extractedInstitution =
      dataList.map<String>((institution) => institution['institution'] as String).toList();

    setState(() {
    institutions = extractedInstitution;
    });
  } else {
   print('Fehler beim Laden der Institution: ${institutionResponse.statusCode}');
  }


    // Institutions
    //final instResponse = await http.get(Uri.parse('http://localhost:3006/api/v1/teams'));
    //if (instResponse.statusCode == 200) {
    //  final instData = jsonDecode(instResponse.body);
    //  institutions = List<String>.from(instData['institutions']); // Falls JSON {"institutions": [...]}
    //}

    // Wenn alles geladen wurde
    setState(() {
      loading = false;
    });
  } catch (e) {
    print('Fehler beim Laden der Dropdown-Daten: $e');
  }
}




  Future<void> _saveTeamFavorite() async {
    if (selectedLeague == null || selectedTeam == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Bitte wähle deine Liga UND Team aus'), duration: Duration(seconds: 1)),
      );
      return;
    }

    final prefs = await SharedPreferences.getInstance();
    List<String> favorites = prefs.getStringList('teamFavorites') ?? [];
    favorites.add('Liga: $selectedLeague, Team: $selectedTeam');
    await prefs.setStringList('teamFavorites', favorites);

    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Favorit unter „Teams“ gespeichert'), duration: Duration(seconds: 1)),
    );

    setState(() {
      selectedLeague = null;
      selectedTeam = null;
    });
  }

  Future<void> _saveInstitutionFavorite() async {
    if (selectedInstitution == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Bitte wähle deine Institution aus'), duration: Duration(seconds: 1)),
      );
      return;
    }

    final prefs = await SharedPreferences.getInstance();
    List<String> favorites = prefs.getStringList('institutionFavorites') ?? [];
    favorites.add('Institution: $selectedInstitution');
    await prefs.setStringList('institutionFavorites', favorites);

    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Favorit unter „Institution“ gespeichert'), duration: Duration(seconds: 1)),
    );

    setState(() {
      selectedInstitution = null;
    });
  }

  @override
  Widget build(BuildContext context) {
    return loading
        ? Center(child: CircularProgressIndicator())
        : SingleChildScrollView(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              children: [
                Text(
                  'Robocup Junior Vöhringen',
                  style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
                ),
                SizedBox(height: 24),
                Row(
                  children: [
                    Expanded(
                      child: DropdownButton<String>(
                        hint: Text('Liga auswählen'),
                        value: selectedLeague,
                        onChanged: (value) => setState(() => selectedLeague = value),
                        isExpanded: true,
                        items: leagues.map((league) {
                          return DropdownMenuItem(
                            value: league,
                            child: Text(league),
                          );
                        }).toList(),
                      ),
                    ),
                    SizedBox(width: 16),
                    Expanded(
                      child: DropdownButton<String>(
                        hint: Text('Team auswählen'),
                        value: selectedTeam,
                        onChanged: (value) => setState(() => selectedTeam = value),
                        isExpanded: true,
                        items: selectedLeague != null
                            ? teams.map((team) {
                                return DropdownMenuItem(
                                  value: team,
                                  child: Text(team),
                                );
                              }).toList()
                            : [],
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 8),
                ElevatedButton(
                  onPressed: _saveTeamFavorite,
                  child: Text('Als Favorit unter „Teams“ speichern'),
                ),
                SizedBox(height: 24),
                DropdownButton<String>(
                  hint: Text('Institution auswählen'),
                  value: selectedInstitution,
                  onChanged: (value) => setState(() => selectedInstitution = value),
                  isExpanded: true,
                  items: institutions.map((institution) {
                    return DropdownMenuItem(
                      value: institution,
                      child: Text(institution),
                    );
                  }).toList(),
                ),
                SizedBox(height: 8),
                ElevatedButton(
                  onPressed: _saveInstitutionFavorite,
                  child: Text('Als Favorit unter „Institution“ speichern'),
                ),
              ],
            ),
          );
  }
}
