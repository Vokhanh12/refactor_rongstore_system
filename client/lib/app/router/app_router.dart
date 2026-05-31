import 'package:client/features/dashbroad/presentation/web/page/dashbroad_page.dart';
import 'package:go_router/go_router.dart';

import '../shell/app_shell.dart';

class AppRouter {
  static final router = GoRouter(
    initialLocation: '/dashboard',
    routes: [
      ShellRoute(
        builder: (context, state, child) {
          return AppShell(child: child);
        },
        routes: [
          GoRoute(
            path: '/dashboard',
            builder: (_, __) => const DashboardPage(),
          ),

          // GoRoute(
          //   path: '/iam/',
          //   builder: (_, __) => const UserPage(),
          // ),

          // GoRoute(
          //   path: '/fnb/',
          //   builder: (_, __) => const UserPage(),
          // ),

          // GoRoute(
          //   path: '/hr/',
          //   builder: (_, __) => const UserPage(),
          // ),
        ],
      ),
    ],
  );
}