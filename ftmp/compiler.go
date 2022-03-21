package ftmp

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func CompileDir(path string) {
	f, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if !f.IsDir() {
		panic("is not dir")
	}

	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	var names []string
	for _, file := range files {
		name := file.Name()
		if name == "." || name == ".." {
			continue
		}

		if !file.IsDir() {
			names = append(names, name)
		} else {
			subPath := filepath.Join(path, name)
			CompileDir(subPath)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		if strings.HasSuffix(name, ".tmpl") {
			filename := filepath.Join(path, name)
			compileFile(filename)
		}
	}
}

func compileFile(infile string) {
	//Итоговый файл в который пишем
	outfile := infile + ".go"

	//Сам файл, который компилим
	inf, err := os.Open(infile)
	if err != nil {
		panic(err)
	}

	//Файл с временными значениями
	tmpfile := outfile + ".tmp"
	outf, err := os.Create(tmpfile)
	if err != nil {
		panic(err)
	}

	//Получить точный путь к файлу
	packageName, err := getPackageName(infile)
	if err != nil {
		panic(err)
	}

	//Отправляем на парс
	//func ParseV2(w io.Writer, r io.Reader, filename, pkg string) error
	if err = ParseV2(outf, inf, infile, packageName); err != nil {
		panic(err)
	}

	//Закрытие файлов на чтение и запись
	if err = outf.Close(); err != nil {
		panic(err)
	} 

	if err = inf.Close(); err != nil {
		panic(err)
	}
}

func getPackageName(filename string) (string, error) {
	filenameAbs, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}

	dir, _ := filepath.Split(filenameAbs)
	return filepath.Base(dir), err
}