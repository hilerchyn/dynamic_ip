package cmd

import (
	"dynamic_ip/pkg"
	"fmt"
	"net/http"

	"github.com/alibabacloud-go/alidns-20150109/v4/client"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newUpdateDomainRecordCmd())
}

var recordId, recordRr, recordType, recordValue string

func newUpdateDomainRecordCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "updateDomainRecord",
		Short: "Update domain record",
		RunE:  updateDomainRecordCmdFunc,
	}

	cmd.Flags().StringVarP(&recordId, "recordId", "i", "", "")
	cmd.Flags().StringVarP(&recordRr, "rr", "r", "", "")
	cmd.Flags().StringVarP(&recordType, "recordType", "t", "A", "A")
	cmd.Flags().StringVarP(&recordValue, "ip", "v", "", "127.0.0.1")

	return cmd
}

func updateDomainRecordCmdFunc(cmd *cobra.Command, args []string) error {
	if recordId == "" || recordRr == "" || recordType == "" || recordValue == "" {
		return fmt.Errorf("please specify required arguments")
	}

	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(recordId),
		RR:       tea.String(recordRr),
		Type:     tea.String(recordType),
		Value:    tea.String(recordValue),
	}
	runtime := &util.RuntimeOptions{}
	updateResponseBody := &client.UpdateDomainRecordResponseBody{}

	tryErr := func(client *alidns20150109.Client) error {
		result, err := client.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
		if err != nil {
			return err
		}

		if *result.StatusCode != http.StatusOK {
			return fmt.Errorf("error with http status code: %d", *result.StatusCode)
		}

		updateResponseBody = result.Body
		return nil
	}

	err := pkg.Invoke(tryErr)
	if err != nil {
		return err
	}

	println(updateResponseBody.String())

	return nil
}
