/*
 * Simple Inventory API
 *
 * This is API for cloud-based microservice
 *
 * API version: 1.0.0
 * Contact: newmail@mail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DeleteAllMetrics(w http.ResponseWriter, r *http.Request) {
	deleteAll(METRICS)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteMetricById(w http.ResponseWriter, r *http.Request) {
	type ViewData struct {
		Id string
	}
	data := ViewData{
		Id: r.URL.Query().Get("id"),
	}

	deleteFromDB(data.Id, METRICS)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetMetricById(w http.ResponseWriter, r *http.Request) {
	type ViewData struct {
		Id string
	}
	data := ViewData{
		Id: r.URL.Query().Get("id"),
	}

	type Response struct {
		Id   string
		User int
		App string
		Text string
	}

	tmp := get(data.Id, METRICS)
	text, err1 := ioutil.ReadFile(tmp.Path)
	if err1 != nil {
		panic(err1)
	}

	responseRaw := Response{
		Id:   data.Id,
		User: tmp.UserId,
		App: tmp.AppId,
		Text: string(text),
	}

	response, err2 := json.Marshal(responseRaw)
	if err2 != nil {
		panic(err2)
	}
	w.Write(response)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PostMetric(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	put(string(body), METRICS)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
