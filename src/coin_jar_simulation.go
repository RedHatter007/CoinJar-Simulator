
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

// cylinderVolume calculates the volume of a cylinder given the radius and height.
func cylinderVolume(radius, height float64) float64 {
    return math.Pi * math.Pow(radius, 2) * height
}

// simulateJarFilling simulates dropping coins into the jar until the specified volume is reached.
func simulateJarFilling(jarVolume, coinVolume, packingDensity float64) int {
    filledVolume := 0.0
    count := 0
    for filledVolume < jarVolume*packingDensity {
        filledVolume += coinVolume
        count++
    }
    return count
}

// monteCarloSimulation runs the simulation a specified number of times and averages the results.
func monteCarloSimulation(simulations int, jarVolume, coinVolume, packingDensity float64) int {
    var totalCoins int
    for i := 0; i < simulations; i++ {
        totalCoins += simulateJarFilling(jarVolume, coinVolume, packingDensity)
    }
    averageCoins := float64(totalCoins) / float64(simulations)
    return int(math.Floor(averageCoins))
}

func main() {
    // Constants for the jar and coins
    const jarDiameterCM = 15.0
    const jarHeightCM = 22.0
    const coinDiameterCM = 4.0
    const coinThicknessCM = 0.2
    const packingDensity = 0.5
    const fullness = 0.9

    // Calculate the volume of the jar (90% full)
    jarRadiusCM := jarDiameterCM / 2
    jarVolume := fullness * cylinderVolume(jarRadiusCM, jarHeightCM)

    // Calculate the volume of a single coin
    coinRadiusCM := coinDiameterCM / 2
    coinVolume := cylinderVolume(coinRadiusCM, coinThicknessCM)

    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Number of simulations to run
    const simulations = 10000

    // Run the Monte Carlo simulation
    estimatedCoins := monteCarloSimulation(simulations, jarVolume, coinVolume, packingDensity)

    // Output the result
    fmt.Printf("Estimated number of coins in the jar: %d\n", estimatedCoins)
}
