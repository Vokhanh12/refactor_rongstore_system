import 'package:client/app/shell/mobile_shell.dart';
import 'package:client/app/shell/web_shell.dart';
import 'package:flutter/material.dart';

class AppShell extends StatelessWidget {
  final Widget child;

  const AppShell({super.key, required this.child});

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) {
        if (constraints.maxWidth > 900) {
          return WebShell(child: child);
        }
        return MobileShell(child: child);
      },
    );
  }
}