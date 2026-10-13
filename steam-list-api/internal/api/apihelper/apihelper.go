package apihelper

import (
	"net/http"
	"strconv"
)

func GetPage(r *http.Request) (int, error) {
	pageString := r.URL.Query().Get("page")
	if pageString == "" {
		return 1, nil
	}
	return strconv.Atoi(pageString)
}

func InternalServerErrorHandler(w http.ResponseWriter, _ *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error" + err.Error()))
}
