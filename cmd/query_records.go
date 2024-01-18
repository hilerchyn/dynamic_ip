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
	rootCmd.AddCommand(newQueryRecordsCmd())
}

var recordName string

func newQueryRecordsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queryRecords",
		Short: "Query records of specific domain",
		RunE:  queryRecordsCmdFunc,
	}

	cmd.Flags().StringVarP(&recordName, "recordName", "r", "", "")

	return cmd
}

func queryRecordsCmdFunc(cmd *cobra.Command, args []string) error {
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

	if recordName != "" {
		ip := ""
		for _, record := range domainRecords.Record {
			if *record.RR != recordName {
				continue
			}
			ip = *record.Value
			break
		}

		if ip != "" {
			fmt.Println(ip)
			return nil
		}
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
