// Package client описывает воронку продаж: этапы сделки и разрешённые переходы.
package client

import "fmt"

// Stage — этап работы с клиентом.
type Stage string

const (
	FirstContact  Stage = "first_contact" // первое касание
	Qualification Stage = "qualification" // квалификация лида
	Needs         Stage = "needs"         // выявление потребностей
	Proposal      Stage = "proposal"      // коммерческое предложение
	Negotiation   Stage = "negotiation"   // переговоры
	Contract      Stage = "contract"      // договор заключён
	Lost          Stage = "lost"          // сделка проиграна
)

// transitions — куда можно перевести сделку с каждого этапа.
var transitions = map[Stage][]Stage{
	FirstContact:  {Qualification, Lost},
	Qualification: {Needs, Lost},
	Needs:         {Proposal, Lost},
	Proposal:      {Negotiation, Needs, Lost},
	Negotiation:   {Contract, Proposal, Lost},
}

// Deal — сделка с клиентом.
type Deal struct {
	ClientName string
	Stage      Stage
	LostReason string
}

// NewDeal создаёт сделку на этапе первого касания.
func NewDeal(clientName string) *Deal {
	return &Deal{ClientName: clientName, Stage: FirstContact}
}

// MoveTo переводит сделку на этап next, если такой переход разрешён.
func (d *Deal) MoveTo(next Stage) error {
	for _, s := range transitions[d.Stage] {
		if s == next {
			d.Stage = next
			return nil
		}
	}
	return fmt.Errorf("переход %s → %s запрещён", d.Stage, next)
}

// Lose закрывает сделку как проигранную с указанием причины.
func (d *Deal) Lose(reason string) error {
	if err := d.MoveTo(Lost); err != nil {
		return err
	}
	d.LostReason = reason
	return nil
}
