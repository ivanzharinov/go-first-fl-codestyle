package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	war  = "warrior"
	mage = "mage"
	heal = "healer"
)

func attack(charName, charClass string) string {

	var damage int

	switch charClass {
	case war:
		damage = 5 + randint(3, 5)
	case mage:
		damage = 5 + randint(5, 10)
	case heal:
		damage = 5 + randint(-3, -1)
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, damage)
}

// обратите внимание на "if else" и на "else"
func defence(charName, charClass string) string {

	var block int

	switch charClass {
	case war:
		block = 5 + randint(5, 10)
	case mage:
		block = 5 + randint(-2, 2)
	case heal:
		block = 5 + randint(2, 5)
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s блокировал %d урона.", charName, block)

}

// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {

	var spell string
	var value int

	switch charClass {
	case war:
		spell = "Выносливость"
		value = 80 + 25
	case mage:
		spell = "Атака"
		value = 5 + 40
	case heal:
		spell = "Защита"
		value = 10 + 30
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s применил специальное умение `%s %d`", charName, spell, value)

}

// здесь обратите внимание на имена параметров
func training(charName, charClass string) string {

	var description string
	switch charClass {
	case war:
		description = ", ты Воитель - отличный боец ближнего боя.\n"
	case mage:
		description = ", ты Маг - превосходный укротитель стихий.\n"
	case heal:
		description = ", ты Лекарь - чародей, способный исцелять раны.\n"
	}

	//fmt.Printf("%s %s", charName, description)

	fmt.Printf(`%s %s
Потренируйся управлять своими навыками.
Введи одну из команд: attack — чтобы атаковать противника,
defence — чтобы блокировать атаку противника,
special — чтобы использовать свою суперсилу.
Если не хочешь тренироваться, введи команду skip.
`, charName, description)

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		var action string
		switch cmd {
		case "attack":
			action = attack(charName, charClass)
		case "defence":
			action = defence(charName, charClass)
		case "special":
			action = special(charName, charClass)
		}

		fmt.Println(action)

	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiseClass() string {
	var choice string
	var class string

	for choice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &class)
		var description string
		switch class {
		case war:
			description = "Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный."
		case mage:
			description = "Маг — находчивый воин дальнего боя. Обладает высоким интеллектом."
		case heal:
			description = "Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов."
		}
		fmt.Println(description)

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &choice)
		choice = strings.ToLower(choice)
	}
	return class
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := choiseClass()

	fmt.Println(training(charName, charClass))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
