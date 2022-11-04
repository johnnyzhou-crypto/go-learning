package main

import "my-test-service/mock/spider"

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
