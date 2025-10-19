package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Global color functions
var(
	hiGreen = color.New(color.FgHiGreen)
	hiRed   = color.New(color.FgHiRed)
	hiBlue  = color.New(color.FgHiBlue)
	hiYellow = color.New(color.FgHiYellow)
	hiCyan  = color.New(color.FgHiCyan)
	hiMagenta = color.New(color.FgHiMagenta)
	hiWhite = color.New(color.FgHiWhite)
)

// Function to print text with typewriter effect
func typewriterPrompt(message string, colorChoice *color.Color) {
	for _, char := range message {
		colorChoice.Printf("%c", char)
		time.Sleep(70 * time.Millisecond)
	}
	
}

// First letter uppercase function
func firstUpper(name string) string {
	nameArr := strings.Split(name, "")
	nameArr[0] = strings.ToUpper(nameArr[0])
	name2 := strings.Join(nameArr, "")
	return name2
}

func startGame(name string) {
	// Placeholder for game logic
	typewriterPrompt("To be continued...\n", hiWhite)
}

func main() {
	fmt.Println()

	// intro
	typewriterPrompt("The Late Shift Game🍕\n", hiCyan)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// variables
	var (
		name string
		// tipsEarned int
		// rentDue = rand.Intn(150) + 150 // rent between 150-300
	)

	// can you paint a nice scene for the use. friday night at a pizzeria, its late, neon lights, rain outside etc.and you get all sorts of different customer. Nice, mean, weird, rich, flirty, cheap atc. rent is due at the end of the month , which is tomorrow. so you need to make as much money as possible. If you dont, you get evicted. goood luck and make good choices!

	typewriterPrompt("Friday night. The city hums like a tired engine that refuses to stop.\n", hiCyan)
	time.Sleep(1 * time.Second)
	typewriterPrompt("Neon lights flicker outside, painting the rain-soaked sidewalk in streaks of pink and red.\n", hiBlue)
	time.Sleep(1 * time.Second)
	typewriterPrompt("The buzz of the pizza oven mixes with the faint sound of thunder in the distance.\n", hiWhite)
	time.Sleep(1 * time.Second)
	typewriterPrompt("You’re the only one working tonight. Everyone else bailed. Figures.\n", hiYellow)
	time.Sleep(1 * time.Second)

	typewriterPrompt("The shop smells like dough, grease, and a little bit of desperation.\n", hiCyan)
	time.Sleep(900 * time.Millisecond)
	typewriterPrompt("Your apron’s stained, your eyes are heavy, and your rent is due — tomorrow.\n", hiRed)
	time.Sleep(1 * time.Second)
	typewriterPrompt("If you don’t make enough in tips tonight… you’re out. Evicted. Game over.\n", hiRed)
	time.Sleep(1 * time.Second)

	typewriterPrompt("But hey — it’s not all bad.\n", hiGreen)
	time.Sleep(900 * time.Millisecond)
	typewriterPrompt("Gus's Pizza is open late, and you get *all kinds* of customers this time of night:\n", hiWhite)
	time.Sleep(1 * time.Second)
	typewriterPrompt("The nice ones who make you smile.\n", hiGreen)
	typewriterPrompt("The cheap ones who make you question your life choices.\n", hiYellow)
	typewriterPrompt("The weird ones who talk to their slices.\n", hiCyan)
	typewriterPrompt("And the flirty ones who tip... depending on how you play your cards.\n", hiMagenta)
	time.Sleep(1 * time.Second)

	// get player name
	fmt.Println()
	typewriterPrompt("Before we get started, what’s your name?", hiWhite)
	fmt.Print("> ")
	fmt.Scanln(&name)
	fmt.Println()
	
	// welcome message
	typewriterPrompt("Welcome to the late shift, "+firstUpper(name)+".\n", hiCyan)
	time.Sleep(900 * time.Millisecond)

	typewriterPrompt("Your goal is simple:\n", hiYellow)
	typewriterPrompt("Survive the night.\n", hiRed)
	typewriterPrompt("Make enough tips to pay the rent.\n", hiGreen)
	typewriterPrompt("And try not to lose your mind in the process.\n", hiMagenta)
	time.Sleep(1 * time.Second)
	fmt.Println()

	typewriterPrompt("You take a deep breath. The rain picks up outside.\n", hiBlue)
	time.Sleep(900 * time.Millisecond)
	typewriterPrompt("A neon reflection dances across the counter.\n", hiWhite)
	time.Sleep(900 * time.Millisecond)
	typewriterPrompt("Then — *DING!*\n", hiMagenta)
	time.Sleep(700 * time.Millisecond)
	typewriterPrompt("The bell above the door chimes.\n", hiCyan)
	time.Sleep(900 * time.Millisecond)
	typewriterPrompt("Your first customer of the night walks in...\n", hiYellow)

	// start game
	fmt.Println()
	startGame(name)
}




