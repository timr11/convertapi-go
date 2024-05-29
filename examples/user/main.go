package main

import (
	"fmt"
	"github.com/timr11/convertapi-go"
	"github.com/timr11/convertapi-go/config"
	"os"
)

func main() {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	if user, err := convertapi.UserInfo(nil); err == nil {
		fmt.Println("User information: ")
		fmt.Printf("%+v\n", user)
	} else {
		fmt.Println(err)
	}
}
