package template

var ViewTPL = `import 'package:flutter/material.dart';
import 'package:get/get.dart';
import '{{.PageName}}_controller.dart';

class {{.PageNameUp}}View extends GetView<{{.PageNameUp}}Controller> {
  const {{.PageNameUp}}View({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(body: Center(child: Text("{{.PageName}} view")));
  }
}
`
