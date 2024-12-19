package tracking

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func TrackingPackage(trackingNumber string, token string) (track map[string]interface{}) {
	url := fmt.Sprintf("https://api.usps.com/tracking/v3/tracking/%s", trackingNumber)
	bearerToken := fmt.Sprintf("Bearer %s", token)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Autorization", bearerToken)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	client := &http.Client{}
	var tracking map[string]interface{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error al enviar la solicitud HTTP: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error en la requesta HTTP: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error al leer el cuerpo de la requesta: ", err)
	}

	err = json.Unmarshal([]byte(body), &tracking)
	if err != nil {
		log.Fatal("Error al deserializar JSON: ", err)
	}

	return tracking
}
