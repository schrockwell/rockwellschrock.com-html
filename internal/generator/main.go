package generator

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	siteDirPath = "./web"
	outDirPath  = "./_site"
)

func Run() {
	// Ensure site dir exists
	if _, err := os.Stat(siteDirPath); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Directory '%s' not found", siteDirPath)
	}

	// Parse all templates in /site/templates/*
	templatesGlob := filepath.Join(siteDirPath, "_templates", "*")
	templates, err := template.ParseGlob(templatesGlob)
	if err != nil {
		log.Fatal(err)
	}

	// Wipe _site directory
	err = os.RemoveAll(outDirPath)
	if err != nil {
		log.Fatal(err)
	}

	// Recursively walk through site directory
	filepath.Walk(siteDirPath, func(sourcePath string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		relPath, _ := filepath.Rel(siteDirPath, sourcePath)
		if strings.HasPrefix(relPath, "_") {
			return nil
		}

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			log.Fatal(err)
		}

		destPath := filepath.Join(outDirPath, relPath)

		if filepath.Ext(sourcePath) == ".html" {
			executePageTemplate(templates, sourcePath, destPath)
		} else if !fileInfo.IsDir() {
			copyFile(sourcePath, destPath)
		}

		return nil
	})
}

func executePageTemplate(templates *template.Template, sourcePath string, destPath string) {
	bytes, err := os.ReadFile(sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	pageTemplate, err := template.Must(templates.Clone()).New("content").Parse(content)
	if err != nil {
		log.Fatal(err)
	}

	destDirPath := filepath.Dir(destPath)
	err = os.MkdirAll(destDirPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	err = pageTemplate.ExecuteTemplate(destFile, "root.html", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(destPath)
}

func copyFile(sourcePath string, destPath string) {
	destDirPath := filepath.Dir(destPath)
	err := os.MkdirAll(destDirPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	io.Copy(destFile, sourceFile)

	fmt.Println(destPath)
}
