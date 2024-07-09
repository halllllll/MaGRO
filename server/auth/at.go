package auth

/*

👺 this is temporary implementation til asign entra app to user manager privilege 👹

*/

import (
	"context"
	"fmt"
	"os"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

/*

👺 this is temporary implementation til asign entra app to user manager privilege 👹

*/

func GetAccessTokenBySecret(ctx context.Context)(string,error){
	// configを通さず環境変数を直接読む
	s, ok := os.LookupEnv("CLIENT_SECRET")
	if !ok{
		return "", fmt.Errorf("secret not found")
	}
	t, ok := os.LookupEnv("TENANT_ID")
	if !ok{
		return "", fmt.Errorf("tenant id not found")
	}
	c, ok := os.LookupEnv("ENTRA_CLIENT_ID")
	if !ok{
		return "", fmt.Errorf("client id not found")
	}

	cred, err := confidential.NewCredFromSecret(s)
	if err != nil {
		return "", err
	}

	confidentialClient, err := confidential.New(fmt.Sprintf("https://login.microsoftonline.com/%s", t),c, cred)
	if err != nil {
		return "", err
	}

	scopes := []string{"https://graph.microsoft.com/.default"}
	result, err := confidentialClient.AcquireTokenSilent(ctx, scopes)
	if err != nil {
		result, err = confidentialClient.AcquireTokenByCredential(context.TODO(), scopes)
		if err != nil {
			return "", err
		}
	}

	return result.AccessToken, nil
}