package main

import (
	kv "github.com/hashicorp/vault-plugin-secrets-kv"
	"log"
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/builtin/credential/approle"
	vaulthttp "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/vault"
)

const (
	// TestVaultToken is the Vault token used for tests
	testVaultToken = "unittesttoken"
)

type vaultTest struct {
	Cluster       *vault.TestCluster
	Client        *api.Client
	AppRoleID     string
	AppRoleSecret string
}

// GetMockedVaultServer Creates the test server
func createMockedKVTestVault(t *testing.T) vaultTest {
	t.Helper()

	// CoreConfig parameterizes the Vault core config
	coreConfig := &vault.CoreConfig{
		LogicalBackends: map[string]logical.Factory{
			"kv": kv.VersionedKVFactory,
		},
	}
	cluster := vault.NewTestCluster(t, coreConfig, &vault.TestClusterOptions{
		HandlerFunc: vaulthttp.Handler,
	})
	cluster.Start()

	core := cluster.Cores[0].Core
	client := cluster.Cores[0].Client
	vault.TestWaitActive(t, core)

	// Mount a KVv2 backend
	err := client.Sys().Mount("kv", &api.MountInput{
		Type: "kv-v2",
	})
	if err != nil {
		t.Fatal(err)
	}
	//kvData := map[string]interface{}{
	//	"data": populateEntries(fixturesVault),
	//}
	//
	//secretRaw, err := client.Logical().Write("kv/data/vaultrepo", kvData)
	//
	//if err != nil {
	//	t.Fatalf("write failed, err :%v, resp: %#v", err, secretRaw)
	//}
	return vaultTest{
		Cluster:       cluster,
		Client:        client,
		AppRoleID:     "roleID",
		AppRoleSecret: "secretID",
	}
}
func createMockedAppRoleTestVault(t *testing.T) vaultTest {
	t.Helper()

	cluster := vault.NewTestCluster(t, &vault.CoreConfig{
		DevToken: testVaultToken,
		CredentialBackends: map[string]logical.Factory{
			"approle": approle.Factory,
		},
	}, &vault.TestClusterOptions{
		HandlerFunc: vaulthttp.Handler,
	})
	cluster.Start()

	core := cluster.Cores[0].Core
	vault.TestWaitActive(t, core)
	client := cluster.Cores[0].Client

	// Enable approle
	err := client.Sys().EnableAuthWithOptions("approle", &api.EnableAuthOptions{
		Type: "approle",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create an approle
	_, err = client.Logical().Write("auth/approle/role/unittest", map[string]interface{}{
		"policies": []string{"unittest"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Gets the role ID, that is basically the 'username' used to log into vault
	res, err := client.Logical().Read("auth/approle/role/unittest/role-id")
	if err != nil {
		t.Fatal(err)
	}

	// Keep the roleID for later use
	roleID, ok := res.Data["role_id"].(string)
	if !ok {
		t.Fatal("Could not read the approle")
	}
	log.Printf("roleID: " + roleID)

	// Create a secretID that is basically the password for the approle
	res, err = client.Logical().Write("auth/approle/role/unittest/secret-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Use thre secretID later
	secretID, ok := res.Data["secret_id"].(string)
	if !ok {
		t.Fatal("Could not generate the secret id")
	}
	log.Printf("secretID: " + secretID)

	// Create a broad policy to allow the approle to do whatever
	err = client.Sys().PutPolicy("unittest", `
        path "*" {
            capabilities = ["create", "read", "list", "update", "delete"]
        }
    `)
	if err != nil {
		t.Fatal(err)
	}
	return vaultTest{
		Cluster:       cluster,
		Client:        client,
		AppRoleID:     roleID,
		AppRoleSecret: secretID,
	}
}
