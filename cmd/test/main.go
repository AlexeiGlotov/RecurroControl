package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// checkFile проверяет файл на соответствие правилу.
func checkFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	foundError := false
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.Contains(line, "newErrorResponse") && scanner.Scan() {
			lineNumber++
			nextLine := scanner.Text()
			if !strings.Contains(nextLine, "return") {
				foundError = true
				fmt.Printf("Найдено нарушение в %s:%d\n", filePath, lineNumber)
			}
		}
	}

	if !foundError {
		fmt.Printf("В файле %s ошибок не найдено\n", filePath)
	}

	return scanner.Err()
}

// processDirectory обрабатывает все Go-файлы в директории.
func processDirectory(rootPath string) error {
	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			return checkFile(path)
		}
		return nil
	})
}

func main() {
	rootPath := "." // замените на путь к вашему проекту
	if err := processDirectory(rootPath); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
	}
}
