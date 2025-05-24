import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class EndpointTestPage extends StatefulWidget {
  const EndpointTestPage({Key? key}) : super(key: key);

  @override
  State<EndpointTestPage> createState() => _EndpointTestPageState();
}

class _EndpointTestPageState extends State<EndpointTestPage> {
  late Future<List<dynamic>> _daten;

  @override
  void initState() {
    super.initState();
    _daten = fetchDaten();
  }

  Future<List<dynamic>> fetchDaten() async {
    final url = Uri.parse('http://10.0.2.2:3006/api/leagues'); // ggf. anpassen

    try {
      final response = await http.get(url);

      if (response.statusCode == 200) {
        final List<dynamic> daten = json.decode(response.body);
        return daten;
      } else {
        throw Exception('Serverfehler: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Verbindungsfehler: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('API-Daten')),
      drawer: Drawer(
        child: ListView(
          padding: EdgeInsets.zero,
          children: [
            const DrawerHeader(
              decoration: BoxDecoration(color: Colors.blue),
              child: Text('Navigation', style: TextStyle(color: Colors.white, fontSize: 24)),
            ),
            ListTile(
              title: const Text('Startseite'),
              onTap: () {
                Navigator.pushNamed(context, '/');
              },
            ),
            ListTile(
              title: const Text('API Test'),
              onTap: () {
                Navigator.pushNamed(context, '/endpoint');
              },
            ),
          ],
        ),
      ),
      body: FutureBuilder<List<dynamic>>(
        future: _daten,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(
              child: Text(
                'Fehler:\n${snapshot.error}',
                textAlign: TextAlign.center,
                style: const TextStyle(color: Colors.red),
              ),
            );
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(child: Text('Keine Daten gefunden.'));
          } else {
            final daten = snapshot.data!;
            return ListView.builder(
              itemCount: daten.length,
              itemBuilder: (context, index) {
                return ListTile(
                  title: Text(daten[index].toString()),
                );
              },
            );
          }
        },
      ),
    );
  }
}


/*import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class EndpointTestPage extends StatefulWidget {
  const EndpointTestPage({Key? key}) : super(key: key);

  @override
  State<EndpointTestPage> createState() => _EndpointTestPageState();
}

class _EndpointTestPageState extends State<EndpointTestPage> {
  late Future<List<dynamic>> _daten;

  @override
  void initState() {
    super.initState();
    _daten = fetchDaten();
  }

  Future<List<dynamic>> fetchDaten() async {
    final url = Uri.parse('http://10.0.2.2:3006/api/dein-endpunkt'); // ← hier ggf. anpassen

    try {
      final response = await http.get(url);

      if (response.statusCode == 200) {
        final List<dynamic> daten = json.decode(response.body);
        return daten;
      } else {
        throw Exception('Serverfehler: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Verbindungsfehler: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Endpoint Test')),
      body: FutureBuilder<List<dynamic>>(
        future: _daten,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(
              child: Text(
                'Fehler beim Laden:\n${snapshot.error}',
                textAlign: TextAlign.center,
                style: const TextStyle(color: Colors.red),
              ),
            );
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(child: Text('Keine Daten gefunden.'));
          } else {
            final daten = snapshot.data!;
            return ListView.builder(
              itemCount: daten.length,
              itemBuilder: (context, index) {
                final item = daten[index];
                return ListTile(
                  title: Text(item.toString()),
                );
              },
            );
          }
        },
      ),
    );
  }
}*/
