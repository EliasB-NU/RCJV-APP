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
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              Expanded(
                child: DropdownButton<String>(
                  isExpanded: true,
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
              ),
              SizedBox(width: 16),
              Expanded(
                child: DropdownButton<String>(
                  isExpanded: true,
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
              ),
            ],
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
