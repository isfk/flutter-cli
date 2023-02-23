package template

var RequestTPL = `import 'package:get/get.dart';
import 'package:get/get_connect.dart';

class Request {
  static late GetConnect connect;

  static Future init() async {
    connect = GetConnect();
  }

  Future<Response> get(String url) async {
    return connect.get(url);
  }

  Future<Response> post(
      String url, Map data, Map<String, String> headers) async {
    return connect.post(url, data, headers: headers);
  }
}
`
