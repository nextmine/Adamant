package main

import (
	"Adamant/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	serverPort = "19132"
)

func main() {
	setupLogging()

	log.Println("Запуск сервера...")

	srv, err := server.New(serverPort)
	if err != nil {
		log.Fatalln("Ошибка при создании сервера:", err)
	}

	go srv.Run()

	log.Printf("Сервер запущен и готов к работе на порту %s\n", serverPort)

	waitForShutdownSignal()

	log.Println("\nЗавершение работы сервера...")
}

func setupLogging() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func waitForShutdownSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}
