package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	bj38_cli "bj38_cli/proto/bj38"
	"github.com/micro/go-micro/v2/client"
)

func Bj38_cliCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bj38_cliCall收到了")
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("request解析成功")

	// call the backend service
	bj38_cliClient := bj38_cli.NewBj38Service("go.micro.service.bj38", client.DefaultClient)
	rsp, err := bj38_cliClient.Call(context.TODO(), &bj38_cli.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("回调成功")

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}
	fmt.Println("response：",response)
	fmt.Println("response：",response["msg"])

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("response json解析成功")

}
