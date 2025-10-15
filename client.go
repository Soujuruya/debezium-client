package debezium_client

import (
	"net/http"
	"strings"
	"time"
)

//Любой HTTP-запрос должен быть покрыт таймаутами

type Client struct {
	cc      *http.Client
	baseURL string
}

func New(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		cc:      &http.Client{Timeout: timeout},
	}
}
