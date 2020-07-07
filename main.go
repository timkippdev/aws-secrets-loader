package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type secret struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

var region *string

func main() {
	filePath := flag.String("file", "data/data.json", "file to read JSON data from")
	delete := flag.Bool("delete", false, "delete all secrets contained in the JSON data file")
	region = flag.String("region", "us-west-2", "AWS region")
	flag.Parse()
	file, error := ioutil.ReadFile(*filePath)
	if error != nil {
		log.Fatal(error)
	}

	var secrets []secret
	err := json.Unmarshal([]byte(file), &secrets)
	if err != nil {
		panic(err)
	}

	sm := getSMClient()

	for i := 0; i < len(secrets); i++ {
		secret := secrets[i]
		if *delete {
			deleteSecret(sm, secret)
		} else {
			createSecret(sm, secret)
		}
	}
}

// TODO: convert to upsertSecret
func createSecret(sm *secretsmanager.SecretsManager, secret secret) {
	jsonAsBytes, err := json.Marshal(secret.Value)
	if err != nil {
		panic(err)
	}

	jsonAsString := string(jsonAsBytes)

	log.Printf("Creating secret with name: %s", secret.Name)

	output, err := sm.CreateSecret(&secretsmanager.CreateSecretInput{
		Name:         &secret.Name,
		SecretString: &jsonAsString,
	})

	if err != nil {
		// TODO: update to skip if secret already exists / upsert
		log.Panic(err)
	}

	log.Printf("Secret successfully created with name: %s | region: %s", *output.Name, *region)
}

func deleteSecret(sm *secretsmanager.SecretsManager, secret secret) {
	log.Printf("Deleting secret with name: %s", secret.Name)

	output, err := sm.DeleteSecret(&secretsmanager.DeleteSecretInput{
		SecretId:                   &secret.Name,
		ForceDeleteWithoutRecovery: aws.Bool(true),
	})

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Secret successfully deleted with name: %s | region: %s", *output.Name, *region)
}

func getSMClient() *secretsmanager.SecretsManager {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(*region),
	})

	if err != nil {
		log.Panic(err)
	}

	return secretsmanager.New(session)
}
