package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	for true {
		bs := make([]byte, 1024)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[0:n]))

		if n == 0 || err != nil {
			break
		}
	}

	// Еще можно использовать сокращенную версию без цикла
	//io.Copy(os.Stdout, resp.Body)

	fmt.Println("Done")
}
