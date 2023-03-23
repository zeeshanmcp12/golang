package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Set up Workload Identity Federation credentials
	creds, err := google.WorkloadIdentityPoolCredentials(ctx, &google.WorkloadIdentityPoolCredentialsOptions{
		Audience:                       "AUDIENCE",
		SubjectTokenType:               "urn:ietf:params:oauth:token-type:jwt",
		TokenURL:                       "https://sts.googleapis.com/v1/token",
		ServiceAccountImpersonationURL: "https://iamcredentials.googleapis.com/v1/projects/-/serviceAccounts/SERVICE_ACCOUNT_EMAIL:generateAccessToken",
		CredentialSource: &google.CredentialSource{
			EnvironmentID:               "azure",
			URL:                         "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https://sts.windows.net/{tenant-id}",
			RegionalCredVerificationURL: "https://management.azure.com/subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.Compute/virtualMachines/{vm-name}/identity/token?api-version=2019-12-01&resource=https://sts.windows.net/{tenant-id}",
		},
	})
	// GET 'http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https://management.azure.com/' HTTP/1.1 Metadata: true
	if err != nil {
		fmt.Printf("Error creating credentials: %v\n", err)
		os.Exit(1)
	}

	// Impersonate the service account
	ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: "SERVICE_ACCOUNT_EMAIL",
		Scopes:          []string{"https://www.googleapis.com/auth/bigquery"},
	}, option.WithCredentials(creds))

	if err != nil {
		fmt.Printf("Error creating impersonated token source: %v\n", err)
		os.Exit(1)
	}

	// Use the impersonated credentials to access BigQuery, for example
	client, err := bigquery.NewClient(ctx, "YOUR_PROJECT_ID", option.WithTokenSource(ts))
	if err != nil {
		fmt.Printf("Error creating BigQuery client: %v\n", err)
		os.Exit(1)
	}

	// Use the client to interact with BigQuery (e.g., list datasets)
	it := client.Datasets(ctx)
	for {
		dataset, err := it.Next()
		if err == bigquery.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error listing datasets: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Dataset: %s\n", dataset.DatasetID)
	}
}
