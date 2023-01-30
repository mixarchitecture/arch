package config

type MySQLExample struct {
	Address  string `env:"EXAMPLE_MYSQL_ADDRESS" envDefault:"localhost:3306"`
	Username string `env:"EXAMPLE_MYSQL_USERNAME" envDefault:"root"`
	Password string `env:"EXAMPLE_MYSQL_PASSWORD" envDefault:"root"`
	Database string `env:"EXAMPLE_MYSQL_DATABASE" envDefault:"boilerplate"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Server struct {
	Host string `env:"SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"SERVER_PORT" envDefault:"3000"`
}

type Cors struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
}

type Topics struct {
	Example ExampleTopics
}

type ExampleTopics struct {
	Created string `env:"STREAMING_TOPIC_EXAMPLE_CREATED"`
	Updated string `env:"STREAMING_TOPIC_EXAMPLE_UPDATED"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type App struct {
	Protocol     string `env:"PROTOCOL" envDefault:"http"`
	MySQLExample MySQLExample
	Server       Server
	Cors         Cors
	I18n         I18n
	Topics       Topics
	Nats         Nats
}
