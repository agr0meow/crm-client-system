package client

import "testing"

func TestHappyPath(t *testing.T) {
	d := NewDeal("Тестовый клиент")
	for _, s := range []Stage{Qualification, Needs, Proposal, Negotiation, Contract} {
		if err := d.MoveTo(s); err != nil {
			t.Fatalf("неожиданная ошибка: %v", err)
		}
	}
	if d.Stage != Contract {
		t.Fatalf("ожидался этап %s, получен %s", Contract, d.Stage)
	}
}

func TestForbiddenJump(t *testing.T) {
	d := NewDeal("Тестовый клиент")
	if err := d.MoveTo(Contract); err == nil {
		t.Fatal("переход сразу к договору должен быть запрещён")
	}
}

func TestLose(t *testing.T) {
	d := NewDeal("Тестовый клиент")
	if err := d.Lose("нет бюджета"); err != nil {
		t.Fatal(err)
	}
	if d.Stage != Lost || d.LostReason != "нет бюджета" {
		t.Fatalf("сделка закрыта неверно: %+v", d)
	}
}
