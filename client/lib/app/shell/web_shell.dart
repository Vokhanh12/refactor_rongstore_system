import 'package:client/app/router/web/sidebar/sidebar.dart';
import 'package:flutter/material.dart';

class WebShell extends StatelessWidget {
  final Widget child;

  const WebShell({super.key, required this.child});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFFF5F6FA),
      body: Row(
        children: [
          // SIDEBAR
          const SizedBox(
            width: 260,
            child: Sidebar(),
          ),

          // MAIN AREA
          Expanded(
            child: Column(
              children: [
                // HEADER SLOT (nếu bạn đã có AppHeader)
                Container(
                  height: 64,
                  decoration: const BoxDecoration(
                    color: Colors.white,
                    border: Border(
                      bottom: BorderSide(
                        color: Color(0xFFEAEAEA),
                        width: 1,
                      ),
                    ),
                  ),
                  padding: const EdgeInsets.symmetric(horizontal: 16),
                  child: const Row(
                    children: [
                      Text(
                        "Dashboard",
                        style: TextStyle(
                          fontSize: 16,
                          fontWeight: FontWeight.w600,
                        ),
                      ),
                    ],
                  ),
                ),

                // CONTENT AREA
                Expanded(
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: child,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}