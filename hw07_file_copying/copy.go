package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrOpenFile              = errors.New("error opening file")
)

func Copy(fromPath, toPath string, offset, limit int64) error {

	srcF, err := os.OpenFile(fromPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return ErrOpenFile
	}
	defer srcF.Close()

	fi, _ := srcF.Stat()
	if fi.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if fi.Size() == 0 {
		return ErrUnsupportedFile
	}

	if limit+offset > fi.Size() || limit == 0 {
		limit = fi.Size() - offset
	}
	cleanF := serialization(srcF)

	dstF, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer dstF.Close()

	// start new bar
	bar := progressbar.DefaultBytes(
		limit,
		"copying",
	)

	_, err = io.CopyN(io.MultiWriter(dstF, bar), bytes.NewReader(cleanF[offset:]), limit)
	if err != nil && err != io.EOF {
		return ErrUnsupportedFile
	}

	// finish bar
	bar.Finish()

	return nil
}

func serialization(f *os.File) []byte {
	fi, _ := f.Stat()
	result := make([]byte, fi.Size())
	_, err := f.Read(result)
	if err != nil {
		return nil
	}
	return bytes.ReplaceAll(result, []byte("\r"), []byte(""))
}

//type CopyData struct {
//	fromPath string
//	toPath   string
//	limit    int64
//	offset   int64
//}
//
//var FullData = []CopyData{
//	CopyData{
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\input.txt",
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\out_offset0_limit0_my.txt",
//		0,
//		0,
//	},
//	CopyData{
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\input.txt",
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\out_offset0_limit10_my.txt",
//		0,
//		10,
//	},
//	CopyData{
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\input.txt",
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\out_offset0_limit1000_my.txt",
//		0,
//		1000,
//	},
//	CopyData{
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\input.txt",
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\out_offset0_limit10000_my.txt",
//		0,
//		10000,
//	},
//	CopyData{
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\input.txt",
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\out_offset100_limit1000_my.txt",
//		100,
//		1000,
//	},
//	CopyData{
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\input.txt",
//		"C:\\Users\\user\\GolandProjects\\otus-go-pro\\hw07_file_copying\\testdata\\out_offset6000_limit1000_my.txt",
//		6000,
//		1000,
//	},
//}

func main() {
	//Чтение аргументов
	from := flag.String("from", "", "Путь к исходному файлу")
	to := flag.String("to", "", "Путь к копии")
	offset := flag.Int64("offset", 0, "Отступ в исходном файле")
	limit := flag.Int64("limit", 0, "Количество копируемых байт (0 = весь файл)")
	flag.Parse()

	// Проверка аргументов
	if *from == "" || *to == "" {
		fmt.Println("Ошибка: Параметры -from и -to обязательны")
		return
	}

	err := Copy(*from, *to, *offset, *limit)
	if err != nil {
		fmt.Println(err)
	}

	//for _, item := range FullData {
	//	err := Copy(item.fromPath, item.toPath, item.limit, item.offset)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//}

}
