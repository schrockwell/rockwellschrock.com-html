package resizer

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func Run(siteDirPath string) {
	// Recursively walk through input directory
	filepath.Walk(siteDirPath, func(sourcePath string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !isImage(sourcePath) || isRenamed(sourcePath) {
			return nil
		}

		srcDir := filepath.Dir(sourcePath)
		destBasename, err := findNextImageBasename(srcDir)
		if err != nil {
			log.Fatal(err)
		}

		destExt := normalizeExt(sourcePath)
		destFilename := destBasename + destExt
		destPath := filepath.Join(srcDir, destFilename)

		fmt.Printf("%s -> %s\n", sourcePath, destPath)

		resizeImage(sourcePath, destPath)
		err = os.Remove(sourcePath)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})
}

func isImage(path string) bool {
	return normalizeExt(path) == ".jpg"
}

func isRenamed(path string) bool {
	basename := filepath.Base(path)
	isMatch, _ := regexp.MatchString(`image[\d]+`, basename)
	return isMatch
}

func findNextImageBasename(dir string) (string, error) {
	basenames, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(basenames); i++ {
		ext := filepath.Ext(basenames[i])
		basenames[i] = strings.TrimSuffix(filepath.Base(basenames[i]), ext)
	}

	for i := 1; i < 1000; i++ {
		basename := "image" + fmt.Sprintf("%03d", i)

		// Look for the existing basename
		j := 0
		for j = 0; j < len(basenames); j++ {
			if basenames[j] == basename {
				break
			}
		}

		// If we reached the end of the list, then this basename is free
		if j == len(basenames) {
			return basename, nil
		}
	}

	return "", errors.New("Couldn't find a free image basename")
}

func normalizeExt(path string) string {
	ext := strings.ToLower(filepath.Ext(path))

	if ext == ".jpeg" {
		return ".jpg"
	} else {
		return ext
	}
}

// Source: https://chat.openai.com/share/d598dc7c-75b8-471d-847b-77cff0b79d14
func resizeImage(inputPath string, outputPath string) {
	// 1520px is the max-width of the site (760px, see main.css) at 2x resolution
	cmd := exec.Command("convert", inputPath, "-resize", "1520x1520>", "-quality", "80", outputPath)

	// Capture command output and error
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Command output:", string(cmdOutput))
		os.Exit(1)
	}
}
