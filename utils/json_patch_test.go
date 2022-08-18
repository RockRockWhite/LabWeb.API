package utils

import (
	"fmt"
	"gin/entities"
	"testing"
)

func TestApplyJsonPatch(t *testing.T) {
	article := entities.Article{
		UserId:   0,
		Title:    "标题",
		Content:  "内容",
		Views:    0,
		Tags:     nil,
		Comments: nil,
		Stars:    nil,
	}
	patchJson := []byte("[{ \"op\": \"replace\", \"path\": \"/Title\", \"value\": \"标题patched\" },{ \"op\": \"replace\", \"path\": \"/Content\", \"value\": \"内容patched\" }]")
	ApplyJsonPatch(&article, patchJson)

	t.Log(fmt.Sprintf("%+v", article))
}
