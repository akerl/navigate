package cmd

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

func goRunner(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	ws, err := flags.GetString("ws")
	if err != nil {
		return err
	}

	if len(args) < 1 {
		return fmt.Errorf("must provide url")
	} else if len(args) > 1 {
		return fmt.Errorf("too many args provided")
	}
	url := args[0]

	allocatorContext, _ := chromedp.NewRemoteAllocator(context.Background(), ws)
	ctx, _ := chromedp.NewContext(allocatorContext)

	tabs, err := chromedp.Targets(ctx)
	if err != nil {
		return err
	}
	if len(tabs) == 0 {
		return fmt.Errorf("no tabs found")
	}

	tabCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(tabs[0].TargetID))

	return chromedp.Run(tabCtx, chromedp.Navigate(url))
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Navigate to the provided URL",
	RunE:  goRunner,
}

func init() {
	rootCmd.AddCommand(goCmd)
	goCmd.Flags().StringP("ws", "w", "ws://localhost:8080", "Chrome DevTools websocket URL")
}
