package main

import "fmt"

type concert interface {
	getPrice() int
}

type ticket struct {
}

func (t *ticket) getPrice() int {
	return 10000
}

type vipZone struct {
	concert concert
}

func (o *vipZone) getPrice() int {
	ticketPrice := o.concert.getPrice()
	return ticketPrice + 6000
}

type armyBomb struct {
	concert concert
}

func (o *armyBomb) getPrice() int {
	ticketPrice := o.concert.getPrice()
	return ticketPrice + 6000
}

func main() {

	concertTicket := &ticket{}

	concertWithArmyBomb := &armyBomb{
		concert: concertTicket,
	}

	concertWithArmyBombInVipZone := &vipZone{
		concert: concertWithArmyBomb,
	}

	fmt.Printf("Price of BTS concert with army bomb in vip zone is %d\n", concertWithArmyBombInVipZone.getPrice())

}
