package config

type MySQLAuth struct {
	Address  string `env:"AUTH_MYSQL_ADDRESS" envDefault:"localhost:3306"`
	Username string `env:"AUTH_MYSQL_USERNAME" envDefault:"root"`
	Password string `env:"AUTH_MYSQL_PASSWORD" envDefault:"root"`
	Database string `env:"AUTH_MYSQL_DATABASE" envDefault:"auth"`
}

type MongoAuth struct {
	Host       string `env:"AUTH_MONGO_HOST" envDefault:"localhost"`
	Port       string `env:"AUTH_MONGO_PORT" envDefault:"27017"`
	Username   string `env:"AUTH_MONGO_USERNAME" envDefault:""`
	Password   string `env:"AUTH_MONGO_PASSWORD" envDefault:""`
	Database   string `env:"AUTH_MONGO_DATABASE" envDefault:"auth"`
	Collection string `env:"AUTH_MONGO_COLLECTION" envDefault:"users"`
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
	Auth AuthTopics
}

type AuthTopics struct {
	Registered    string `env:"STREAMING_TOPIC_AUTH_REGISTERED"`
	LoggedIn      string `env:"STREAMING_TOPIC_AUTH_LOGGED_IN"`
	LoginFailed   string `env:"STREAMING_TOPIC_AUTH_LOGIN_FAILED"`
	TokenExtended string `env:"STREAMING_TOPIC_AUTH_TOKEN_EXTENDED"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type App struct {
	Protocol  string `env:"PROTOCOL" envDefault:"http"`
	MySQLAuth MySQLAuth
	MongoAuth MongoAuth
	Server    Server
	Cors      Cors
	I18n      I18n
	Topics    Topics
	Nats      Nats
}
