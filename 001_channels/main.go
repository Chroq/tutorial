package main

import (
	"fmt"
	"time"
)

const BalrogHP = 20

func LegolasShootArrow(damage chan int) {
	for damage != nil {
		damage <- 1
		time.Sleep(100 * time.Millisecond)
	}
}

func GandalfCastsSpell(damage chan int) {
	for damage != nil {
		damage <- 5
		time.Sleep(250 * time.Millisecond)
	}
}

func DisplayBalrogHP(LegolasDamage, GandalfDamage chan int) {
	var balrogHP = BalrogHP
	for LegolasDamage != nil && GandalfDamage != nil {

		var incomingDamage int
		select {
		case incomingDamage = <-LegolasDamage:
			fmt.Println("Legolas shoots an arrow !")
		case incomingDamage = <-GandalfDamage:
			fmt.Println("Gandalf casts a spell !")
		}

		balrogHP -= incomingDamage

		fmt.Printf("Balrog HP : %d\n", balrogHP)

		if balrogHP <= 0 {
			LegolasDamage, GandalfDamage = nil, nil
			fmt.Printf("Balrog is dead !")
		}
	}
}

func main() {
	var LegolasDamage, GandalfDamage = make(chan int), make(chan int)

	go LegolasShootArrow(LegolasDamage)
	go GandalfCastsSpell(GandalfDamage)
	go DisplayBalrogHP(LegolasDamage, GandalfDamage)

	var a string
	fmt.Scanln(&a)
}
