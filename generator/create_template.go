package generator

import (
	"fmt"
	"strings"

	"github.com/isfk/flutter-cli/generator/template"
)

func CreatePage(appName, pageName string) error {
	filesPath := "./lib/pages"
	g := New(
		PageName(pageName),
		PageNameUp(strings.ToUpper(pageName[:1])+pageName[1:]),
	)
	if len(appName) > 0 {
		filesPath = fmt.Sprintf("./%v/lib/pages", appName)
	}

	files := []File{
		{Path: fmt.Sprintf("%v/%v/%v_binding.dart", filesPath, pageName, pageName), Template: template.BindingTPL},
		{Path: fmt.Sprintf("%v/%v/%v_controller.dart", filesPath, pageName, pageName), Template: template.ControllerTPL},
		{Path: fmt.Sprintf("%v/%v/%v_model.dart", filesPath, pageName, pageName), Template: template.ModelTPL},
		{Path: fmt.Sprintf("%v/%v/%v_view.dart", filesPath, pageName, pageName), Template: template.ViewTPL},
	}

	err := g.Generate(files)
	if err != nil {
		return err
	}
	return nil
}

func CreateFiles(appName string) error {
	if err := CreatePage(appName, "home"); err != nil {
		return err
	}

	if err := CreatePage(appName, "notfound"); err != nil {
		return err
	}

	g := New(
		AppName(appName),
	)

	files := []File{
		{Path: fmt.Sprintf("./%v/lib/main.dart", appName), Template: template.MainTPL},
		{Path: fmt.Sprintf("./%v/lib/global.dart", appName), Template: template.GlobalTPL},
		{Path: fmt.Sprintf("./%v/lib/routes/app_routes.dart", appName), Template: template.AppRoutesTPL},
		{Path: fmt.Sprintf("./%v/lib/routes/app_pages.dart", appName), Template: template.AppPagesTPL},
		{Path: fmt.Sprintf("./%v/lib/utils/request.dart", appName), Template: template.RequestTPL},
		{Path: fmt.Sprintf("./%v/lib/utils/storage.dart", appName), Template: template.StorageTPL},
		{Path: fmt.Sprintf("./%v/pubspec.yaml", appName), Template: template.PubspecTPL},
	}

	err := g.Generate(files)
	if err != nil {
		return err
	}
	return nil
}
