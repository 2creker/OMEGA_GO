package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Метод login выполнен!.")
}

func verify(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Метод verify выполнен!")
}

func main_page(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OMEGA_GO")
}

func main() {
	http.HandleFunc("/login/", login)
	http.HandleFunc("/verify/", verify)
	http.HandleFunc("/", main_page)

	server := &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка сервера: %s", err)
		}
	}()

	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	log.Println("Завершение работы сервера.")

	if err := server.Shutdown(nil); err != nil {
		log.Fatalf("Ошибка при остановке сервера: %s", err)
	}

	log.Println("Сервер успешно остановлен.")
}
