package client

import (
	"fmt"
	"net/http"
)

var BasePath = "http://46.101.42.27:8080/"

func GetAll() string {
	res := ""
	resp, err := http.Get(BasePath + "all")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		res = res + string(bs[:n])
		//fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
	fmt.Println(res)
	return res
}

func DeleteAll() {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, BasePath+"all", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
}
