

// This file is auto-generated, don't edit it. Thanks.
package main

import (
  "encoding/json"
  "strings"
  "fmt"
  "os"
  alidns20150109  "github.com/alibabacloud-go/alidns-20150109/v4/client"
  openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  "github.com/alibabacloud-go/tea/tea"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
  config := &openapi.Config{
    // 必填，您的 AccessKey ID
    AccessKeyId: accessKeyId,
    // 必填，您的 AccessKey Secret
    AccessKeySecret: accessKeySecret,
  }
  // Endpoint 请参考 https://api.aliyun.com/product/Alidns
  config.Endpoint = tea.String("alidns.cn-shanghai.aliyuncs.com")
  _result = &alidns20150109.Client{}
  _result, _err = alidns20150109.NewClient(config)
  return _result, _err
}

func _main (args []*string) (_err error) {
  // 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
  // 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
  client, _err := CreateClient(tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")))
  if _err != nil {
    return _err
  }

  describeDomainsRequest := &alidns20150109.DescribeDomainsRequest{}
  runtime := &util.RuntimeOptions{}
  tryErr := func()(_e error) {
    defer func() {
      if r := tea.Recover(recover()); r != nil {
        _e = r
      }
    }()
    // 复制代码运行请自行打印 API 的返回值
    result, _err := client.DescribeDomainsWithOptions(describeDomainsRequest, runtime)
    if _err != nil {
      return _err
    }

    println(result.String())

    return nil
  }()

  if tryErr != nil {
    var error = &tea.SDKError{}
    if _t, ok := tryErr.(*tea.SDKError); ok {
      error = _t
    } else {
      error.Message = tea.String(tryErr.Error())
    }
    // 错误 message
    fmt.Println(tea.StringValue(error.Message))
    // 诊断地址
    var data interface{}
    d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
    d.Decode(&data)
    if m, ok := data.(map[string]interface{}); ok {
      recommend, _ := m["Recommend"]
      fmt.Println(recommend)
    }
    _, _err = util.AssertAsString(error.Message)
    if _err != nil {
      return _err
    }
  }
  return _err
}


//func main() {
//  err := _main(tea.StringSlice(os.Args[1:]))
//  if err != nil {
//    panic(err)
//  }
//}
