package startones

import (


	"code.google.com/p/gcfg"
	"log"
	"log/syslog"
	"domains"
)

var config domains.Config

func Start() (syslog.Writer, domains.Config) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")	

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	
	err = gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		
		golog.Crit("cannot read configuration file config.ini" + err.Error())

	}
	
	golog.Info(config.Database.ConStr)

	return *golog, config


}
