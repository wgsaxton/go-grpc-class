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
	Method []MethodConfig `json:"methodConfig,omitempty"`
}

type MethodConfig struct {
	Name    []NameConfig `json:"name,omitempty"`
	Timeout string       `json:"timeout,omitempty"`
}

type NameConfig struct {
	Service string `json:"service,omitempty"`
	Method  string `json:"method,omitempty"`
}
