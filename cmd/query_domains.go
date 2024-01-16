package cmd

import "github.com/spf13/cobra"


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

    return nil
}
