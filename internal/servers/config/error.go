package config

import "errors"

var (
	ErrNoServerHost = errors.New("no server host provided")
	ErrNoServerPort = errors.New("no server port provided")
)
