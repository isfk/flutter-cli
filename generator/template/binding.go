package template

var BindingTPL = `import 'package:get/get.dart';
import '{{.PageName}}_controller.dart';

class {{.PageNameUp}}Binding implements Bindings {
  @override
  void dependencies() {
    Get.lazyPut<{{.PageNameUp}}Controller>(() => {{.PageNameUp}}Controller());
  }
}
`
