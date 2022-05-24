package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 6 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://google.com", nil)

	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)

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
