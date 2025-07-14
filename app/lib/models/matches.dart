// === Datei: models/match.dart ===

class Referee {
  final String firstName;
  final String lastName;

  Referee({required this.firstName, required this.lastName});

  factory Referee.fromJson(Map<String, dynamic> json) {
    return Referee(
      firstName: json['first_name'] ?? '',
      lastName: json['last_name'] ?? '',
    );
  }
}

class Match {
  final String name;
  final String field;
  final DateTime startTime;
  final String teamName;
  final String institutionName;
  final String league;
  final List<Referee>? referees; // ⬅️ optional

  Match({
    required this.name,
    required this.field,
    required this.startTime,
    required this.teamName,
    required this.institutionName,
    required this.league,
    this.referees, // ⬅️ optional
  });

  factory Match.fromJson(Map<String, dynamic> json) {
    return Match(
      name: json['name'] ?? '',
      field: json['field'] ?? '',
      startTime: DateTime.tryParse(json['startTime'] ?? '') ?? DateTime(2000),
      teamName: json['teamName'] ?? '',
      institutionName: json['institutionName'] ?? '',
      league: json['league'] ?? '',
      referees: (json['referees'] as List?)?.map((e) => Referee.fromJson(e)).toList(),
    );
  }
}

