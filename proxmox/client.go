package proxmox

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

type ProxmoxClient = pveSDK.Client

func createClient(apiURL string) (*ProxmoxClient, error) {
	return pveSDK.NewClient(apiURL, nil, "", &tls.Config{InsecureSkipVerify: true}, "", 30)
}

func NewClient(ctx context.Context, apiURL, username, password string) *ProxmoxClient {
	client, err := createClient(apiURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}

	err = client.Login(ctx, username, password, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error logging in: %v\n", err)
		os.Exit(1)
	}

	return client
}
