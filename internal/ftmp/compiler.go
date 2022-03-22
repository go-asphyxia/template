package ftmp

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func CompileDir(path string) {
	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	var names []string
	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}

		if !file.IsDir() {
			names = append(names, file.Name())
		} else {
			subPath := filepath.Join(path, file.Name()) 
			CompileDir(subPath)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		if strings.HasSuffix(name, "tmpl") {
			filename := filepath.Join(path, name)
			compilefile(filename)
		}
	}
	dir.Close()
}

func compilefile(filename string) {
	//Итоговый файл в который пишем
	//writeFile := filename + ".go"

	//Сам файл, который компилим
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	//Файл с временными значениями
	outf := filename + ".tmp"
	tmpfile, err := os.Create(outf)
	if err != nil {
		panic(err)
	}

	//Отправляем на парс
	//func ParseV2(w io.Writer, r io.Reader, filename string) error {
	if err = ParseFile(tmpfile, readFile, filename); err != nil {
		panic(err)
	}

	//Закрытие файлов на чтение и запись
	if err = tmpfile.Close(); err != nil {
		panic(err)
	} 

	if err = readFile.Close(); err != nil {
		panic(err)
	}
}