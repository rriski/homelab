package main

import (
	"fmt"
	"log"
	"os"
    "time"
    "math/rand"

	vault "github.com/hashicorp/vault/api"
	"github.com/sethvargo/go-password/password"
	"gopkg.in/yaml.v2"
)

type RandomPassword struct {
	Path string
	Data []struct {
		Key     string
		Length  int
		Special bool
	}
}

func main() {
    rand.Seed(time.Now().UnixNano())

	data, err := os.ReadFile("./config.yaml")

	if err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}

	randomPasswords := []RandomPassword{}

	err = yaml.Unmarshal([]byte(data), &randomPasswords)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	config := vault.DefaultConfig()

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	for _, randomPassword := range randomPasswords {
		path := fmt.Sprintf("/secret/data/%s", randomPassword.Path)

		secret, _ := client.Logical().Read(path)

		if secret == nil {
			secretData := map[string]interface{}{
				"data": map[string]interface{}{},
			}

			for _, randomKey := range randomPassword.Data {
                numDigits := rand.Intn(randomKey.Length / 2)
                numSpecial := 0
                if randomKey.Special {
                    numSpecial = rand.Intn(randomKey.Length / 2)
                }
                // First argument is length, second is the number of digits, third is
                // the number of special symbols, fourth is allowing uppercase and lowercase
                // characters and fifth is allowing repeat characters
				res, err := password.Generate(randomKey.Length, numDigits, numSpecial, false, true)
				if err != nil {
					log.Fatal(err)
				}

				secretData["data"].(map[string]interface{})[randomKey.Key] = res
			}

			_, err = client.Logical().Write(path, secretData)

			if err != nil {
				log.Fatalf("Unable to write secret: %v", err)
			} else {
				log.Println("Secret written successfully.")
			}
		} else {
			log.Println("Key at path %s already exists.", path)
		}
	}
}
