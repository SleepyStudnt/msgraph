package main

import (
	"fmt"

	"github.com/open-networks/go-msgraph"
	"github.com/spf13/viper"
)

func config() *msgraph.GraphClient {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	graphClient, err := msgraph.NewGraphClient(viper.GetString("TENANTID"), viper.GetString("CLIENTID"), viper.GetString("SECRET"))
	if err != nil {
		fmt.Println("Credentials are invalid: ", err)
	}
	return graphClient
}

func main() {
	client := config()
	groups, err := client.ListGroups()
	if err != nil {
		fmt.Printf("Error: " + err.Error())
	} else {
		for _, g := range groups {
			fmt.Printf("\nGroup: %s\n===================\n\n", g.DisplayName)
			members, _ := g.ListMembers()
			for _, u := range members {
				fmt.Printf("User: %s\n", u.DisplayName)
			}
		}
	}
}
