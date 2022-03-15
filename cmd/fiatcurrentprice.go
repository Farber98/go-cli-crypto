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

// fiatcurrentpriceCmd represents the fiatcurrentprice command
var fiatcurrentpriceCmd = &cobra.Command{
	Use:   "fiatcurrentprice",
	Short: "Get actual top 10 crypto.",
	Long:  "Get actual top 10 crypto.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one crypto id needed.")
		} else {
			fiatCurrentPrice(args[0])
		}
	},
}

type fiatCurrentPriceStruct struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	MarketData struct {
		CurrentPrice struct {
			Aud  int `json:"aud"`
			Cad  int `json:"cad"`
			Chf  int `json:"chf"`
			Cny  int `json:"cny"`
			Eur  int `json:"eur"`
			Gbp  int `json:"gbp"`
			Hkd  int `json:"hkd"`
			Jpy  int `json:"jpy"`
			Nzd  int `json:"nzd"`
			Usd  int `json:"usd"`
			Sats int `json:"sats"`
		} `json:"current_price"`
	} `json:"market_data"`
}

func fiatCurrentPrice(id string) {
	url := "https://api.coingecko.com/api/v3/coins/" + id + "?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false"
	responseByte := helpers.GetData(url)
	crypto := &fiatCurrentPriceStruct{}
	err := json.Unmarshal(responseByte, &crypto)
	if err != nil {
		log.Println("Unable to unmarshall response.", err)
	}
	fmt.Printf("\n=====================\t%s\t====================== \n\n", strings.Title(crypto.ID))
	fmt.Printf("Symbol: %s\nName: %s\nCurrent Price: %+v\n\n", strings.ToUpper(crypto.Symbol), crypto.Name, crypto.MarketData.CurrentPrice)
}

func init() {
	rootCmd.AddCommand(fiatcurrentpriceCmd)

}
