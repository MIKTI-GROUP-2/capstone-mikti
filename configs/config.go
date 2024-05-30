package configs

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type ProgrammingConfig struct {
	Server     int
	DBPort     int
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	Email      string
	Password   string
	Secret     string
	RefSecret  string
	Cloud_URL  string
}

func InitConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	res, errorRes := loadConfig()

	logrus.Error(errorRes)
	if res == nil {
		logrus.Error("Config : Cannos start program, Failed to load configuration")
		return nil
	}
	return res
}

func readData() *ProgrammingConfig {
	var data = new(ProgrammingConfig)
	data, _ = loadConfig()

	if data == nil {
		err := godotenv.Load(".env")
		data, errorData := loadConfig()

		fmt.Println(errorData)

		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func loadConfig() (*ProgrammingConfig, error) {
	var errorLoad error
	var res = new(ProgrammingConfig)
	var permit = true

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid Port Value, ", err.Error())
			permit = false
		}
		res.Server = port
	} else {
		permit = false
		errorLoad = errors.New("SERVER PORT UNDEFINED")
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	} else {
		permit = false
		errorLoad = errors.New("DBHOST UNDEFINED")
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid DB Port Value, ", err.Error())
			permit = false
		}
		res.DBPort = port
	} else {
		permit = false
		errorLoad = errors.New("DBPORT UNDEFINED")
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	} else {
		permit = false
		errorLoad = errors.New("DBNAME UNDEFINED")
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	} else {
		permit = false
		errorLoad = errors.New("DBUSER UNDEFINED")
	}

	if val, found := os.LookupEnv("DBPASSWORD"); found {
		res.DBPassword = val
	} else {
		permit = false
		errorLoad = errors.New("DBPASSWORD UNDEFINED")
	}

	if val, found := os.LookupEnv("EMAIL"); found {
		res.Email = val
	} else {
		permit = false
		errorLoad = errors.New("EMAIL UNDEFINED")
	}

	if val, found := os.LookupEnv("PASSWORD"); found {
		res.Password = val
	} else {
		permit = false
		errorLoad = errors.New("PASSWORD UNDEFINED")
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	} else {
		permit = false
		errorLoad = errors.New("SECRET UNDEFINED")
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefSecret = val
	} else {
		permit = false
		errorLoad = errors.New("REFSECRET UNDEFINED")
	}

	if val, found := os.LookupEnv("CLOUDURL"); found {
		res.Cloud_URL = val
	} else {
		permit = false
		errorLoad = errors.New("CLOUDURL UNDEFINED")
	}

	if !permit {
		return nil, errorLoad
	}

	return res, nil

}
