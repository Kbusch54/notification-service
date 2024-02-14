package config

type Config struct {
	Persistence Persistence `mapstructure:"persistence"`
	Services    Services    `mapstructure:"services"`
	Server      Server      `mapstructure:"server"`
	Stream      Stream      `mapstructure:"stream"`
}

type Persistence struct {
	MongoDB MongoDB `mapstructure:"mongodb"`
}

type MongoDB struct {
	URL      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
}

type Services struct {
	Brevo Brevo `mapstructure:"brevo"`
}

type InvestingServices struct {
	Host string `mapstructure:"host"`
}

type Server struct {
	Port string `mapstructure:"port"`
}
type Stream struct {
	Kafka Kafka `mapstructure:"kafka"`
}

type Kafka struct {
	Brokers []string `mapstructure:"brokers"`
}

type Brevo struct {
	APIKey string
	Email  string
	Host   string
	Port   int
}
