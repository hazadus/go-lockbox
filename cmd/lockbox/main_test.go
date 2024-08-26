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
	lockboxFileName = ".lockbox_test_data"
	secret          = "1234567890123456"
)

const envVarWithFileName = "GO_LOCKBOX_FILENAME"
const envVarWithSecret = "GO_LOCKBOX_SECRET"

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// Сохраним значения переменных окружения, чтобы
	// изменить на время тестов, а потом вернуть, как
	// было:
	envLockboxFileName := os.Getenv(envVarWithFileName)
	envSecret := os.Getenv(envVarWithSecret)
	os.Setenv(envVarWithFileName, lockboxFileName)
	os.Setenv(envVarWithSecret, secret)

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
	os.Setenv(envVarWithSecret, envSecret)
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

	t.Run("ListRecords", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%s\n", title)
		if expected != string(output) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(output))
		}
	})

	t.Run("DeleteRecord", func(t *testing.T) {
		// Добавить запись, которую будем удалять
		tempTitle := "temporaryRecord"
		cmd := exec.Command(cmdPath, "-add", tempTitle, "-pwd", "12345678")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		// Удаляем
		cmd = exec.Command(cmdPath, "-del", tempTitle)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		// Пробуем получить удалённую запись: ничего не должно вернуться
		cmd = exec.Command(cmdPath, "-get", tempTitle)

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		if string(output) != "" {
			t.Errorf("Got %q instead of empty string.\n", string(output))
		}
	})
}
