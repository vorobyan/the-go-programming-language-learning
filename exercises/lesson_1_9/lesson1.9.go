// Измените программу fetch так, чтобы она выводила код состояния HTTP, содержащийся в resp.Status.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
