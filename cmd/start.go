package cmd

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kmdkuk/try-ebiten/pkg/game"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		ebiten.SetWindowSize(640, 480)
		ebiten.SetWindowTitle("Hello, World!")
		if err := ebiten.RunGame(game.NewGame()); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
