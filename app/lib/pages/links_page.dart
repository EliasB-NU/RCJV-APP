// pages/links_page.dart
import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

class LinksPage extends StatelessWidget {
  final String robocupUrl = 'https://robocup.rocci.net/';
  final String scoringUrl = 'https://portal.robocup.de/rescue/scoring/uebersicht/';

  Future<void> _launchURL(String url) async {
    if (await canLaunchUrl(Uri.parse(url))) {
      await launchUrl(Uri.parse(url));
    } else {
      throw 'Could not open: $url';
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          ElevatedButton(
            onPressed: () => _launchURL(robocupUrl),
            child: Text('Rocci Robocup Webseite'),
          ),
          SizedBox(height: 16),
          ElevatedButton(
            onPressed: () => _launchURL(scoringUrl),
            child: Text('Scoring Übersicht Vöhringen'),
          ),
        ],
      ),
    );
  }
}
