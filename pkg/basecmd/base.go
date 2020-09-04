package basecmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
)

func NewCommand() (*cobra.Command, context.Context, context.CancelFunc) {
	cancel, cancelFunc := context.WithCancel(context.TODO())
	config := &ApplicationConfig{}
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

			err := viper.Unmarshal(config)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%+v", config)

			//
			<-cancel.Done()
			return nil
		},
	}
	//command.AddCommand()
	viper.AutomaticEnv()
	viper.AddConfigPath(`.`)
	command.PersistentFlags().String("test", "abcd", "abcd")
	command.PersistentFlags().StringToString("a", nil, "")
	err := viper.BindPFlags(command.PersistentFlags())
	if err != nil {
		panic(err.Error())
	}
	err = viper.BindPFlags(command.Flags())
	if err != nil {
		panic(err.Error())
	}
	viper.SetDefault("test", "dcba")

	return command, cancel, cancelFunc
}

type ApplicationConfig struct {
	Test string
	A    map[string]string
}
