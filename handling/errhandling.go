package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err := WriteFile(filename, "this is WriteFile")
		if err != nil {
			fmt.Println("파일 생성 실패", err)
			return
		}
		line, err := ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기 실패", err)
			return
		} else {
			fmt.Println("파일 내용 : ", line)
		}
	} else {
		fmt.Println("파일 내용 : ", line)
	}
}
