package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/nero-15/calcio-app/config"
)

func main() {
	// 取得したいデータのURL作成
	url, _ := url.Parse("https://api.football-data.org/v2/")
	url.Path = path.Join(url.Path, "teams", "108")

	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("X-Auth-Token", config.Config.FootballDataApiToken) // アカウント登録時に送られてきたAPIトークンをリクエストヘッダーに追加
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
