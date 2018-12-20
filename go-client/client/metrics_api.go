package client

import (
	"fmt"
	"net/http"
	"os"
)

func GetAppStatMetrics(app string) string {
	res := ""
	resp, err := http.Get(BasePath + "metrics/getstat?id=" + app)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		res += string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return res
}

func GetMetricsById(id string) {
	resp, err := http.Get(BasePath + "metrics?id=" + id)
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

func DeleteMetricsById(id string) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, BasePath+"metrics?id="+id, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v %v", req.URL, req.Method)

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

func PostMetrics(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	resp, err := http.Post(BasePath+"metrics", "text/plain", file)
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

func DeleteAllMetrics() {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, BasePath+"metrics/deleteAll", nil)
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
