package template

var GlobalTPL = `import 'package:{{.AppName}}/utils/request.dart';
import 'package:{{.AppName}}/utils/storage.dart';

class Global {
  static bool isLogined = false;

  static Future init() async {
    await Request.init();
    await Storage.init();
  }
}
`
