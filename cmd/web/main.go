package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get-updates", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			fmt.Fprint(w, "unsupported method")
		}
		var bodyBytes []byte
		var err error
		if req.Body != nil {
			bodyBytes, err = ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Printf("Body reading error: %v", err)
				return
			}
			defer req.Body.Close()
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
	})
	log.Printf("About to listen on 9000. Go to https://notionassistant.publicvm.com/")
	err := http.ListenAndServe(":9000", nil)
	log.Fatal(err)
}
