package common

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

// Pagination 分页器
type Pagination struct {
	Request *http.Request
	Total   int64
	Size    int
	Pname   string
	Page    int
}

// NewPagination 新建分页器
func NewPagination(req *http.Request, total int64, size int, pname string, page int) *Pagination {
	return &Pagination{
		Request: req,
		Total:   total,
		Size:    size,
		Pname:   pname,
		Page:    page,
	}
}

// Pages 渲染生成html分页标签
func (p *Pagination) Pages() string {

	//计算总页数
	var totalPageNum = int(math.Ceil(float64(p.Total) / float64(p.Size)))

	//首页链接
	var firstLink string
	//上一页链接
	var prevLink string
	//下一页链接
	var nextLink string
	//末页链接
	var lastLink string

	var pageLinks string
	//中间页码链接
	// var pageLinks []string

	//首页和上一页链接
	if p.Page > 1 {
		firstLink += fmt.Sprintf(`<li class="page-item"><a class="page-link" href="%s">首页</a></li>`, p.pageURL("1"))
		prevLink = fmt.Sprintf(`<li class="page-item"><a class="page-link" href="%s">&laquo;</a></li>`, p.pageURL(strconv.Itoa(p.Page-1)))
	} else {
		firstLink += `<li class="page-item disabled"><a class="page-link" href="#">首页</a></li>`
		prevLink = `<li class="page-item disabled"><a class="page-link" href="#">&laquo;</a></li>`
	}

	//末页和下一页
	if p.Page < totalPageNum {
		nextLink = fmt.Sprintf(`<li class="page-item"><a class="page-link" href="%s">&raquo;</a></li>`, p.pageURL(strconv.Itoa(p.Page+1)))
		lastLink = fmt.Sprintf(`<li class="page-item"><a class="page-link" href="%s">尾页</a></li>`, p.pageURL(strconv.Itoa(totalPageNum)))
	} else {
		nextLink = `<li class="page-item disabled"><a class="page-link" href="#">&raquo;</a></li>`
		lastLink = `<li class="page-item disabled"><a class="page-link" href="#">尾页</a></li>`
	}

	//生成中间页码链接
	// pageLinks = make([]string, 0, 10)
	startPos := p.Page - 3
	endPos := p.Page + 3
	if startPos < 1 {
		endPos = endPos + int(math.Abs(float64(startPos))) + 1
		startPos = 1
	}
	if endPos > totalPageNum {
		endPos = totalPageNum
	}
	for i := startPos; i <= endPos; i++ {
		if i == p.Page {
			pageLinks += fmt.Sprintf(`<li class="page-item active"><a class="page-link" href="%s">%d</a></li>`, p.pageURL(strconv.Itoa(i)), i)
		} else {
			pageLinks += fmt.Sprintf(`<li class="page-item"><a class="page-link" href="%s">%d</a></li>`, p.pageURL(strconv.Itoa(i)), i)
		}
	}
	return fmt.Sprintf(`<ul class="pagination">%s%s%s%s%s</ul>`, firstLink, prevLink, pageLinks, nextLink, lastLink)
}

// pageURL 生成分页url
func (p *Pagination) pageURL(page string) string {
	//基于当前url新建一个url对象
	u, _ := url.Parse(p.Request.URL.String())
	q := u.Query()
	q.Set(p.Pname, page)
	u.RawQuery = q.Encode()
	return u.String()
}
