package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var urlStore = make(map[string]string)

func generateShortCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the URL: ")
	url, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error occured while reading URL %s", err)
	}
	url = strings.TrimSpace(url)
	shortcode := generateShortCode(6)
	urlStore[shortcode] = url
	fmt.Printf("Shortened URL: http://localhost:8080/%s\n", shortcode)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := strings.TrimPrefix(r.URL.Path, "/")
		log.Println(code)
		original, ok := urlStore[code]
		if ok {
			http.Redirect(w, r, original, http.StatusFound)
		} else {
			http.NotFound(w, r)
		}
	})
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	// fmt.Printf("You entered:%s", shortcode)
}
