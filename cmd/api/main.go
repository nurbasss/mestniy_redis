package main

import (
	"github.com/nurbasss/mestniy_redis/pkg/add"
	"github.com/nurbasss/mestniy_redis/pkg/get"
	"github.com/nurbasss/mestniy_redis/pkg/recover"
	"github.com/nurbasss/mestniy_redis/pkg/rest"
	"github.com/nurbasss/mestniy_redis/pkg/save"
	"github.com/nurbasss/mestniy_redis/pkg/store"

	"fmt"
	"log"
	"net/http"
)

func main() {
	mestniyStore := store.NewStore()
	addService := add.NewService(mestniyStore)
	getService := get.NewService(mestniyStore)
	saveService := save.NewService(mestniyStore)
	recover.RecoverService.Recover(mestniyStore)
	mux := http.NewServeMux()
	go save.StartTask(saveService)
	//Can not implement route like in task so please use:
	
	//localhost:4200/put?key=keyname&value=valuename
	//localhost:4200/get?key=keyname

	/*also it save data into local json, so even after server terminating, 
	you will be able to retrieve data from repo*/

	mux.HandleFunc("/put", rest.AddHandler(addService))
	mux.HandleFunc("/get", rest.GetHandler(getService))
	fmt.Println("Server running in port 4200")
	log.Fatal(http.ListenAndServe(":4200", mux))
}