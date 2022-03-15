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
	"time"

	"github.com/spf13/cobra"
)

// cryptop10Cmd represents the cryptop10 command
var cryptop10Cmd = &cobra.Command{
	Use:   "cryptop10",
	Short: "Get actual top 10 crypto.",
	Long:  "Get actual top 10 crypto.",
	Run: func(cmd *cobra.Command, args []string) {
		top10()
	},
}

type crypto []struct {
	MarketCapRank     int       `json:"market_cap_rank,omitempty"`
	ID                string    `json:"id,omitempty"`
	Symbol            string    `json:"symbol,omitempty"`
	Name              string    `json:"name,omitempty"`
	CurrentPrice      float32   `json:"current_price,omitempty"`
	DayHigh           float32   `json:"high_24h,omitempty"`
	DayLow            float32   `json:"low_24h,omitempty"`
	CirculatingSupply float32   `json:"circulating_supply,omitempty"`
	TotalSupply       float32   `json:"total_supply,omitempty"`
	Ath               float32   `json:"ath,omitempty"`
	AthDate           time.Time `json:"ath_date,omitempty"`
	Atl               float32   `json:"atl,omitempty"`
	AtlDate           time.Time `json:"atl_date,omitempty"`
	LastUpdate        time.Time `json:"last_updated,omitempty"`
}

func top10() {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"
	responseByte := helpers.GetData(url)
	crypto := &crypto{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		log.Println("Unable to unmarshall response.", err)
	}
	for i, item := range *crypto {
		fmt.Printf("\n=====================\t#%d %s\t====================== \n\n", i+1, strings.Title(item.ID))
		fmt.Printf("Symbol: %s\nCurrent price: %.2f\n24hHigh: %.2f\n24hLow: %.2f\nCirculating Supply: %.2f\nTotal Supply: %.2f\nATH: %.2f\nATHDate: %v\nATL: %.2f\nATLDate: %v\nLast Update: %v\n\n", strings.ToUpper(item.Symbol), item.CurrentPrice, item.DayHigh, item.DayLow, item.CirculatingSupply, item.TotalSupply, item.Ath, item.AthDate, item.Atl, item.AtlDate, item.LastUpdate)

	}
}

func init() {
	rootCmd.AddCommand(cryptop10Cmd)

}
