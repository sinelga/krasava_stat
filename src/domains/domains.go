package domains

import (

)

type Config struct {
	Database struct {
		ConStr string
	}
	Store struct {
		StoreDir string
	}
	Redis struct {
		Prot string
		Host string
	}
}