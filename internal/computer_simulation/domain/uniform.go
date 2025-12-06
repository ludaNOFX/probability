package domain

type UniformGenerator interface {
	// возвращает U ∈ (0,1)
	Float64() float64
}
