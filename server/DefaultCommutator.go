package server

import (
	"net/http"

	"github.com/ujooju/tlsCommutator/config"
)

func DefaultCommutator(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://"+config.GlobalConfig.DestIp+":"+config.GlobalConfig.Port, http.StatusSeeOther)
}
