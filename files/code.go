package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func CopyFromSourceToDest(source, dest string) error {
	info, err := os.Stat(source)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Не найдено файла с именем: ", source)
		} else {
			fmt.Println("ОШибка при работе с файлом", source, ", проверьте файл на валидность и достаточные права")
		}
		return fmt.Errorf("Ошибка при работе с файлом источником!")
	}
	if info.IsDir() {
		return fmt.Errorf("%s является директорией, а не файлом", source)
	}
	data, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("Ошибка открытия файла с именем : %s", source)
	}
	if _, err := os.Stat(dest); err != nil {
		fmt.Println("Принимающего файла не существует.будет создан файл...")
	}
	fmt.Println("Копирование ", source, "->", dest, "...")
	err2 := os.WriteFile(dest, data, 0644)
	if err2 != nil {
		return fmt.Errorf("Ошибка записи в файл с именем : %s", dest)
	}
	fmt.Println("Копирование прошло успешно")
	return nil
}

func ReadSomeBytesUSingBuffer(filename string, bytes int) (string, error, int) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("Не существует указанного файла\n"), 0
		} else {
			return "", err, 0
		}
	}
	defer file.Close()
	buf := make([]byte, bytes)
	readBytes, err := file.Read(buf)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Прочитано", readBytes, "байт данных, прочитано до конца")
			return string(buf[:readBytes]), nil, readBytes
		} else {
			return "", fmt.Errorf("Ошибка при работе с буфером!\n"), 0
		}
	}
	fmt.Println("Прочитано", readBytes, "байт данных! В файле содержатся еще данные!")
	return string(buf), nil, readBytes
}

func ReadAndWriteFileUSingBufioScannedAndWriter(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("Файл не найден\n")
		}
		return fmt.Errorf("Ошибка при открытии файла!")
	}
	defer file.Close()
	Scanner := bufio.NewScanner(file)

	var i int
	for Scanner.Scan() {
		i++
		line := Scanner.Text()
		fmt.Println("Строка ", i, ": "+line)
	}
	fmt.Println()

	Writer := bufio.NewWriter(file)
	Writer.WriteString("new new")
	Writer.Flush()
	newData, myerror, countRead := ReadSomeBytesUSingBuffer(filename, 100)
	if myerror != nil {
		return fmt.Errorf("Ошибка при вызове функции ReadSomeBytesUSingBuffer внутри метода")
	}
	fmt.Println("\nВот прочитанные", countRead, "байтов файла после перезаписи:\n"+string(newData))
	return nil
}

func main() {
	fmt.Println("КОПИРОВАНИЕ СОДЕРЖИМОГО ФАЙЛА В ДРУГОЙ ФАЙЛ")
	fmt.Println(strings.Repeat("-", 100))
	err1 := CopyFromSourceToDest("files/SourceTest", "files/DestinationTest")
	if err1 != nil {
		panic("Ошибки при работе с файлами, проверьте все как нужно")
	}
	fmt.Println()

	fmt.Println("РАБОТА С БУФЕРОМ УКАЗАННОГО РАЗМЕРА")
	fmt.Println(strings.Repeat("-", 100))
	info, err, _ := ReadSomeBytesUSingBuffer("files/DestinationTest", 5)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Содержимое прочитанного буфера:", info)
	fmt.Println()

	fmt.Println("РАБОТА СО СКАНЕРОМ И ВРАЙТЕРОМ ИЗ bufio")
	fmt.Println(strings.Repeat("-", 100))
	err = ReadAndWriteFileUSingBufioScannedAndWriter("files/TestBufio")
	if err == nil {
		fmt.Println("\nВсё хорошо, функция отработала нужным образом, ошибок не возвращалось")
	} else {
		fmt.Println("К сожалению последняя функция не сработала должным обазом...")
		fmt.Print(err.Error())
	}

}
