package goarea

const Pi = 3.1415

func Circle(r float64) float64 {
	return Pi * r * r
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
