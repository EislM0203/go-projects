package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type FileManager struct {	
	InputFilePath string
	OutputDir string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string 
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error scanning the file: %v", err)
	}

	return lines, nil
}

func (fm FileManager) WriteResult(job interface{}) error {
	time.Sleep(3 * time.Second) // simulated slow file write
	
	dir := "results"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return err
	}

	fullPath := filepath.Join(dir, fm.OutputDir)

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(job); err != nil {
		log.Printf("Failed to encode JSON: %v", err)
		return err
	}

	return nil
}

func New(inputFilePath, outputDir string) *FileManager {
	return &FileManager{
		InputFilePath: inputFilePath,
		OutputDir: outputDir,
	}
}