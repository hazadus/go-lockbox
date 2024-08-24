package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName         = "lockbox_test"
	lockboxFileName = ".lockbox_test.json"
)

const envVarWithFileName = "GO_LOCKBOX_FILENAME"

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// Сохраним значение переменной окружения, чтобы
	// изменить на время тестов, а потом вернуть, как
	// было:
	envLockboxFileName := os.Getenv(envVarWithFileName)
	os.Setenv(envVarWithFileName, lockboxFileName)

	buildCmd := exec.Command("go", "build", "-o", binName)
	if err := buildCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Can't build binary %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(lockboxFileName)

	os.Setenv(envVarWithFileName, envLockboxFileName)
	os.Exit(result)
}

func TestLockboxCLI(t *testing.T) {
	title := "website"
	password := "password"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewRecordFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", title, "-pwd", password)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetRecord", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-get", title)

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := password
		if expected != string(output) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(output))
		}
	})
}
