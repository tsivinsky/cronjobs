package env

import "github.com/tsivinsky/goenv"

type env struct {
	DBUser     string `env:"POSTGRES_USER,required"`
	DBPassword string `env:"POSTGRES_PASSWORD,required"`
	DBName     string `env:"POSTGRES_DB,required"`
	DBHost     string `env:"DB_HOST,required"`

	WEB_APP_URL   string `env:"WEB_APP_URL,required"`
	COOKIE_DOMAIN string `env:"COOKIE_DOMAIN,required"`

	GHClientId     string `env:"GITHUB_CLIENT_ID,required"`
	GHClientSecret string `env:"GITHUB_CLIENT_SECRET,required"`
}

var Env = new(env)

func Init() error {
	return goenv.Load(Env)
}
