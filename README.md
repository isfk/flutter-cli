# `Flutter + Get` 一键生成新项目

> 注意事项

1. 请先安装 `Flutter`，否则不能正常生成
2. 请查看文件列表结构
3. 生成新页面后，请在 `routes` 下面两个文件添加
4. 页面导航 `Get.toNamed("/user")`
5. `Request` 使用:
   5.1 `Request().get(url);`
   5.2 `Request().post(url, data, headers);`
6. `Storage` 使用:
   6.1 `Storage.get("test_data");`
   6.2 `Storage.set("test_data", data);`
   6.3 `Storage.has("test_data");`
   6.4 `Storage.remove("test_data");`
   6.4 `Storage.erase();`

## 使用(使用直接看这里就可以了)

```shell
go install github.com/isfk/flutter-cli@latest

flutter-cli create myapp
flutter-cli page user
```

生成:

```
lib
├── global.dart
├── main.dart
├── pages
│   ├── home
│   │   ├── home_binding.dart
│   │   ├── home_controller.dart
│   │   ├── home_model.dart
│   │   └── home_view.dart
│   ├── notfound
│   │   ├── notfound_binding.dart
│   │   ├── notfound_controller.dart
│   │   ├── notfound_model.dart
│   │   └── notfound_view.dart
│   └── user
│       ├── user_binding.dart
│       ├── user_controller.dart
│       ├── user_model.dart
│       └── user_view.dart
├── routes
│   ├── app_pages.dart
│   └── app_routes.dart
└── utils
    ├── request.dart
    └── storage.dart
```

## 安装 `cobra-cli`

`go install github.com/spf13/cobra-cli@latest`

## 项目初始化

```shell
go mod init
cobra-cli init --author "sfk@live.cn"
```
