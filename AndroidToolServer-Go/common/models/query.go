package models

const PAGE_SIZE = 5 //默认页码数据大小

type PageReq struct {
	PageNum  int `json:"pageNum"`  //页数
	PageSize int `json:"pageSize"` //每一页大小
}

func InitPageIfAbsent(page, pageSize *int) {
	if *page <= 0 {
		*page = 1
	}
	if *pageSize <= 0 {
		*pageSize = PAGE_SIZE
	}
}
