package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"gtool/pkg/common"
	"gtool/pkg/flen"
	"gtool/pkg/utCover"
	"strings"
)

// Config ...
type Config struct {
	Flen    bool
	Gocyclo bool

	Help map[string]string
}

// CmdInfo ...
type CmdInfo struct {
	name  string
	isRun bool
}

var (
	flenCfg = Config{}
	rootCmd = &cobra.Command{Use: "gtool"}
)

// CmdFlen ...
var CmdFlen = &cobra.Command{
	Use:   "flen [string to flen]",
	Short: "flen anything to the screen",
	Long: `flen is for printing anything back to the screen.
            For many years people have printed back to the screen.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		pkgs := common.ReadArgs().CheckPkgs
		flen.GenerateFuncLens(pkgs[0])
	},
}

// CmdGocyclo ...
var CmdGocyclo = &cobra.Command{
	Use:   "gocyclo [string to gocyclo]",
	Short: "gocyclo anything to the screen",
	Long: `gocyclo is for echoing anything back.
            gocyclo works a lot like print, except it has a child command.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gocyclo: " + strings.Join(args, " "))
	},
}

// Utcover ...
var CmdUtCover = &cobra.Command{
	Use:   "utCover [string to utCover]",
	Short: "utCover anything to the screen",
	Long: `utCover is for test pkg's ytCover'.
            utCover works a lot like print, except it has a child command.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("checkPath:", common.ReadArgs().CheckPkgs)
		utCover.UtCover(common.ReadArgs().CheckPkgs)
	},
}

// CmdConfig ...
func CmdConfig() {
	CmdFlen.Flags().BoolVarP(&flenCfg.Flen, "test", "t", false, "test to add test files")
	rootCmd.AddCommand(CmdUtCover, CmdFlen, CmdGocyclo)
	rootCmd.Execute()

}
