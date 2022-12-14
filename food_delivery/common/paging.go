package common

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
	// Support cursor  with UID
	FakeCursor string `json:"fake_cursor" form:"fake_cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Fullfill() {
	if p.Limit == 0 {
		p.Limit = 50
	}
}
