package helper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/spf13/viper"
)

type payload struct {
	Text string `json:"text"`
}

//TeamsLog webHook send msg to Teams頻道
func TeamsLog(msg string, dest string) {
	payloads := payload{msg}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	ba, _ := json.Marshal(payloads)
	resp, _ := http.Post(viper.GetString("TeamsHook."+dest), "application/json", bytes.NewBuffer([]byte(ba)))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("[Helper] Notify teams channel response with result: %s\n", string(body))
}
