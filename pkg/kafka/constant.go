package kafka

import "time"

const (
	MIN_BYTES                = 10e3 // 10KB
	MAX_BYTES                = 10e6 // 10MB
	QUEUE_CAPACITY           = 100
	HEARTBEAT_INTERVAL       = 3 * time.Second
	COMMIT_INTERVAL          = 0
	PARTITION_WATCH_INTERVAL = 5 * time.Second
	MAX_ATTEMPTS             = 3
	DIAL_TIMEOUT             = 3 * time.Minute
	MAX_WAIT                 = 1 * time.Second
	WRITER_READ_TIMEOUT      = 10 * time.Second
	WRITER_WRITE_TIMEOUT     = 10 * time.Second
	WRITER_REQUIRED_ACKS     = -1
	WRITER_MAX_ATTEMPTS      = 3
)
