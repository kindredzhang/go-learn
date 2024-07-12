package response

import "go-learn/model/manage"

type GoodsCategoryResponse struct {
	GoodsCategory manage.MallGoodsCategory `json:"mallGoodsCategory"`
}
