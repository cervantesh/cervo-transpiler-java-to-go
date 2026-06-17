package algorithm

type DiffAlgorithmListener interface {
	DiffStart()
	DiffStep(value int, max int)
	DiffEnd()
}
