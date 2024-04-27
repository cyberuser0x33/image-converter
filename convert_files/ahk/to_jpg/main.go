/*
FOR J.A.R.V.I.S. project
original project https://github.com/Priler/jarvis

made by cyberuser0x33 https://github.com/cyberuser0x33
*/

package main

import (
    "github.com/disintegration/imaging"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    // Путь к папке c файлами 
    // важно!!! путь к папке в Windows нужно указывать чере двойной слэш "\\"
    // как пример D:\\folder\\files
    
    filesDir := "Путь к папке с файлами"

    files, err := os.ReadDir(filesDir)
    if err != nil {
        log.Fatalf("Error reading directory: %v", err)
    }

    // Создаем папку "output" внутри папки с файлами
    outputDir := filepath.Join(filesDir, "output")
    err = os.MkdirAll(outputDir, 0755)
    if err != nil {
        log.Fatalf("Error creating output directory: %v", err)
    }

    for _, file := range files {
        // Проверяем, является ли файл графическим файлом поддерживаемого формата
        if !file.IsDir() && isValidImageFile(file.Name()) {
            inputPath := filepath.Join(filesDir, file.Name())
            outputPath := filepath.Join(outputDir, file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))]+".jpg")
            convertFile(inputPath, outputPath)
        }
    }
}

func isValidImageFile(filename string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    return ext == ".jpg" || ext == ".png" ||  ext == ".jfif"
}

func convertFile(inputPath, outputPath string) {
    inputImage, err := imaging.Open(inputPath)
    if err != nil {
        log.Printf("Error opening file %s: %v", inputPath, err)
        return
    }

    outputFile, err := os.Create(outputPath)
    if err != nil {
        log.Printf("Error creating file %s: %v", outputPath, err)
        return
    }
    defer outputFile.Close()

    err = imaging.Encode(outputFile, inputImage, imaging.JPEG, imaging.JPEGQuality(100))
    if err != nil {
        log.Printf("Error encoding file %s: %v", outputPath, err)
        return
    }

    log.Printf("Converted %s to %s", inputPath, outputPath)
} 
