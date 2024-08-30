package util

import (
	"os"
	"os/exec"
	"bytes"
	"path/filepath"
	"time"
	"fmt"
	"strconv"
	"errors"
)

func RunShell(cmd string) (string, error) {
    command := exec.Command("sh", "-c", cmd)

    var stdout bytes.Buffer
    command.Stdout = &stdout

    command.Stderr = os.Stderr

    err := command.Run()
    if err != nil { return "", err }

    return stdout.String(), nil
}

func GetContent(name, path string) (string, error) {
	file_path, err := filePath(name, path)
	if (err != nil) { return "", err }

	content, err := os.ReadFile(file_path)
	if err != nil { return "", fmt.Errorf("failed to read file: %w", err) }

	return string(content), nil
}

func SetContent(name, path, content string) error {
	file_path, err := filePath(name, path)
	if (err != nil) { return err }

	dir := filepath.Dir(file_path)
	if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directories: %w", err)
	}

	err = os.WriteFile(file_path, []byte(content), 0644)
	if err != nil { return fmt.Errorf("failed to write file: %w", err) }

	return nil
}

func IsCacheValid(name, path, max_age string) (bool, error) {
	file_age, err := FileAge(name, path)
	if (err != nil) { return false, nil }
	if (file_age == -1) { return false, nil }

	max_age_seconds, err := maxAgeSeconds(max_age)
	if (err != nil) { return false, nil }

	return file_age < max_age_seconds, nil
}

func FileAge(name, path string) (int, error) {
	file_path, err := filePath(name, path)
	if (err != nil) { return -1, nil }

	fileInfo, err := os.Stat(file_path)
	if err != nil { return -1, nil }

	modTime := fileInfo.ModTime()
	now := time.Now()
	age := now.Sub(modTime).Seconds()

	return int(age), nil
}

func filePath(name, path string) (string, error) {
    if path[:2] == "~/" {
        home, err := os.UserHomeDir()
        if err != nil { return "", err }
        path = filepath.Join(home, path[2:])
    }

    if path[len(path)-1] == os.PathSeparator { return filepath.Join(path, name), nil }
    return path, nil
}

func maxAgeSeconds(max_age string) (int, error) {
	val, err := strconv.Atoi(max_age[:len(max_age)-1])
	if err != nil { return 0, errors.New("Can't parse max age flag")}

	units := map[string]int {"s": 1, "m": 60, "h": 60*60, "d": 60*60*24 }
	multiplier, ok := units[max_age[len(max_age)-1:]]
	if !ok { return 0, errors.New("Can't parse max age flag")}

	return val*multiplier, nil
}

