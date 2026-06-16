package semantic

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
)

func TestBuildProjectModelResolvesMultiFileInternalImport(t *testing.T) {
	root := writeSemanticProject(t, map[string]string{
		"src/main/java/com/example/model/User.java": `package com.example.model;

public class User {
  public static int id() {
    return 1;
  }
}`,
		"src/main/java/com/example/service/UserService.java": `package com.example.service;

import com.example.model.User;

public class UserService {
  public static int read(User user) {
    return User.id();
  }
}`,
	})

	model := buildModel(t, root)
	if len(model.Diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got %#v", model.Diagnostics)
	}
	if _, ok := model.Packages["com.example.model"]; !ok {
		t.Fatalf("expected model package in %#v", model.Packages)
	}
	user, ok := model.Classes["com.example.model.User"]
	if !ok {
		t.Fatalf("expected User class in %#v", model.Classes)
	}
	idMethods := user.Methods["id"]
	if len(idMethods) != 1 || idMethods[0].ReturnType.Kind != ir.KindInt {
		t.Fatalf("expected int id method, got %#v", idMethods)
	}
	service := model.Classes["com.example.service.UserService"]
	readMethods := service.Methods["read"]
	if len(readMethods) != 1 || len(readMethods[0].Params) != 1 {
		t.Fatalf("expected read(User) method, got %#v", readMethods)
	}
	if readMethods[0].Params[0].Type.Kind != ir.KindObject || readMethods[0].Params[0].Type.Name != "User" {
		t.Fatalf("expected object param User, got %#v", readMethods[0].Params[0])
	}
}

func TestBuildProjectModelReportsUnresolvedInternalImport(t *testing.T) {
	root := writeSemanticProject(t, map[string]string{
		"src/main/java/com/example/model/User.java": `package com.example.model;

public class User {
  public static int id() {
    return 1;
  }
}`,
		"src/main/java/com/example/service/UserService.java": `package com.example.service;

import com.example.model.Missing;

public class UserService {
  public static int read() {
    return 1;
  }
}`,
	})

	model := buildModel(t, root)
	if !hasSemanticCode(model.Diagnostics, CodeImportUnresolved) {
		t.Fatalf("expected unresolved import diagnostic, got %#v", model.Diagnostics)
	}
}

func TestBuildProjectModelReportsDuplicateClass(t *testing.T) {
	root := writeSemanticProject(t, map[string]string{
		"src/main/java/com/example/User.java": `package com.example;

public class User {
  public static int one() {
    return 1;
  }
}`,
		"src/main/java/com/example/UserCopy.java": `package com.example;

public class User {
  public static int two() {
    return 2;
  }
}`,
	})

	model := buildModel(t, root)
	if !hasSemanticCode(model.Diagnostics, CodeDuplicateSymbol) {
		t.Fatalf("expected duplicate class diagnostic, got %#v", model.Diagnostics)
	}
}

func TestBuildProjectModelReportsDuplicateField(t *testing.T) {
	root := writeSemanticProject(t, map[string]string{
		"src/main/java/com/example/User.java": `package com.example;

public class User {
  private int value;
  private double value;
}`,
	})

	model := buildModel(t, root)
	if !hasSemanticCode(model.Diagnostics, CodeDuplicateSymbol) {
		t.Fatalf("expected duplicate field diagnostic, got %#v", model.Diagnostics)
	}
}

func TestBuildProjectModelReportsUnsupportedOverload(t *testing.T) {
	root := writeSemanticProject(t, map[string]string{
		"src/main/java/com/example/MathUtil.java": `package com.example;

public class MathUtil {
  public static int value(int input) {
    return input;
  }

  public static double value(double input) {
    return input;
  }
}`,
	})

	model := buildModel(t, root)
	if !hasSemanticCode(model.Diagnostics, CodeOverloadUnsupported) {
		t.Fatalf("expected overload diagnostic, got %#v", model.Diagnostics)
	}
}

func buildModel(t *testing.T, root string) ProjectModel {
	t.Helper()
	project, err := javaproject.Scan(root)
	if err != nil {
		t.Fatalf("scan project: %v", err)
	}
	return BuildProjectModel(project)
}

func hasSemanticCode(diagnostics []Diagnostic, code string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code {
			return true
		}
	}
	return false
}

func writeSemanticProject(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for path, body := range files {
		fullPath := filepath.Join(root, filepath.FromSlash(path))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			t.Fatalf("mkdir %s: %v", fullPath, err)
		}
		if err := os.WriteFile(fullPath, []byte(body), 0644); err != nil {
			t.Fatalf("write %s: %v", fullPath, err)
		}
	}
	return root
}
