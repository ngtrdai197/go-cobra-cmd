.PHONY: worker client producer consumer

# Delevery email client
client:
	go run main.go client-cmd

# Delevery email worker
worker:
	go run main.go worker-cmd

producer:
	go run main.go kafka-producer-cmd

consumer:
	go run main.go kafka-consumer-cmd