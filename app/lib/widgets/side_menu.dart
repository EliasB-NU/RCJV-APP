import 'package:flutter/material.dart';

class SideMenu extends StatelessWidget {
  final Function(int) onItemTapped;
  const SideMenu({required this.onItemTapped});

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: [
          DrawerHeader(
            decoration: BoxDecoration(color: Colors.blue),
            child: Text('Menü', style: TextStyle(color: Colors.white, fontSize: 24)),
          ),
          ListTile(
            leading: Icon(Icons.home),
            title: Text('Home'),
            onTap: () => onItemTapped(0),
          ),
          ListTile(
            leading: Icon(Icons.schedule),
            title: Text('Spielplan'),
            onTap: () => onItemTapped(1),
          ),
          ListTile(
            leading: Icon(Icons.info),
            title: Text('Infos'),
            onTap: () => onItemTapped(2),
          ),
          ListTile(
            leading: Icon(Icons.link),
            title: Text('Links'),
            onTap: () => onItemTapped(3),
          ),
          ListTile(
            leading: Icon(Icons.settings),
            title: Text('Einstellungen'),
            onTap: () => onItemTapped(4),
          ),
        ],
      ),
    );
  }
}
