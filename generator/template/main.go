package template

var MainTPL = `import 'package:{{.AppName}}/pages/home/home_binding.dart';
import 'package:{{.AppName}}/global.dart';
import 'package:{{.AppName}}/routes/app_pages.dart';
import 'package:get/get.dart';
import 'package:flutter/material.dart';

void main() {
  Global.init().then((value) => runApp(const MyApp()));
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'Flutter And Get',
      theme: ThemeData(useMaterial3: true),
      getPages: AppViews.routes,
      initialRoute: AppViews.initial,
      initialBinding: HomeBinding(),
    );
  }
}
`