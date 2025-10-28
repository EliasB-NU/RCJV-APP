import 'dart:convert';
import 'package:http/http.dart' as http;

Future<void> fetchJson() async {
  final url = Uri.parse('http://localhost:3006/data/test.json');

  try {
    final response = await http.get(url);

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      print('Daten empfangen: $data');
    } else {
      print('Fehler: ${response.statusCode}');
    }
  } catch (e) {
    print('Fehler beim Laden der Daten: $e');
  }
}

void main() {
  fetchJson();
}
