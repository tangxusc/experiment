package main

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"testgo/pkg/myshell/config"
	"testgo/pkg/myshell/viewer"
	"testgo/pkg/myshell/web"
)

func NewCommand() (*cobra.Command, context.Context, context.CancelFunc) {
	cancel, cancelFunc := context.WithCancel(context.TODO())
	cfg := &config.ApplicationConfig{}
	command := &cobra.Command{
		Use:   ``,
		Short: ``,
		Long:  ``,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			go func() {
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Kill)
				<-c
				cancelFunc()
			}()

			err := viper.Unmarshal(cfg)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			//run web server
			go web.StartWebServer(cancel, cfg)
			//run viewer
			go viewer.StartViewer(cancel, cancelFunc, cfg)

			<-cancel.Done()
			return nil
		},
	}
	//command.AddCommand()
	viper.AutomaticEnv()
	viper.AddConfigPath(`.`)
	command.PersistentFlags().Int("port", 9999, "web端口")
	command.PersistentFlags().Bool("debug", true, "debug模式")
	err := viper.BindPFlags(command.PersistentFlags())
	if err != nil {
		panic(err.Error())
	}
	err = viper.BindPFlags(command.Flags())
	if err != nil {
		panic(err.Error())
	}
	viper.SetDefault("port", 9999)
	viper.SetDefault("debug", true)

	return command, cancel, cancelFunc
}

func main() {
	command, _, _ := NewCommand()
	err := command.Execute()
	if err != nil {
		panic(err.Error())
	}
}
