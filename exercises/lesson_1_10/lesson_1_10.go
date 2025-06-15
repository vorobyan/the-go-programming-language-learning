// Найдите веб-сайт, который содержит большое количество данных. Исследуйте работу кеширования путем двукратного запуска f e t c h a l l и сравнения времени запросов. Получаете ли вы каждый раз одно и то же содержимое?
// Измените fetchall так, чтобы вывод осуществлялся в файл и чтобы затем можно
// было его изучить.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	filename := "data.txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		//fmt.Println(<-ch)
		if _, err := file.Write([]byte(<-ch)); err != nil {
			log.Fatal(err)
		}
	}
	final := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	_, err = file.Write([]byte(final))
	if err != nil {
		log.Fatal(err)
	}

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err, "\n")
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
