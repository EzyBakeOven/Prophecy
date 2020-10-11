package setup

import (
	"fmt"

	"github.com/spf13/viper"
)

// ProcessConfig : Processes the config file so a world can be generated.
func ProcessConfig() {
	viper.SetConfigName("simConfig") // name of config file (without extension)
	viper.AddConfigPath(".")         // look for config in the working directory
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("Welcome to", viper.Get("Realm"))
}
