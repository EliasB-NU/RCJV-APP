// pages/home_page.dart V1
/*import 'package:flutter/material.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String? selectedLeague;
  String? selectedTeam;

  final List<String> leagues = ['Liga A', 'Liga B', 'Liga C'];
  final List<String> teams = ['Team Alpha', 'Team Beta', 'Team Gamma'];

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Text(
            'Robocup Junior Vöhringen',
            style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
            textAlign: TextAlign.center,
          ),
          SizedBox(height: 24),
          DropdownButton<String>(
            hint: Text('Liga auswählen'),
            value: selectedLeague,
            onChanged: (value) {
              setState(() {
                selectedLeague = value;
              });
            },
            items: leagues.map((league) {
              return DropdownMenuItem(
                value: league,
                child: Text(league),
              );
            }).toList(),
          ),
          SizedBox(height: 16),
          DropdownButton<String>(
            hint: Text('Team auswählen'),
            value: selectedTeam,
            onChanged: (value) {
              setState(() {
                selectedTeam = value;
              });
            },
            items: teams.map((team) {
              return DropdownMenuItem(
                value: team,
                child: Text(team),
              );
            }).toList(),
          ),
          SizedBox(height: 24),
          ElevatedButton(
            onPressed: () {
              if (selectedLeague != null && selectedTeam != null) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Favorit gespeichert: $selectedTeam in $selectedLeague')),
                );
              } else {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Bitte Liga und Team auswählen!')),
                );
              }
            },
            child: Text('Als Favorit speichern'),
          ),
        ],
      ),
    );
  }
}

// pages/home_page.dart V2
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String? selectedLeague;
  String? selectedTeam;

  final List<String> leagues = ['Liga A', 'Liga B', 'Liga C'];
  final List<String> teams = ['Team Alpha', 'Team Beta', 'Team Gamma'];

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Text(
            'Robocup Junior Vöhringen',
            style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
            textAlign: TextAlign.center,
          ),
          SizedBox(height: 24),
          DropdownButton<String>(
            hint: Text('Liga auswählen'),
            value: selectedLeague,
            onChanged: (value) {
              setState(() {
                selectedLeague = value;
              });
            },
            items: leagues.map((league) {
              return DropdownMenuItem(
                value: league,
                child: Text(league),
              );
            }).toList(),
          ),
          SizedBox(height: 16),
          DropdownButton<String>(
            hint: Text('Team auswählen'),
            value: selectedTeam,
            onChanged: (value) {
              setState(() {
                selectedTeam = value;
              });
            },
            items: teams.map((team) {
              return DropdownMenuItem(
                value: team,
                child: Text(team),
              );
            }).toList(),
          ),
          SizedBox(height: 24),
          ElevatedButton(
            onPressed: () async {
              if (selectedLeague != null && selectedTeam != null) {
                final prefs = await SharedPreferences.getInstance();
                await prefs.setString('favoriteLeague', selectedLeague!);
                await prefs.setString('favoriteTeam', selectedTeam!);
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Favorit gespeichert: $selectedTeam in $selectedLeague')),
                );
              } else {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Bitte Liga und Team auswählen!')),
                );
              }
            },
            child: Text('Als Favorit speichern'),
          ),
        ],
      ),
    );
  }
}

*/

// pages/home_page.dart
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String? selectedLeague;
  String? selectedTeam;
  String? selectedInstitution;

  final List<String> leagues = ['Liga A', 'Liga B', 'Liga C'];
  final List<String> teams = ['Team Alpha', 'Team Beta', 'Team Gamma'];
  final List<String> institutions = ['Schule 1', 'Schule 2', 'Schule 3'];

  Future<void> _saveTeamFavorite() async {
    if (selectedLeague == null || selectedTeam == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Bitte wähle deine Liga UND Team aus')),
      );
      return;
    }

    final prefs = await SharedPreferences.getInstance();
    List<String> favorites = prefs.getStringList('teamFavorites') ?? [];
    favorites.add('Liga: $selectedLeague, Team: $selectedTeam');
    await prefs.setStringList('teamFavorites', favorites);

    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Favorit unter „Teams“ gespeichert')),
    );

    setState(() {
      selectedLeague = null;
      selectedTeam = null;
    });
  }

  Future<void> _saveInstitutionFavorite() async {
    if (selectedInstitution == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Bitte wähle deine Institution aus')),
      );
      return;
    }

    final prefs = await SharedPreferences.getInstance();
    List<String> favorites = prefs.getStringList('institutionFavorites') ?? [];
    favorites.add('Institution: $selectedInstitution');
    await prefs.setStringList('institutionFavorites', favorites);

    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Favorit unter „Institution“ gespeichert')),
    );

    setState(() {
      selectedInstitution = null;
    });
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: [
          Text(
            'Robocup Junior Vöhringen',
            style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
          ),
          SizedBox(height: 24),

          // Liga & Team Dropdowns nebeneinander
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

          // Institution Dropdown
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
