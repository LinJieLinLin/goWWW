package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	dir:= "www"
	getDir:=GetConfig("WWW")
	if ""!=getDir.(string){
		dir = getDir.(string)
	}
	http.Handle("/", http.FileServer(http.Dir(dir)))
	srvPort := GetConfig("PORT")
	if err := http.ListenAndServe("0.0.0.0:"+srvPort.(string), nil); nil != err {
		fmt.Println(err)
	}
}

var configMap map[string]interface{} = make(map[string]interface{})
var configLock *sync.RWMutex = new(sync.RWMutex)

func ServerConfig() {
	project_dir, _ := os.Getwd()
	project_dir = strings.Replace(strings.SplitAfterN(project_dir, "src", -1)[0], "/src", "", 1)
	project_dir = strings.Replace(project_dir, "\\src", "", 1)
	fmt.Println("project dir is ", project_dir)
	configMap["PROJECT_DIR"] = project_dir
	if config_byte, err := ioutil.ReadFile(project_dir + "/c.properties"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(json.Unmarshal(config_byte, &configMap))
	}
	fmt.Println("server config:")
	for k, v := range configMap {
		fmt.Printf("%s:%v \n", k, v)
	}

}

func GetConfig(k string) interface{} {
	configLock.RLock()
	defer configLock.RUnlock()
	v := configMap[k]
	return v
}

func SetConfig(k string, v interface{}) {
	configLock.Lock()
	defer configLock.Unlock()
	configMap[k] = v
}

func init() {
	ServerConfig()
}
