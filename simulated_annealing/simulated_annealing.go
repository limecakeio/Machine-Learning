package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	cityCount        = 100
	distanceLimit    = 100
	startTemperature = 1000.0
	epsilon          = 0.01
)

var (
	cityMatrix      [cityCount][cityCount]int
	hypothesis      = newHypothesis()
	savedHypothesis [cityCount]int
	temp            = startTemperature
)

/*newCityMatrix initiates an elementary, hollow matrix filled with random ints */
func newCityMatrix() {
	for i := 0; i < len(cityMatrix); i++ {
		for j := 0; j < len(cityMatrix[i]); j++ {
			if j == i {
				cityMatrix[i][j] = 0
			} else {
				randomDistance := rand.Intn(distanceLimit)
				for randomDistance == 0 { //Distances should between cities a, b are not 0
					randomDistance = rand.Intn(distanceLimit)
				}
				cityMatrix[i][j] = randomDistance
			}
		}
	}
}

/*newHypothesis returns an initial randomly shuffled hypothesis with a length of cityCount*/
func newHypothesis() (hyp [cityCount]int) {
	for i := 0; i < len(hyp); i++ {
		hyp[i] = i
	}
	//Initial travel path is shuffled
	rand.Shuffle(len(hyp), func(i, j int) {
		hyp[i], hyp[j] = hyp[j], hyp[i]
	})
	return hyp
}

/*getDistance returns the distance at the point (x, y) within the cityMatrix*/
func getDistance(x, y int) int {
	return int(cityMatrix[x][y])
}

/*fitness returns the total distance of the hypothesis (tp) in the elemtary matrix (cm)*/
func fitness(hyp [cityCount]int) (fitness int) {
	for i := 0; i < len(hyp)-1; i++ {
		x, y := hyp[i], hyp[i+1]
		fitness -= getDistance(x, y)
	}

	fitness -= getDistance(hyp[len(hyp)-1], hyp[0]) //Return home from last city

	return fitness
}

/*moveOneStepAtRandom takes a hypothesis and pseudo-randomly swaps two points to move one step at random*/
func moveOneStepAtRandom(hypothesis [cityCount]int) (newHypo [cityCount]int) {
	newHypo = hypothesis

	x, y := 0, 0

	for x == y { //Don't swap the same position
		x = rand.Intn(cityCount - 1)
		y = rand.Intn(cityCount - 1)
	}

	//Classic swap
	tmp := newHypo[x]
	newHypo[x] = newHypo[y]
	newHypo[y] = tmp

	return newHypo
}

func main() {
	rand.Seed(time.Now().UnixNano()) //Reseed Random
	newCityMatrix()
	lastFitness := fitness(hypothesis)
	rnd := rand.Float64()
	for temp > epsilon {
		savedHypothesis = hypothesis
		hypothesis = moveOneStepAtRandom(hypothesis)
		newFitness := fitness(hypothesis)
		if newFitness > lastFitness {
			fmt.Printf("Fitness: %v\n", newFitness)
			lastFitness = newFitness
		} else if  rnd < math.Exp((float64(newFitness)-float64(lastFitness))/temp) {
			fmt.Print("###########################################\n")
			fmt.Printf("e^(%+v) = %+v \n", (float64(newFitness)-float64(lastFitness))/temp, math.Exp((float64(newFitness)-float64(lastFitness))/temp))
			fmt.Printf("Go down: %+v \n", newFitness)
			fmt.Print("###########################################\n\n")
			lastFitness = newFitness
		} else {
			hypothesis = savedHypothesis
		}
		temp -= epsilon
	}
}
