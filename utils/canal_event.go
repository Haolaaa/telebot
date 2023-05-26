package utils

import (
	"encoding/json"
	"fmt"
	"telebot/global"
	"telebot/model"
	"time"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func GetParseValue(canalRowsEvent *canal.RowsEvent) (*canal.RowsEvent, map[string]interface{}) {
	var (
		rows = map[string]interface{}{}
	)

	for columnIndex, currColumn := range canalRowsEvent.Table.Columns {
		columnValue := canalRowsEvent.Rows[len(canalRowsEvent.Rows)-1][columnIndex]
		switch columnValue.(type) {
		// parse text field into string
		case []uint8:
			columnValue = Uint8ToString(columnValue.([]uint8)...)
		case decimal.Decimal:
			v, _ := columnValue.(decimal.Decimal).Float64()
			columnValue = float32(v)
		case string:
			switch currColumn.Name {
			case "created_at", "updated_at":
				if columnValue.(string) != "" && canalRowsEvent.Table.Name == "video_changes" {
					var err error
					columnValue, err = time.Parse("2006-01-02 15:04:05", columnValue.(string))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}

		rows[currColumn.Name] = columnValue
	}

	return canalRowsEvent, rows
}

func TableEventDispatcher(canalRowsEvent *canal.RowsEvent, row map[string]interface{}) {
	modelStruct := model.GetModelStruct(canalRowsEvent.Table.Name)
	rowModel, ok := MapStructureRow(canalRowsEvent, modelStruct, row)
	if !ok {
		return
	}

	switch canalRowsEvent.Action {
	case canal.InsertAction:
		switch canalRowsEvent.Table.Name {
		case "video_changes":
			global.LOG.Info("insert video_changes", zap.Any("row", rowModel))
			_, err := json.Marshal(rowModel)
			if err != nil {
				global.LOG.Error("json.Marshal", zap.Error(err))
			}
		}
	}
}

func MapStructureRow(canalRowsEvent *canal.RowsEvent, model interface{}, row map[string]interface{}) (interface{}, bool) {
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &model,
	})
	err := decoder.Decode(row)
	if err != nil {
		return nil, false
	}

	return model, true
}
