package constant

import "time"

// config echo runtime.
const (
	DefaultPort    = 8000
	DefaultTimeout = 10 * time.Second
)

// config header middleware.
const (
	DefaultAllowHeaderAccessToken  = "Authorization"
	DefaultAllowHeaderRefreshToken = "refresh-token"
)
