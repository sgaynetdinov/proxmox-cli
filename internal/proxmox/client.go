package proxmox

import (
	"context"
	"crypto/tls"
	"fmt"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

type ProxmoxClient = pveSDK.Client

func createClient(apiURL string) (*ProxmoxClient, error) {
	return pveSDK.NewClient(apiURL, nil, "", &tls.Config{InsecureSkipVerify: true}, "", 30)
}

func NewClient(ctx context.Context, apiURL, username, password string) (*ProxmoxClient, error) {
	client, err := createClient(apiURL)
	if err != nil {
		return nil, fmt.Errorf("creating client: %w", err)
	}

	if err := client.Login(ctx, username, password, ""); err != nil {
		return nil, fmt.Errorf("logging in: %w", err)
	}

	return client, nil
}
