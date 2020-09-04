package viewer

import (
	"context"
	"fmt"
	"github.com/webview/webview"
	"testgo/pkg/myshell/config"
	"testgo/pkg/myshell/web"
)

func StartViewer(ctx context.Context, cancelFunc context.CancelFunc, config *config.ApplicationConfig) {
	w := webview.New(config.Debug)
	defer w.Destroy()
	w.SetTitle("myshell")
	w.Bind("quit", func() {
		cancelFunc()
		w.Terminate()
	})
	w.Navigate(fmt.Sprintf("http://127.0.0.1:%v%s", config.Port, web.IndexPath))
	w.Run()
}
