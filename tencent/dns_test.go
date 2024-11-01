package tencent

import (
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	v20210323 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func TestCreateAliasDomain(t *testing.T) {
	credential := common.NewCredential("SecretId", "SecretKey")
	client, _ := v20210323.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())

	request := v20210323.NewCreateDomainAliasRequest()
	response, err := client.CreateDomainAlias(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response.ToJsonString())
}
