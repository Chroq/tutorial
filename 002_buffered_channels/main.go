package main

import (
	"fmt"
)

const BalrogHP = 20

func LegolasShootArrows(damage chan int) {
	damage <- 1
	damage <- 2
	damage <- 3
	close(damage)
}

func GandalfCastsSpell(dead chan bool, damage chan int) {
	for !<-dead {
		damage <- 5
	}
}

func DisplayBalrogHP(dead chan bool, LegolasDamage, GandalfDamage chan int) {
	var balrogHP = BalrogHP
	fmt.Print("Balrog is coming ! ")
	for balrogHP > 0 {
		if balrogHP > 0 {
			dead <- false
			fmt.Printf("Balrog HP : %d\n", balrogHP)
		} else {
			dead <- true
		}

		var incomingDamage int
		for damage := range LegolasDamage {
			fmt.Printf("Legolas shoots an arrow and deals %d damages ! \n", damage)
			incomingDamage += damage
		}

		var open bool
		if incomingDamage, open = <-GandalfDamage; open {
			fmt.Printf("Gandalf casts a spell and deals %d damages ! \n", incomingDamage)
		}

		balrogHP -= incomingDamage
	}

	fmt.Printf("Balrog is dead !")
}

func main() {
	dead, LegolasDamage, GandalfDamage := make(chan bool), make(chan int, 3), make(chan int)

	go LegolasShootArrows(LegolasDamage)
	go GandalfCastsSpell(dead, GandalfDamage)

	go DisplayBalrogHP(dead, LegolasDamage, GandalfDamage)

	var a string
	fmt.Scanln(&a)
}
