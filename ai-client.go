// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// )

// func contactAI(url string, ollamaReq Request) (*Response, error) {
// 	js, err := json.Marshal(&ollamaReq)
// 	if err != nil {
// 		return nil, err
// 	}
// 	client := http.Client{}
// 	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(js))
// 	if err != nil {
// 		return nil, err
// 	}
// 	httpResp, err := client.Do(httpReq)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer httpResp.Body.Close()
// 	ollamaResp := Response{}
// 	err = json.NewDecoder(httpResp.Body).Decode(&ollamaResp)
// 	return &ollamaResp, err
// }
