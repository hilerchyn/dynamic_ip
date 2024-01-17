package cmd

import (
	"dynamic_ip/pkg"
	"fmt"
	"net/http"
	"os"

	"github.com/alibabacloud-go/alidns-20150109/v4/client"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(queryDomainsCmd)
}

var queryDomainsCmd = &cobra.Command{
	Use:   "queryDomains",
	Short: "Query domains from Aliyun",
	RunE:  queryDomainsCmdFunc,
}

func queryDomainsCmdFunc(cmd *cobra.Command, args []string) error {
	describeDomainsRequest := &alidns20150109.DescribeDomainsRequest{}
	runtime := &util.RuntimeOptions{}
	domains := &client.DescribeDomainsResponseBodyDomains{}
	tryErr := func(client *alidns20150109.Client) error {
		// 复制代码运行请自行打印 API 的返回值
		result, _err := client.DescribeDomainsWithOptions(describeDomainsRequest, runtime)
		if _err != nil {
			return _err
		}

		if *result.StatusCode != http.StatusOK {
			return fmt.Errorf("error with http status code: %d", *result.StatusCode)
		}

		domains = result.Body.Domains
		return nil
	}

	err := pkg.Invoke(tryErr)
	if err != nil {
		return err
	}

	// print domains
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Domain Id", "Domain Name", "Record Count", "Create Time"})

	for idx, domain := range domains.Domain {
		t.AppendRow(table.Row{idx, *domain.DomainId, *domain.DomainName, *domain.RecordCount, *domain.CreateTime})
		t.AppendSeparator()
	}

	t.Render()

	return nil
}
