.PHONY: worker client

# Delevery email client
client:
	go run main.go client-cmd

# Delevery email worker
worker:
	go run main.go worker-cmd