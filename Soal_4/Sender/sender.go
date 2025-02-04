package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

// Function to connect to RabbitMQ with retry mechanism
func connectRabbitMQWithRetry(retries int, delay time.Duration) (*amqp.Connection, *amqp.Channel, error) {
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	for i := 0; i < retries; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				_, err = ch.QueueDeclare("my-queue", true, false, false, false, nil)
				if err == nil {
					return conn, ch, nil
				}
			}
		}

		fmt.Printf("🔄 Retry %d/%d - Gagal koneksi: %v\n", i+1, retries, err)
		time.Sleep(delay)
	}

	return nil, nil, fmt.Errorf("gagal koneksi ke RabbitMQ setelah %d percobaan", retries)
}

// Function to send messages with retry mechanism
func sendMessageWithRetry(ch *amqp.Channel, message string, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = ch.Publish(
			"",         // Exchange
			"my-queue", // Routing key
			false,      // Mandatory
			false,      // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err == nil {
			return nil 
		}

		delay := time.Duration(math.Pow(2, float64(i))) * time.Second
		fmt.Printf("⚠️ Gagal mengirim pesan. Retry %d/%d dalam %v...\n", i+1, maxRetries, delay)
		time.Sleep(delay)
	}

	return fmt.Errorf("gagal mengirim pesan setelah %d percobaan", maxRetries)
}

func main() {
	fmt.Println("Tunggu sebentar..")

	conn, ch, err := connectRabbitMQWithRetry(5, 2*time.Second) 
	if err != nil {
		log.Fatalf("❌ Error koneksi ke RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	fmt.Println("📩 Ketik pesan dan tekan Enter (ketik 'exit' untuk keluar)")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(text) == "exit" {
			fmt.Println("👋 Keluar dari sender...")
			break
		}

		err := sendMessageWithRetry(ch, text, 5) // Coba kirim pesan 5 kali
		if err != nil {
			log.Printf("❌ Gagal mengirim pesan: %v", err)
		} else {
			fmt.Println("✅ Pesan terkirim!")
		}
	}
}
