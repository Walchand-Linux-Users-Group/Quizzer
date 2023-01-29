package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type options struct{
	A string
	B string
	C string
	D string
}
type dtype struct {
	Question string
	Options options
	Answer string
}

func questionPuller(url string) ([]dtype){
	resp , err := http.Get(url)


	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)


	if err != nil {
		fmt.Println(err)
	}

	var data []dtype

	json.Unmarshal(body, &data)
	return data;

}