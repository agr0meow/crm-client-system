package main

import (
	"fmt"
	"log"

	"github.com/agr0meow/crm-client-system/internal/client"
)

func main() {
	deal := client.NewDeal("ООО «Ромашка»")

	path := []client.Stage{
		client.Qualification,
		client.Needs,
		client.Proposal,
		client.Negotiation,
		client.Contract,
	}
	for _, stage := range path {
		if err := deal.MoveTo(stage); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: этап %s\n", deal.ClientName, deal.Stage)
	}
}
