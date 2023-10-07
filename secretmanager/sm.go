package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/jsonandrx/twitterGo/awsgo"
	"github.com/jsonandrx/twitterGo/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret
	fmt.Println("> Get Secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &dataSecret)
	fmt.Println("-> Lectura secret Ok" + secretName)
	return dataSecret, nil
}
