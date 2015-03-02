package main 

import (
		"log"
	"log/syslog"
	"github.com/rs/cors"
	"github.com/zenazn/goji"

)

func main() {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	goji.Use(c.Handler)
	

}

