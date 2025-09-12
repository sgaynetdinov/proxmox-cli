package proxmox

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/joho/godotenv"
)

type ProxmoxClient = pveSDK.Client

func createClient() (*ProxmoxClient, error) {
	apiURL := os.Getenv("PM_API_URL")
	if apiURL == "" {
		return nil, fmt.Errorf("PM_API_URL environment variable must be set")
	}

	return pveSDK.NewClient(apiURL, nil, "", &tls.Config{InsecureSkipVerify: true}, "", 30)
}

func loginClient(client *ProxmoxClient, ctx context.Context) error {
	user := os.Getenv("PM_USER")
	pass := os.Getenv("PM_PASS")

	if user == "" || pass == "" {
		return fmt.Errorf("PM_USER and PM_PASS environment variables must be set")
	}

	return client.Login(ctx, user, pass, "")
}

func Login(ctx context.Context) *ProxmoxClient {
	godotenv.Load()

	client, err := createClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}

	err = loginClient(client, ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error logging in: %v\n", err)
		os.Exit(1)
	}

	return client
}
