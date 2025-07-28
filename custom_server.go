package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	port := "3999"
	presentPort := "3998" // Go present будет работать на этом порту

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Запускаем Go present на внутреннем порту
	fmt.Println("🚀 Starting Go present server on internal port", presentPort)
	cmd := exec.Command("go", "run", "golang.org/x/tools/cmd/present", "-http=:"+presentPort, "-play=false")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go func() {
		if err := cmd.Run(); err != nil {
			log.Printf("Go present server error: %v", err)
		}
	}()

	// Даем время Go present серверу запуститься
	time.Sleep(2 * time.Second)

	// Создаем прокси к Go present серверу
	presentURL, _ := url.Parse("http://localhost:" + presentPort)
	proxy := httputil.NewSingleHostReverseProxy(presentURL)

	// Создаем кастомный обработчик
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Если запрашивается корень, показываем наш index.html
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}

		// Если запрашивается наш кастомный контент (presentations/), обрабатываем сами
		if strings.HasPrefix(r.URL.Path, "/presentations/") {
			http.ServeFile(w, r, "."+r.URL.Path)
			return
		}

		// Для всех остальных запросов используем прокси к Go present
		proxy.ServeHTTP(w, r)
	})

	fmt.Printf("🌐 Custom presentation server started on http://localhost:%s\n", port)
	fmt.Println("")
	fmt.Println("📖 Access presentations:")
	fmt.Printf("   🌐 Language Selection: http://localhost:%s/\n", port)
	fmt.Printf("   🇺🇸 English:           http://localhost:%s/presentations/en/\n", port)
	fmt.Printf("   🇷🇺 Russian:           http://localhost:%s/presentations/ru/\n", port)
	fmt.Println("")
	fmt.Println("🛑 Press Ctrl+C to stop the server")
	fmt.Println("")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
