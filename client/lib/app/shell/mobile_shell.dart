import 'package:flutter/material.dart';

class MobileShell extends StatelessWidget {
  final Widget child;
  const MobileShell({super.key, required this.child});

  @override
  Widget build(BuildContext context) {
    return Scaffold(body: child);
  }
}
