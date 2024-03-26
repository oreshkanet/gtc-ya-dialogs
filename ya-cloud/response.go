package main

type Response struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}
