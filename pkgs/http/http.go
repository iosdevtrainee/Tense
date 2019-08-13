package http

import (
	"net/http"
	"time"

	"../internal/blog"
	"../internal/comment"
)

type Server interface {
	Start(cfg Config, p blog.Service, c comment.Service /* t thread.Service, u user.Service */) *http.Server
}

type Config struct {
	Port         uint
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}
