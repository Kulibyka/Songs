package models

type Pagination struct {
	PageNum  int `json:"page_num"`  // Номер страницы
	LimitNum int `json:"limit_num"` // Количество элементов на странице
}
