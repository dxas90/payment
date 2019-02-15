package event_handler

import (
	"github.com/screwyprof/payment/internal/pkg/observer"
	"github.com/screwyprof/payment/pkg/event"
	"github.com/screwyprof/payment/pkg/report"
)

type MoneyReceived struct {
	accountReporter report.AccountUpdater
}

func NewMoneyReceived(accountReporter report.AccountUpdater) *MoneyReceived {
	return &MoneyReceived{
		accountReporter: accountReporter,
	}
}

func (h *MoneyReceived) Handle(e observer.Event) {
	evn, ok := e.(event.MoneyReceived)
	if !ok {
		return
	}

	//fmt.Printf("MoneyReceivedEventHandler: %s - %s => %s, %s = %s\n",
	//	evn.From, evn.Amount.Display(), evn.To, evn.To, evn.Balance.Display())
	h.accountReporter.Update(&report.Account{Number: string(evn.To), Balance: evn.Balance})
}