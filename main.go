package main

import (
	"fmt"
	"github.com/keybase/go-keychain"
	"os"
	"os/user"
)

func main() {
	u, _ := user.Current()

	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(os.Args[1])
	query.SetAccount(u.Username)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		// Error
	} else if len(results) != 1 {
		// Not found
	} else {
		password := string(results[0].Data)
		fmt.Printf("%s\n", password)
	}
}
