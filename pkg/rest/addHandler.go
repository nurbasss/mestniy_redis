package rest

import (
	"net/http"

	"github.com/nurbasss/mestniy_redis/pkg/add"
)

func AddHandler(addService add.AddService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Please provide key name", http.StatusBadRequest)
			return
		}
		val := r.URL.Query().Get("value")
		if val == "" {
			http.Error(w, "Please provide value", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		addService.Set(key, val)
		w.WriteHeader(http.StatusOK)
	}
}
