import 'package:client/app/router/moblie/bottomnav/bottomnav_config.dart';
import 'package:flutter/material.dart';

class Bottomnav extends StatelessWidget {
  const Bottomnav({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 240,
      color: Colors.grey.shade900,
      child: ListView(
        children: bottomItems.map((item) {
          return ListTile(
            title: Text(item.title, style: const TextStyle(color: Colors.white)),
            onTap: () => Navigator.pushNamed(context, item.path),
          );
        }).toList(),
      ),
    );
  }
}