package main

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"my-test-service/mock/spider"
	"testing"
)

// mockgen -destination spider/mock_spider.go -package spider github.com/cz-it/blog/blog/Go/testing/gomock/example/spider Spider
func TestGetVersion(t *testing.T) {
	mockController := gomock.NewController(t)
	mockSpider := spider.NewMockSpider(mockController)
	mockSpider.EXPECT().GetBody().Return("go1.19.2")
	goVersion := GetGoVersion(mockSpider)

	if goVersion != "go1.19.2" {
		fmt.Println("wrong target version")
	}

}
