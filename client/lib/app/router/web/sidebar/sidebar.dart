import 'package:flutter/material.dart';
import 'sidebar_config.dart';

class Sidebar extends StatelessWidget {
  const Sidebar({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 240,
      color: Colors.grey.shade900,
      child: ListView(
        children: sidebarItems.map((item) {
          return ListTile(
            title: Text(item.title, style: const TextStyle(color: Colors.white)),
            onTap: () => Navigator.pushNamed(context, item.path),
          );
        }).toList(),
      ),
    );
  }
}