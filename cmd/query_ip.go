package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(queryIpCmd)
}

var queryIpCmd = &cobra.Command{
	Use:   "queryIp",
	Short: "Query dynamic ip",
	RunE:  queryIpCmdFunc,
}

func queryIpCmdFunc(cmd *cobra.Command, args []string) error {
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
