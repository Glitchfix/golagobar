package server

import (
	"fmt"

	"github.com/Glitchfix/golagobar/config"
)

// Init - Initialize server
func Init() {

	r := NewRouter()
	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	r.Run(address)
}
