// @author AlphaSnow

package transform

import (
	"fmt"
	"github.com/jinzhu/copier"
)

var IntToStrConverter = copier.TypeConverter{
	SrcType: uint(0),
	DstType: copier.String,
	Fn: func(src interface{}) (interface{}, error) {
		return fmt.Sprintf("%d", src), nil
	},
}

func SelectCopyOption(src interface{}, dst interface{}, labelField string, valueField string) copier.Option {
	return copier.Option{
		Converters: []copier.TypeConverter{
			IntToStrConverter,
		},
		FieldNameMapping: []copier.FieldNameMapping{
			{
				SrcType: src,
				DstType: dst,
				Mapping: map[string]string{
					valueField: "Value",
					labelField: "Label",
				},
			},
		}}
}
