package rest

import (
	"fmt"
	"net/http"

	"github.com/nurbasss/mestniy_redis/pkg/get"
	"github.com/nurbasss/mestniy_redis/pkg/store"
)

func GetHandler(getService get.GetService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Please provide keyname", http.StatusBadRequest)
			return
		}
		val, err := getService.Get(key)
		if err == store.ErrNotFound {
			http.Error(w, "No data", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("Error when getting data from repo: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}
}
