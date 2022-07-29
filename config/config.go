package config

//в файл COnfig клдаутся данные, которые могут изменится например название БД
// ПОРТА и пр.
type Config struct {
	// поля конфигурации
	// в конкретном случае мы укажем всё, что в строке dsn

	DBHost     string `env:"Host"`
	DBUser     string `env:"User"`
	DBPassword string `env:"Password"`
	DBName     string `env:"Name"`
	DBPort     string `env:"Port"`
	SSLmode    string `env:"SSLmode"`
}
