package cwapi

import (
	"encoding/base64"
	"fmt"
)

// Set the name so it appears at the top of the file list.

func ConnectwiseConfiguration(cwClientId, cwCompanyName, cwPublicKey, cwPrivateKey string) *Configuration {
	config := NewConfiguration()
	config.AddDefaultHeader("User-Agent", "Nexigen-Connectwise-Client")
	config.AddDefaultHeader("Accept", "application/vnd.connectwise.com+json")
	config.AddDefaultHeader("Content-Type", "application/json")
	// This removes the need to set the clientId in each request
	config.AddDefaultHeader("clientId", cwClientId)
	// Create the authentication string: company+publicKey:privateKey
	auth := fmt.Sprintf("%s+%s:%s", cwCompanyName, cwPublicKey, cwPrivateKey)
	// Encode to base64
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	config.AddDefaultHeader("Authorization", "Basic "+encodedAuth)
	return config
}
