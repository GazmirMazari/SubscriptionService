package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	// Session holds the session secret
	Session   *scs.SessionManager
	DB        *sql.DB
	Info      *log.Logger
	Error     *log.Logger
	WaitGroup *sync.WaitGroup
}
