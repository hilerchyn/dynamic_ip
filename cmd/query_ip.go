package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var ipQueryApi string
var apiMap map[string]func(*cobra.Command, []string) error

func init() {
	apiMap = map[string]func(*cobra.Command, []string) error{}
	apiMap["icanhazip"] = queryIpViaIcanhazip
	apiMap["chinaz"] = queryIpViaChinaz

	rootCmd.AddCommand(newQueryIpCmd())
}

func newQueryIpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queryIp",
		Short: "Query dynamic ip",
		RunE:  queryIpCmdFunc,
	}

	cmd.Flags().StringVarP(&ipQueryApi, "ipQueryApi", "q", "icanhazip", "support value: icanhazip, chinaz")

	return cmd
}

func queryIpCmdFunc(cmd *cobra.Command, args []string) error {
	queryFunc, ok := apiMap[ipQueryApi]
	if !ok {
		return fmt.Errorf("IP query API of [%s] not found", ipQueryApi)
	}

	return queryFunc(cmd, args)
}

func queryIpViaChinaz(cmd *cobra.Command, args []string) error {
	resp, err := http.Get("https://ip.tool.chinaz.com/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("api status[%d] incorrect", resp.StatusCode)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// query content with CSS selector
	htmlContent := ""
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	doc.Find(".WhoIpWrap").Each(func(idx int, s *goquery.Selection) {
		html, err := s.Html()
		if err != nil {
			println("error:", err.Error())
		}

		htmlContent = html
	})
	if htmlContent == "" {
		return errors.New("selector .WhoIpWrap without content")
	}

	// fetch ip
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock
	regEx, err := regexp.Compile(regexPattern)
	if err != nil {
		return err
	}
	ips := regEx.FindAllString(htmlContent, -1)
	if len(ips) == 0 {
		return errors.New("no matched IP")
	}

	fmt.Println(ips[0])

	return nil
}

func queryIpViaIcanhazip(cmd *cobra.Command, args []string) error {
	resp, err := http.Get("https://ipv4.icanhazip.com/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("api status[%d] incorrect", resp.StatusCode)
	}

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(strings.Trim(string(ip), "\n"))

	return nil
}
