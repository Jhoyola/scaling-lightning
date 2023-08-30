package scalinglightning

import (
	"fmt"

	sl "github.com/scaling-lightning/scaling-lightning/pkg/network"
	"github.com/spf13/cobra"
)

var channelbalanceCmd = &cobra.Command{
	Use:   "channelbalance",
	Short: "Get the onchain wallet balance of a node",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		processDebugFlag(cmd)
		nodeName := cmd.Flag("node").Value.String()
		slnetwork, err := sl.DiscoverStartedNetwork(kubeConfigPath)
		if err != nil {
			fmt.Printf(
				"Problem with network discovery, is there a network running? Error: %v\n",
				err.Error(),
			)
			return
		}
		allNodes := slnetwork.LightningNodes
		for _, node := range allNodes {
			if node.GetName() == nodeName {
				channelBalance, err := node.ChannelBalance()
				if err != nil {
					fmt.Printf("Problem getting wallet balance: %v\n", err.Error())
					return
				}
				fmt.Printf("%d sats\n", channelBalance.AsSats())
				return
			}
		}
		allNames := []string{}
		for _, node := range allNodes {
			allNames = append(allNames, node.GetName())
		}
		fmt.Printf(
			"Can't find node with name %v, here are the nodes that are running: %v\n",
			nodeName,
			allNames,
		)
	},
}

func init() {
	rootCmd.AddCommand(channelbalanceCmd)

	channelbalanceCmd.Flags().
		StringP("node", "n", "", "The name of the node to get the wallet balance of")
	channelbalanceCmd.MarkFlagRequired("node")
}
