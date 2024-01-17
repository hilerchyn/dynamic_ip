package cmd

import (
	"dynamic_ip/pkg"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alibabacloud-go/alidns-20150109/v4/client"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(querySubDomainsCmd)
}

var querySubDomainsCmd = &cobra.Command{
	Use:   "querySubDomains",
	Short: "Query sub-domains of specific domain",
	RunE:  querySubDomainsCmdFunc,
}

func querySubDomainsCmdFunc(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify domain name")
	}

	domainName := args[0]
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(domainName),
	}
	runtime := &util.RuntimeOptions{}
	domainRecords := &client.DescribeDomainRecordsResponseBodyDomainRecords{}

	tryErr := func(client *alidns20150109.Client) error {
		result, err := client.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
		if err != nil {
			return err
		}

		if *result.StatusCode != http.StatusOK {
			return fmt.Errorf("error with http status code: %d", *result.StatusCode)
		}

		domainRecords = result.Body.DomainRecords
		return nil
	}

	err := pkg.Invoke(tryErr)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"#",
		"Record Id",
		"RR",
		"Domain Name",
		"Status",
		"Type",
		"Value",
		"Update Time",
	})

	for idx, record := range domainRecords.Record {
		updateTime := ""
		if record.UpdateTimestamp != nil {
			updateTime = time.UnixMilli(*record.UpdateTimestamp).Format(time.RFC3339)
		}
		t.AppendRow(table.Row{
			idx,
			*record.RecordId,
			*record.RR,
			*record.DomainName,
			*record.Status,
			*record.Type,
			*record.Value,
			updateTime,
		})
		t.AppendSeparator()
	}

	t.Render()

	return nil
}
