package server

import (
	"net/http"

	"github.com/ujooju/tlsCommutator/config"
)

func DefaultCommutator(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
	http.Redirect(w, r, "http://"+config.GlobalConfig.DestIp+":"+config.GlobalConfig.Port, http.StatusSeeOther)
}
