package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// url := "https://onesignal.com/api/v1/notifications"
		// fmt.Println("URL:>", url)
		// viper.AutomaticEnv()
		// appID := viper.GetString("OSAPPID")
		// APIKEY := viper.GetString("OSAPIKEY")
		// var jsonStr = []byte(`{
		// 	"app_id": "` + appID + `",
		// 	"include_player_ids": ["3ca9a849-bdf7-4986-afea-89bf18c94b6b"],
		// 	"data": {"foo": "bar"},
		// 	"contents": {"en": "kuyyyyy2323"}
		//   }`)
		// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		// req.Header.Set("Authorization", "Basic "+APIKEY)
		// req.Header.Set("Content-Type", "application/json; charset=utf-8")

		// client := &http.Client{}
		// resp, err := client.Do(req)
		// if err != nil {
		// 	panic(err)
		// }
		// defer resp.Body.Close()

		// fmt.Println("response Status:", resp.Status)
		// fmt.Println("response Headers:", resp.Header)
		// body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println("response Body:", string(body))
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	port := viper.GetString("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func genPlayerIDS(arry []string) string {
	mySTR := ``
	for i, str := range arry {
		mySTR = mySTR + `'` + str + `'`
		if i < len(arry)-1 {
			mySTR = mySTR + `,`
		}
	}
	return mySTR
}
