package filemeta

import (
	"fmt"
	"testing"
)

func TestGetMeta(t *testing.T) {
	err := GetMeta("C:\\数据\\TCE视频\\TCS培训.mp4")
	fmt.Println(err)
}
