package c

import "time"

var (
	JWT_EXPIRED_DURATION = 3 * 24 * time.Hour
	JWT_ISSUER           = "citizen-v"
	JWT_SUBJECT          = "subject"
)
