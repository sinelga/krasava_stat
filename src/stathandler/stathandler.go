package main

import (
//	"database/sql"
	"flag"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"log"
//	"log/syslog"
//	"startones"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

//	if *versionFlag {
//		fmt.Println("Version:", APP_VERSION)
//	}
//	golog, err := syslog.New(syslog.LOG_ERR, "golog")
//
//	defer golog.Close()
//	if err != nil {
//		log.Fatal("error writing syslog!!")
//	}
//
//	startpar := startones.Start(*golog)
//
//	fmt.Println(startpar[0])
//	
//
//	db, err := sql.Open("mysql", startpar[0]+":"+startpar[1]+"@/rellinkweb_development")
//
//	rows, err := db.Query(" select number from phones where number like '0600555%'")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var name string
//		if err := rows.Scan(&name); err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("%s \n", name)
//	}
//	if err := rows.Err(); err != nil {
//		log.Fatal(err)
//	}

}
