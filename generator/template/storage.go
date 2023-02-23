package template

var StorageTPL = `import 'package:get_storage/get_storage.dart';

class Storage {
  static late GetStorage box;

  static Future init() async {
    box = GetStorage();
  }

  static Future<void> set(String key, dynamic data) {
    return box.write(key, data);
  }

  static dynamic get(String key) {
    return box.read(key);
  }

  static bool has(String key) {
    return box.hasData(key);
  }

  static Future<void> remove(String key) {
    return box.remove(key);
  }

  static Future<void> erase() {
    return box.erase();
  }
}
`
