package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	url := "http://golang.org/doc/gopher/frontpage.png"
	ext := filepath.Ext(url)
	filename := fmt.Sprintf("teste%s", ext)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
}
