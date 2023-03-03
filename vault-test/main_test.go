package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

type (
	VaultEntries struct {
		REPO_NAME       string `json:"REPO_NAME"`
		VAULT_ENV       string `json:"VAULT_ENV"`
		VAULT_TESTING   string `json:"VAULT_TESTING"`
		TEST_USER       string `json:"TEST_USER"`
		TEST_PASSWORD   string `json:"TEST_PASSWORD"`
		TEST_REQUEST_id string `json:"TEST_REQUEST_id"`
	}
)

var (
	fixturesVault VaultEntries
)

func prepareVaultKVService(c *api.Client, t *testing.T) {

	fmt.Println("loading fixtures file")
	fixturesFile, err := os.Open("./vault_fixture.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Fixture File")
	defer fixturesFile.Close()

	byteValue, _ := io.ReadAll(fixturesFile)
	json.Unmarshal(byteValue, &fixturesVault) //nolint:errcheck // Parse fixtures

	kvData := map[string]interface{}{
		"data": populateEntries(fixturesVault),
	}

	secretRaw, err := c.Logical().Write("kv/data/"+fixturesVault.REPO_NAME+"/"+fixturesVault.VAULT_ENV, kvData)

	if err != nil {
		t.Fatalf("write failed, err :%v, resp: %#v", err, secretRaw)
	}

	data, err := c.Logical().Read("kv/data/" + fixturesVault.REPO_NAME + "/" + fixturesVault.VAULT_ENV)
	if err != nil {
		t.Fatal(err)
	}

	//b, _ := json.Marshal(secretRaw)
	//fmt.Println("secretRaw: ", string(b))
	//b, _ = json.Marshal(data.Data)
	//fmt.Println("secretRaw: ", string(b))
	//
	//// Convert struct
	//var mapData map[string]interface{}
	//if err := json.Unmarshal(b, &mapData); err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println("setupVaultKVService: ", data.Data["data"])

}
func prepareVaultAppRoleService(c *api.Client, t *testing.T) {

	fmt.Println("loading fixtures file")
	fixturesFile, err := os.Open("./vault_fixture.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Fixture File")
	defer fixturesFile.Close()

	byteValue, _ := io.ReadAll(fixturesFile)
	json.Unmarshal(byteValue, &fixturesVault) //nolint:errcheck // Parse fixtures

	valueInjector(c, t)

}
func populateEntries(entries VaultEntries) map[string]interface{} {
	v := reflect.ValueOf(entries)
	typeOfS := v.Type()
	var body = ""
	for i := 0; i < v.NumField(); i++ {
		if i == v.NumField()-1 {
			body += `"` + typeOfS.Field(i).Name + `": "` + v.Field(i).Interface().(string) + `"`
		} else {
			body += `"` + typeOfS.Field(i).Name + `": "` + v.Field(i).Interface().(string) + `",`
		}
	}
	var data = []byte(`{` + body + `}`)
	m := make(map[string]interface{})

	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
	}
	return m

}

func valueInjector(c *api.Client, t *testing.T) {

	err := putSecret(c, populateEntries(fixturesVault), fixturesVault.REPO_NAME+"/"+fixturesVault.VAULT_ENV)
	if err != nil {
		t.Fatal(err)
	}

	data, err := c.Logical().Read("secret/data/" + fixturesVault.REPO_NAME + "/" + fixturesVault.VAULT_ENV)
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(data)
	log.Printf("After valueInjector: " + string(b))

}

func TestPutSecret(t *testing.T) {
	t.Helper()
	// Vault initialization
	//server := createMockedAppRoleTestVault(t)
	server := createMockedKVTestVault(t)

	client := server.Client
	prepareVaultKVService(client, t)
	defer server.Cluster.Cleanup()

	response := GetFromVaultKV(true, "TEST_USER", "testuser", client)
	log.Printf("GetFromVault: TEST_USER: " + response)
	assert.Contains(t, response, fixturesVault.TEST_USER, "Value matched")

}

func TestMyTest(t *testing.T) {
	t.Helper()
	// Vault initialization
	server := createMockedAppRoleTestVault(t)
	client := server.Client
	defer server.Cluster.Cleanup()
	//Populate entries
	prepareVaultAppRoleService(client, t)

	response := GetFromVault(true, "TEST_USER", "testuser", client)
	log.Printf("GetFromVault: TEST_USER: " + response)
	assert.Contains(t, response, fixturesVault.TEST_USER, "Value matched")

}
func GetFromVault(flag bool, key, value string, c *api.Client) string {
	if flag {
		return Retrieve(key, fixturesVault.REPO_NAME, fixturesVault.VAULT_ENV, c)
	}

	return value
}
func GetFromVaultKV(flag bool, key, value string, c *api.Client) string {
	if flag {
		return RetrieveKV(key, fixturesVault.REPO_NAME, fixturesVault.VAULT_ENV, c)
	}

	return value
}

func Retrieve(key, repoName, vaultEnv string, client *api.Client) string {
	log.Printf("Retrieve: " + "secret/data/" + repoName + "/" + vaultEnv + " | " + key)
	secret, err := client.Logical().Read("secret/data/" + repoName + "/" + vaultEnv)

	//token expired related errors.
	if err != nil {
		e := &FetchError{key, err}
		log.Println(e.Error())
		os.Exit(1)
	}

	//token is valid but secret is not present in vault
	if secret == nil {
		e := &FetchError{repoName + "/" + vaultEnv, errors.New("Entry not found in vault")}
		log.Println(e.Error())
		os.Exit(1)
	}

	//No k-v data in vault for the repo
	m, ok := secret.Data[""+key+""].(string)
	if !ok {
		e := &FetchError{key, err}
		log.Println(e.Error())
		os.Exit(1)
	}

	return string(m)
}
func RetrieveKV(key, repoName, vaultEnv string, client *api.Client) string {
	log.Printf("Retrieve: " + "kv/data/" + repoName + "/" + vaultEnv + " | " + key)
	secret, err := client.Logical().Read("kv/data/" + repoName + "/" + vaultEnv)

	//token expired related errors.
	if err != nil {
		e := &FetchError{key, err}
		log.Println(e.Error())
		os.Exit(1)
	}

	//token is valid but secret is not present in vault
	if secret == nil {
		e := &FetchError{repoName + "/" + vaultEnv, errors.New("Entry not found in vault")}
		log.Println(e.Error())
		os.Exit(1)
	}

	//No k-v data in vault for the repo
	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		e := &FetchError{key, err}
		log.Println(e.Error())
		os.Exit(1)
	}

	//Expect the key-value to be only string and not string-Map pairs
	//Need to add support for a nested json as value field.
	return m[key].(string)
}

type FetchError struct {
	Vaultkey string
	Err      error
}

func (e *FetchError) Error() string {
	return "Fetch for key or repo : " + e.Vaultkey + " resulted in " + e.Err.Error()
}
