package security

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/option"
)

// GCPConfig specifies all the gcp
// variables to retrieve a github token.
type GCPConfig struct {
	ProjectID string
	Storage   struct {
		BucketName string
		ObjectName string
	}
	KMS struct {
		KeyringID   string
		CryptoKeyID string
	}
}

// GithubToken returns github token that can be
// used for vanity imports.
func GithubToken(cfg *GCPConfig) (string, error) {
	ctx := context.Background()
	// func GithubToken(projectID, bucketName, objectName, keyringID, cryptoKeyID string) (string, error) {
	c, err := storage.NewClient(ctx, option.WithScopes(storage.ScopeReadOnly))
	if err != nil {
		return "", err
	}
	rc, err := c.Bucket(cfg.Storage.BucketName).Object(cfg.Storage.ObjectName).NewReader(ctx)
	if err != nil {
		return "", err
	}
	defer rc.Close()
	bts, err := ioutil.ReadAll(rc)
	if err != nil {
		return "", err
	}
	bts, err = decrypt(cfg.ProjectID, "global", cfg.KMS.KeyringID, cfg.KMS.CryptoKeyID, bts)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bts)), nil
}

func decrypt(projectID, locationID, keyRingID, cryptoKeyID string, ciphertext []byte) ([]byte, error) {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	cloudkmsService, err := cloudkms.New(client)
	if err != nil {
		return nil, err
	}
	parentName := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		projectID, locationID, keyRingID, cryptoKeyID)

	req := &cloudkms.DecryptRequest{Ciphertext: base64.StdEncoding.EncodeToString(ciphertext)}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Decrypt(parentName, req).Do()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(resp.Plaintext)
}
