// pages/einstellungen_page.dart
import 'package:flutter/material.dart';

class EinstellungenPage extends StatefulWidget {
  @override
  _EinstellungenPageState createState() => _EinstellungenPageState();
}

class _EinstellungenPageState extends State<EinstellungenPage> {
  bool notificationsEnabled = false; // Benachrichtigungen standardmäßig aus

  void _showInfoDialog() {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: Text('App-Version'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Version: 1.0.0'),
            SizedBox(height: 16),
            Text('''
      '''),
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.of(context).pop(),
            child: Text('Schließen'),
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text('Benachrichtigungen', style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold)),
          SwitchListTile(
            title: Text('Benachrichtigungen aktivieren'),
            value: notificationsEnabled,
            onChanged: (bool value) {
              setState(() {
                notificationsEnabled = value;
              });
            },
          ),
          SizedBox(height: 24),
          IconButton(
            onPressed: _showInfoDialog,
            icon: Icon(Icons.info),
          ),
        ],
      ),
    );
  }
}
