package main

// у меня депресстя из-за насти поэтому ебашим калькулятор
// проверок нету ибо мне похуй

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// вычисления с помощью Reverse Polish Notation
// Token represents a token in the expression
type Token struct {
	Type  string
	Value float64
}

// Lexer breaks the expression into tokens
func Lexer(expression string) ([]Token, error) {
	tokens := make([]Token, 0)
	i := 0
	for i < len(expression) {
		if expression[i] == ' ' {
			i++
			continue
		}
		if expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.' {
			j := i
			for j < len(expression) && (expression[j] >= '0' && expression[j] <= '9' || expression[j] == '.') {
				j++
			}
			num, err := strconv.ParseFloat(expression[i:j], 64)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, Token{"Number", num})
			i = j
		} else if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/' || expression[i] == '^' {
			tokens = append(tokens, Token{"Operator", 0})
			tokens[len(tokens)-1].Type = string(expression[i])
			i++
		} else {
			return nil, fmt.Errorf("invalid character '%c' at position %d", expression[i], i)
		}
	}
	return tokens, nil
}

// Evaluator evaluates the expression
func Evaluator(tokens []Token) (float64, error) {
	stack := make([]float64, 0)
	operators := make([]Token, 0)
	for _, token := range tokens {
		if token.Type == "Number" {
			stack = append(stack, token.Value)
		} else {
			for len(operators) > 0 && getPrecedence(operators[len(operators)-1].Type) >= getPrecedence(token.Type) {
				op := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				a := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				switch op.Type {
				case "+":
					stack = append(stack, a+b)
				case "-":
					stack = append(stack, a-b)
				case "*":
					stack = append(stack, a*b)
				case "/":
					if b == 0 {
						return 0, fmt.Errorf("division by zero")
					}
					stack = append(stack, a/b)
				case "^":
					stack = append(stack, math.Pow(a, b))
				}
			}
			operators = append(operators, token)
		}
	}
	for len(operators) > 0 {
		op := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		switch op.Type {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			if b == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			stack = append(stack, a/b)
		case "^":
			stack = append(stack, math.Pow(a, b))
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}

// getPrecedence returns the precedence of an operator
func getPrecedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	default:
		return 0
	}
}

// Evaluation evaluates a mathematical expression
func evaluate(expression string) (float64, error) {
	tokens, err := Lexer(expression)
	if err != nil {
		return 0, err
	}
	return Evaluator(tokens)
}

type minimalDarkTheme struct{}

func (m minimalDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 0xbc, G: 0xd1, B: 0xf3, A: 0xff} // background
	case theme.ColorNameButton:
		return color.RGBA{R: 0x6a, G: 0xb1, B: 0x87, A: 0xff} // button background
	case theme.ColorNameForeground:
		return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} // number on button
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0x40, G: 0xa0, B: 0x70, A: 0xff} // Teal Accent
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 0x99, G: 0x99, B: 0x99, A: 0xff} // Light Gray PlaceHolder
	case theme.ColorNameDisabled:
		return color.RGBA{R: 0x66, G: 0x66, B: 0x66, A: 0xff} // Medium Gray Disabled
	case theme.ColorNameDisabledButton:
		return color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff} // Darker Gray Disabled Button
	case theme.ColorNameScrollBar:
		return color.RGBA{R: 0x77, G: 0x77, B: 0x77, A: 0xff} // Medium Gray ScrollBar
	case theme.ColorNameShadow:
		return color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff} // Black Shadow
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (m minimalDarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m minimalDarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m minimalDarkTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 6
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameText:
		return 14
	default:
		return theme.DefaultTheme().Size(name)
	}
}

func main() {
	//СТРОКА ВЫВОДА и сюрприз
	mainS := ""
	mainResult := widget.NewLabel(mainS)
	errorS := ""
	errorResult := widget.NewLabel(mainS)
	mainResult.TextStyle.Bold = true
	go func() {
		for {
			mainResult.SetText(mainS)
			errorResult.SetText(errorS)
		}
	}()

	// базовые настройки
	app := app.New()
	app.Settings().SetTheme(&minimalDarkTheme{})
	w := app.NewWindow("Semculator")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(400, 500))
	img := canvas.NewImageFromFile("./assets/cool.jpg")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(200, 200))
	img.Hide()
	icon, _ := fyne.LoadResourceFromPath("./assets/icon.png")
	w.SetIcon(icon)
	// title := widget.NewLabelWithStyle("Semculator", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	// настройка заголовка
	title := canvas.NewText("Semculator", color.Black)
	title.TextStyle.Bold = true
	title.TextSize = 30
	title.Alignment = fyne.TextAlignCenter
	// кнопки

	zero := widget.NewButton("0", func() {
		if mainS != "" && mainS != "0" {
			mainS += "0"
		}
	})
	one := widget.NewButton("1", func() {
		mainS += "1"
	})
	two := widget.NewButton("2", func() {
		mainS += "2"
	})
	three := widget.NewButton("3", func() {
		mainS += "3"
	})
	four := widget.NewButton("4", func() {
		mainS += "4"
	})
	five := widget.NewButton("5", func() {
		mainS += "5"
	})
	six := widget.NewButton("6", func() {
		mainS += "6"
	})
	seven := widget.NewButton("7", func() {
		mainS += "7"
	})
	eight := widget.NewButton("8", func() {
		mainS += "8"
	})
	nine := widget.NewButton("9", func() {
		mainS += "9"
	})
	clear := widget.NewButton("C", func() {
		mainS = ""
	})

	plus := widget.NewButton("+", func() {
		mainS += "+"
	})

	minus := widget.NewButton("-", func() {
		mainS += "-"
	})

	multiply := widget.NewButton("*", func() {
		mainS += "*"
	})

	divide := widget.NewButton("/", func() {
		mainS += "/"
	})
	pow := widget.NewButton("^", func() {
		mainS += "^"
	})
	equals := widget.NewButton("=", func() {
		res, err := evaluate(mainS)
		if err != nil {
			errorS = "ты еблан"
		} else {
			errorS = ""
		}
		mainS = fmt.Sprintf("%d", int(res))
	})

	surprise := widget.NewButton("Сюрприз", func() {
		img.Show()
	})
	// сетки
	grid := container.NewGridWithColumns(4,
		one,
		two,
		three,
		plus,
		four,
		five,
		six,
		minus,
		seven,
		eight,
		nine,
		multiply,
		clear,
		zero,
		divide,
		pow,
		equals,
		surprise,
		errorResult,
	)
	w.SetContent(mainResult)
	w.SetContent(container.NewVBox(
		title,
	))
	w.SetContent(container.NewVBox(
		title,
		mainResult, grid, img))
	w.ShowAndRun()
}
