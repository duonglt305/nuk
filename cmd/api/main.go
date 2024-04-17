package main

import (
	"duonglt.net/pkg/utils"
)

func main() {
	h := utils.NewHtmlParser("templates")
	println(h.Parse("email/passwords/forgot.html"))
	/*viper.SetConfigFile(".env")
	  if err := viper.ReadInConfig(); err != nil {
	  	log.Printf("failed to read config: %+v\n", err)
	  	os.Exit(1)
	  }
	  r, err := internal.Initialize()
	  if err != nil {
	  	log.Printf("%+v\n", err)
	  	os.Exit(1)
	  }
	  if err := r.ServeHTTP(); err != nil {
	  	log.Printf("failed to start application: %+v\n", err)
	  	os.Exit(1)
	  }*/
}
