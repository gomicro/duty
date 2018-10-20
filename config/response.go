package config

// Response represents an http response of a status code and a given payload
type Response struct {
	Code    int    `yaml:"code"`
	Payload string `yaml:"payload"`
}
