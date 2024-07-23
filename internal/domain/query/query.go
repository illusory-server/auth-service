package query

const (
	Asc  Order = "ASC"
	Desc Order = "DESC"
)

type (
	Order string

	PaginationQuery struct {
		Page      uint
		Limit     uint
		SortBy    string
		SortOrder Order
	}
)

func (p PaginationQuery) Offset() uint {
	return (p.Page - 1) * p.Limit
}
