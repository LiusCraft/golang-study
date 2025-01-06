package tencent

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	v20210323 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"liuscraft.op/study/util"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println(err)
	}
}

func TestCreateAliasDomain(t *testing.T) {
	secretId := os.Getenv("SecretId")
	secretKey := os.Getenv("SecretKey")
	secretKey, err := util.Base64DesDecrypt(secretKey, "")
	if err != nil {
		t.Error("Base64DesDecrypt failure", err)
		return
	}
	credential := common.NewCredential(secretId, secretKey)
	cpf := profile.NewClientProfile()
	client, _ := v20210323.NewClient(credential, regions.Guangzhou, cpf)

	request := v20210323.NewCreateDomainAliasRequest()
	request.Domain = common.StringPtr("")
	request.DomainAlias = common.StringPtr("")
	response, err := client.CreateDomainAlias(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	t.Logf("%s\n", response.ToJsonString())
}
