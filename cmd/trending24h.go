/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"go-cli-crypto/helpers"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// trending24hCmd represents the trending24h command
var trending24hCmd = &cobra.Command{
	Use:   "trending24h",
	Short: "Get trending24h crypto data.",
	Long:  "Get trending24h crypto data.",
	Run: func(cmd *cobra.Command, args []string) {
		trending24h()
	},
}

type trendingStruct struct {
	Coins []struct {
		Item struct {
			ID            string  `json:"id"`
			CoinID        int     `json:"coin_id"`
			Name          string  `json:"name"`
			Symbol        string  `json:"symbol"`
			MarketCapRank int     `json:"market_cap_rank"`
			PriceBtc      float64 `json:"price_btc"`
			Score         int     `json:"score"`
		} `json:"item"`
	} `json:"coins"`
}

func trending24h() {
	url := "https://api.coingecko.com/api/v3/search/trending"
	responseByte := helpers.GetData(url)
	crypto := &trendingStruct{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		log.Println("Unable to unmarshall response.", err)
	}
	fmt.Println(crypto)
	for _, item := range crypto.Coins {
		fmt.Printf("\n=====================\t #%d %s \t====================== \n\n", item.Item.Score+1, strings.Title(item.Item.Name))
		fmt.Printf("Symbol: %s\nName: %s\nMarket Cap Rank: %d\nPrice BTC: %.2f\n\n", strings.ToUpper(item.Item.Symbol), item.Item.Name, item.Item.MarketCapRank, item.Item.PriceBtc)
	}
}
func init() {
	rootCmd.AddCommand(trending24hCmd)

}
