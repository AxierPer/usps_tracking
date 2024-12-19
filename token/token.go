package token

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	cf "usps/config"
)

type usps struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

func GenerarToken() (tokenGen map[string]interface{}) {
	url := "https://api.usps.com/oauth2/v3/token"
	config := cf.LoadConfig()

	client_id := config["CLIENT_ID"]
	client_secret := config["CLIENT_SECRET"]

	fields := usps{
		GrantType:    "client_credentials",
		ClientId:     client_id,
		ClientSecret: client_secret,
		Scope:        "tracking",
	}

	payload, err := json.Marshal(fields)

	if err != nil {
		log.Fatal("Error en el payload: ", err)
	}

	var token map[string]interface{}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("Error en la solicitud HTTP: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error en la respuesta HTTP: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error al leer el cuerpo de la respuesta: ", err)
	}

	err = json.Unmarshal([]byte(body), &token)
	if err != nil {
		log.Fatal("Error al deserializar JSON: ", err)
	}

	return token
}
