package openchain

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResponse_ParseAccount(t *testing.T) {
	var respJson = `{
    "code":"200",
  "data":"{\"recovery_key\":\"eb716ed3d6755bb9a6cc9b44a5d5c18194ce5243648eb858953504d27e4bb07eb71ee4ff66410f32ff2bb704c4b6e94c8afabce21e36cbf5e59975d87457f0ae\",\"balance\":0,\"auth_map\":[{\"value\":100,\"key\":\"eb716ed3d6755bb9a6cc9b44a5d5c18194ce5243648eb858953504d27e4bb07eb71ee4ff66410f32ff2bb704c4b6e94c8afabce21e36cbf5e59975d87457f0ae\"}],\"encryption_key\":\"\",\"id\":\"3360513e31c13c319a708376f6b3a6dc68f0b655c9cf9545c811c1da63f602cd\",\"recovery_time\":0,\"version\":2,\"status\":0}",
  "success":true
}`

	var resp Response
	err := json.Unmarshal([]byte(respJson), &resp)
	require.Nil(t, err)
	fmt.Println(resp)

	acc, err := resp.ParseAccount()
	require.Nil(t, err)
	fmt.Println(acc)
	fmt.Println("my auth map:", acc.AuthMap)
}
