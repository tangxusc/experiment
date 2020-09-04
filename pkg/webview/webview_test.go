package webview

import (
	"fmt"
	"github.com/webview/webview"
	"io/ioutil"
	"net/http"
	"testing"
)

//CheckNetIsolation.exe LoopbackExempt -a -n="Microsoft.Win32WebViewHost_cw5n1h2txyewy"
//修正localhost不工作
func TestWebServer(t *testing.T) {
	go func() {
		http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
			file, err := ioutil.ReadFile("index.html")
			if err != nil {
				panic(err.Error())
			}
			fmt.Fprint(writer, string(file))
		})
		http.ListenAndServe(":80", nil)

	}()

	select {}

}

func TestView(t *testing.T) {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("myshell")
	w.Bind("quit", func() {
		w.Terminate()
	})
	w.Navigate("http://127.0.0.1/index")

	w.Run()
}
