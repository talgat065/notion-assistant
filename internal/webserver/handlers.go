package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeGetUpdatesHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			fmt.Fprint(w, "unsupported method")
		}
		var bodyBytes []byte
		var err error
		if r.Body != nil {
			bodyBytes, err = ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Body reading error: %v", err)
				return
			}
			defer r.Body.Close()
		}

		if len(bodyBytes) > 0 {
			var prettyJSON bytes.Buffer
			if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
				fmt.Printf("JSON parse error: %v", err)
				return
			}
			fmt.Println(string(prettyJSON.Bytes()))
		} else {
			fmt.Printf("Body: No Body Supplied\n")
		}
	}
}
