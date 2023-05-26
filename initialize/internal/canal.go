package internal

import (
	"telebot/utils"

	"github.com/go-mysql-org/go-mysql/canal"
)

type EventHandler struct {
	canal.DummyEventHandler
}

func (h *EventHandler) OnRow(ev *canal.RowsEvent) error {
	_, rows := utils.GetParseValue(ev)

	if _, ok := utils.InArray(ev.Table.Name, []string{"video_changes"}); ok {
		utils.TableEventDispatcher(ev, rows)
	}

	return nil
}

func (*EventHandler) String() string {
	return "EventHandler"
}
