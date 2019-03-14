package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler returns http.Handler for API endpoint
func Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		// res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// res.Header().Set("Access-Control-Allow-Methods", "PUT")
		fmt.Println(req.RequestURI)
		body, err := json.Marshal(map[string]interface{}{
			"data": "Hello, world dsssh",
		})

		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
		res.Write(body)
	}
}
