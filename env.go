package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func GetValue() {
	name := os.Getenv("NAME")

	fmt.Printf("env NAME is :%s ", name)

	viper.BindEnv("id")

	viper.SetDefault("id", 13)

	id := viper.GetInt("id")

	fmt.Println(id)

	os.Setenv("ID", "50")

}
