package adminupdateconfigstring

import "github.com/SongZihuan/cat-shop/shop-backend/src/model/modeltype"

type Query struct {
	Key   modeltype.ConfigKeyType   `form:"key"`
	Value modeltype.ConfigValueType `form:"value"`
}
