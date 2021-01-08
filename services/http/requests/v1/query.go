package v1

// ListQuery using for ...
type ListQuery struct {
	Paginate bool   `form:"paginate" default:"false"`
	PerPage  string `form:"per_page" default:"15"`
	Search   string `form:"search" default:""`
	OrderBy  string `form:"order_by" default:"atoz"`
	Page     string `form:"page" default:"1"`
}
