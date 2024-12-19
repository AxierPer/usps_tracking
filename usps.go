package main

import (
	"fmt"
	tk "usps/token"
	track "usps/tracking"
)

func main() {
	token := tk.GenerarToken()
	fmt.Println(token["access_token"])
	tokenString := token["access_token"].(string)

	info := track.TrackingPackage("Tracking Number", tokenString)
	fmt.Println(info)
}
