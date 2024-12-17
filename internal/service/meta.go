package service

import "time"

type metadata struct {
	username string
	cookie   string
	delta    time.Duration
}

func (m *metadata) GetUsername() string {
	return m.username
}

func (m *metadata) GetCookie() string {
	return m.cookie
}

func (m *metadata) GetDelta() time.Duration {
	return m.delta
}
