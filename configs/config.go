package configs

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type CM_secrets struct {
	Base_Url       string `yaml:"base_url"`
	Version        string `yaml:"version"`
	Cm_user        string `yaml:"cm_user"`
	Cm_password    string `yaml:"cm_password"`
	Encryption_key string `yaml:"encryption_key"`
}

type AKeyless_secrets struct {
	AkeylessUsername string `yaml:"username"`
	AkeylessPassword string `yaml:"password"`
}

type Config struct {
	Env              string `yaml:"env"`
	CM_secrets       `yaml:"cm_secret"`
	AKeyless_secrets `yaml:"akeyless_secret"`
}

// MustLoad function reads the file, loads the data and store the fields in the Config struct
func MustLoad() *Config {

	configPath := getConfigPath()

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Make an object of the Config struct
	var cnfg Config

	// Unmarshal the fields inside the yaml file and store in Config struct
	if err = yaml.Unmarshal(file, &cnfg); err != nil {
		log.Fatal(err)
	}

	return &cnfg

}

// getConfigPath function helps to verify and load the yaml file which store our secrets being used in this application.
func getConfigPath() string {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	// Defining a new FlagSet on the flag package will "erase" the previous flag set defined by the imported packages
	// and avoid panic: flag redefined: config error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	if configPath == "" {
		flags := flag.String("configfile", "", "place a config file path name")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path not set")

		}
	}

	_, err := os.Stat("configPath")
	if err != nil {
		os.IsNotExist(err)
	}

	return configPath
}
