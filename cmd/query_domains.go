package cmd

import (
	"dynamic_ip/pkg"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(queryDomainsCmd)
}

var queryDomainsCmd = &cobra.Command{
    Use: "queryDomains",
    Short: "Query domains from Aliyun",
    RunE: queryDomainsCmdFunc,
}

func queryDomainsCmdFunc(cmd *cobra.Command, args []string) error {
  println("query domains")
  describeDomainsRequest := &alidns20150109.DescribeDomainsRequest{}
  runtime := &util.RuntimeOptions{}
  tryErr := func(client *alidns20150109.Client)(_e error) {
    // 复制代码运行请自行打印 API 的返回值
    result, _err := client.DescribeDomainsWithOptions(describeDomainsRequest, runtime)
    if _err != nil {
      return _err
    }

    println(result.String())

    return nil
  }

  err := pkg.Invoke(tryErr)

  return err
}
