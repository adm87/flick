package cmd

import (
	"net/http"
	"os"
	"path"
	"path/filepath"

	_ "net/http/pprof"

	assetcmd "github.com/adm87/flick/cmd/assets"

	"github.com/adm87/flick/data"
	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/game"
	"github.com/spf13/cobra"
)

func Boot(g game.Game) *cobra.Command {
	var (
		root    string
		profile bool
	)

	c := &cobra.Command{
		Use:   "finch",
		Short: "Finch Game",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			root, err := filepath.Abs(root)
			if err != nil {
				g.Log().Error("failed to determine absolute path for root", "error", err)
				return err
			}

			assets.SetRoot(root)
			assets.RegisterFilesystem("assets", os.DirFS(path.Join(root, "data", "assets")))
			assets.RegisterFilesystem("embedded", data.EmbeddedFS)
			assets.RegisterImporters(g.Log())

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if profile {
				go func() {
					g.Log().Info("starting profiler on :6060")
					if err := http.ListenAndServe(":6060", nil); err != nil {
						g.Log().Error("profiler failed", "error", err)
					}
				}()
			}
			if err := g.Run(); err != nil {
				g.Log().Error("game exited with error", "error", err)
				return err
			}
			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	c.PersistentFlags().StringVar(&root, "root", ".", "Root directory of the game assets")
	c.PersistentFlags().BoolVar(&profile, "profile", false, "Enable profiling")

	c.AddCommand(assetcmd.GenerateHandles(g.Log()))

	return c
}
