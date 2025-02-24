package server

import (
	"net/http"

	"github.com/ujooju/tlsCommutator/config"
	"golang.org/x/crypto/acme/autocert"
)

func SetUpServer() {
	config.GlobalConfig.SetUpConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", DefaultCommutator)
	http.Serve(autocert.NewListener(config.GlobalConfig.PublicUrl), mux)
}
