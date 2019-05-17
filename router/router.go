package router

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{
		"name":    "Cakazies",
		"old":     19,
		"address": "Surabaya",
		"gender":  "Laki - Laki",
	}
	return data
}

func Home(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data2 := map[string]interface{}{
		"name1": "Cat",
		"name2": "Cow",
		"name3": "Dragon",
	}
	return data2
}
