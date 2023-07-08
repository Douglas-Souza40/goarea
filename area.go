package goarea

import "math"

// PI é um proporção numerica definida pela relação entre  o perimetro de uma circuferenci e seu diametro
const Pi = 3.1416

//Circ é responsavel por calcular a área da circuferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura / 2)
}
