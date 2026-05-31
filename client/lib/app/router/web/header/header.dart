import 'package:flutter/material.dart';

class AppHeader extends StatelessWidget implements PreferredSizeWidget {
  final Widget? title;
  final Widget? search;
  final List<Widget>? actions;
  final Widget? profile;

  const AppHeader({
    super.key,
    this.title,
    this.search,
    this.actions,
    this.profile,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      height: preferredSize.height,
      padding: const EdgeInsets.symmetric(horizontal: 16),
      decoration: const BoxDecoration(
        color: Colors.white,
        border: Border(
          bottom: BorderSide(color: Color(0xFFEAEAEA)),
        ),
      ),
      child: Row(
        children: [
          title ?? const Text("Dashboard"),

          const SizedBox(width: 24),

          Expanded(
            child: search ?? const SizedBox(),
          ),

          const SizedBox(width: 16),

          Row(
            mainAxisSize: MainAxisSize.min,
            children: [
              ...(actions ?? []),
              const SizedBox(width: 12),
              profile ?? const SizedBox(),
            ],
          ),
        ],
      ),
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(64);
}