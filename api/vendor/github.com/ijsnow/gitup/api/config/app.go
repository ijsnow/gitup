package config

import (
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

// App holds all config information for the app. It is populated at
// program initialization time, so its config options can also be used
// at initialization time in any packages that depend on it.
var App struct {
	ENV string `env:"GO_ENV"`

	Server struct {
		Port string `env:"PORT"`
	}

	Database struct {
		Path string `env:"DB_PATH"`
	}

	Security struct {
		Secret string `env:"SECTRET"`
	}

	RepoDir string `env:"REPO_DIR"`
}

func init() {
	// Will get overwritten if production env is set
	App.ENV = "develop"

	setFromEnv(reflect.ValueOf(&App))
}

func setFromEnv(v reflect.Value) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	st := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch k := field.Kind(); k {
		case reflect.Struct:
			setFromEnv(field)
		case reflect.String:
			if key := st.Field(i).Tag.Get("env"); key != "" {
				field.SetString(os.Getenv(key))
			}
		default:
			log.Panicf("Don't know how to handle field '%s' of kind %v\n", st.Field(i).Name, k)
		}
	}
}
