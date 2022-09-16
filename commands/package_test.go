package commands

// 테스트코드 작성
//
// Language: go
// Path: commands\package_test.go
// Compare this snippet from commands\package_test.go:
// package commands
//
import (
	"testing"
)

func TestGetPythonVersion(t *testing.T) {
	pythonVersion, err := GetPythonVersion()
	if err != nil {
		t.Error(err)
	}

	t.Log(pythonVersion)

	if pythonVersion < 3.8 {
		t.Error("python version is lower than 3")
	}
}
