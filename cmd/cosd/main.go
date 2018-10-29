package main

import (
	"fmt"
	"github.com/coschain/contentos-go/cmd/cosd/commands"
	log "github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// NO OTHER CONFIGS HERE EXCEPT NODE CONFIG
func cmdRunNode(cmd *cobra.Command, args []string) {
	// _ is cfg as below process has't used
	node, _ := makeConfig()
	if err := node.Start(); err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(1)
	}

	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		log.Info("Got interrupt, shutting down...")
		go node.Stop()
	}()
	node.Wait()
}

// cosd is the main entry point into the system if no special subcommand is pointed
// It creates a default node based on the command line arguments and runs it
// in blocking mode, waiting for it to be shut down.
var rootCmd = &cobra.Command{
	Use:   "cosd",
	Short: "Cosd is a fast blockchain for content",
	Run:   cmdRunNode,
}

func addCommands() {
	rootCmd.AddCommand(commands.InitCmd)
}

func main() {
	addCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
