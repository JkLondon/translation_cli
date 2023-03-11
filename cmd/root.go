package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "translate_cli",
	Short: "Translate via google api",
	Long:  `Translate via google api`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://translo.p.rapidapi.com/api/v3/translate"

		q := netUrl.Values{}
		q.Set("to", "ru")
		q.Set("text", args[0])
		payload := strings.NewReader(q.Encode())
		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPIKEY"))
		req.Header.Add("X-RapidAPI-Host", "translo.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		resp := TranslateResponse{}
		json.Unmarshal(body, &resp)
		fmt.Println(resp.TranslatedText)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
