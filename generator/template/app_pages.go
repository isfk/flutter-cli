package template

var AppPagesTPL = `import 'package:{{.AppName}}/pages/home/home_binding.dart';
import 'package:{{.AppName}}/pages/home/home_view.dart';
import 'package:{{.AppName}}/pages/notfound/notfound_binding.dart';
import 'package:{{.AppName}}/pages/notfound/notfound_view.dart';
import 'package:get/get.dart';

part 'app_routes.dart';

abstract class AppViews {
  static const initial = AppRoutes.home;

  static final routes = [
    GetPage(
      name: AppRoutes.home,
      page: () => const HomeView(),
      bindings: [
        HomeBinding(),
      ],
    ),
  ];

  static final unknownRoute = GetPage(
    name: AppRoutes.notfound,
    page: () => const NotfoundView(),
    binding: NotfoundBinding(),
  );
}
`