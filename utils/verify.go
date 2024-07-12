package utils

var (
	GoodsCategoryVerify    = Rules{"CategoryRank": {NotEmpty()}, "CategoryName": {NotEmpty()}}
	MallUserRegisterVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	GoodsAddParamVerify    = Rules{"GoodsName": {Le("128")}, "GoodsIntro": {Le("200")}, "GoodsCategoryId": {Ge("1")}, "GoodsCoverImg": {NotEmpty()}, "OriginalPrice": {Ge("1"), Le("1000000")},
		"sellingPrice": {Ge("1"), Le("1000000")}, "stockNum": {Ge("1"), Le("100000")}, "Tag": {Le("16")}, "goodsDetailContent": {NotEmpty()}}
)
