package paging

import "gitlab.com/mcsolutions/find-psy/proto/gen/go/common"

type OrderOffsetLimit struct {
	Order  string
	Offset uint32
	Limit  uint32
}

func GetOrderOffsetLimit(ool *common.OrderOffsetLimit) *OrderOffsetLimit {
	return &OrderOffsetLimit{
		Order:  ool.Order,
		Offset: ool.Offset,
		Limit:  ool.Limit,
	}
}

type Paging struct {
	Order   string
	Offset  uint32
	Limit   uint32
	Total   int
	HasNext bool
}

func (p *Paging) GetReply() *common.Paging {
	return &common.Paging{
		Order:   p.Order,
		Offset:  p.Offset,
		Limit:   p.Limit,
		Total:   uint32(p.Total),
		HasNext: p.HasNext,
	}
}
