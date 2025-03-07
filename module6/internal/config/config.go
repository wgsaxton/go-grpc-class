package config

// `{
//     "methodConfig": [
//       {
//         "name": [{"service": "config.ConfigService", "method": "LongRunning"}],
//         "timeout": "1s"
//       }
//     ]
//   }`

type Config struct {
	MethodConfig []MethodConfig `json:"methodConfig,omitempty"`
}

type MethodConfig struct {
	Name        []NameConfig `json:"name,omitempty"`
	RetryPolicy RetryPolicy  `json:"retryPolicy,omitempty"`
	Timeout     string       `json:"timeout,omitempty"`
}

type RetryPolicy struct {
	MaxAttempts          int      `json:"maxAttempts,omitempty"`
	InitialBackoff       string   `json:"initialBackoff,omitempty"`
	MaxBackoff           string   `json:"maxBackoff,omitempty"`
	BackoffMultiplier    int      `json:"backoffMultiplier,omitempty"`
	RetryableStatusCodes []string `json:"retryableStatusCodes,omitempty"`
}

type NameConfig struct {
	Service string `json:"service,omitempty"`
	Method  string `json:"method,omitempty"`
}
