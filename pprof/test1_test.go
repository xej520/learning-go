package pprof

import (
	"testing"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func TestAdd(t *testing.T)  {
	go func() {
		for  {
			log.Println(add("https://github.com/eddycjy"))
		}
	}()

	http.ListenAndServe("localhost:8080", nil)
}

func add(str string) string {
	var datas []string

	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}


