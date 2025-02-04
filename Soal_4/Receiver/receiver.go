package main

import (
	"fmt"
	"log"
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

func main() {
	fmt.Println("📥 RabbitMQ Receiver - Menunggu pesan...")

	conn, ch, err := connectRabbitMQWithRetry(5, 2*time.Second) // Coba koneksi 5 kali
	if err != nil {
		log.Fatalf("❌ Error koneksi ke RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		"my-queue", // Queue name
		"",         // Consumer
		true,       // Auto-Ack
		false,      // Exclusive
		false,      // No-local
		false,      // No-Wait
		nil,        // Args
	)
	if err != nil {
		log.Fatalf("❌ Gagal menerima pesan: %v", err)
	}

	for msg := range msgs {
		fmt.Println("📩 Pesan diterima:", string(msg.Body))
	}
}
