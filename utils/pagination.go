package utils

type Pagination struct {
	Page       int
	PageSize   int
	Count      int
	TotalPages int
	PrePage    int
	NextPage   int
	Ranges     []int
}

//Paginate分页
type Paginate struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Pages    int `json:"pages"`
}

func Set(page, pageSize, count int) Pagination {
	totalPages := count / pageSize
	if count%pageSize > 0 {
		totalPages += 1
	}

	var prePage int
	if page > 1 {
		prePage = page - 1
	} else {
		prePage = page
	}

	var nextPage int
	if page < totalPages {
		nextPage = page + 1
	} else {
		nextPage = page
	}

	var ranges []int
	if page == 1 {
		max := 0
		if totalPages <= 3 {
			max = totalPages
		} else {
			max = 3
		}
		for i := page; i <= max; i++ {
			ranges = append(ranges, i)
		}

	} else if page == totalPages {
		min := 0
		if totalPages >= 3 {
			min = totalPages - 2
		} else {
			min = page
		}

		for i := min; i <= totalPages; i++ {
			ranges = append(ranges, i)
		}
	} else {
		for i := page - 1; i <= page+1; i++ {
			ranges = append(ranges, i)
		}
	}

	var pg Pagination
	pg.Page = page
	pg.PageSize = pageSize
	pg.PrePage = prePage
	pg.TotalPages = totalPages
	pg.Count = count
	pg.NextPage = nextPage
	pg.Ranges = ranges
	return pg
}
