
package main

import (
    "testing"
)

// TestCylinderVolume tests the cylinderVolume function for accuracy
func TestCylinderVolume(t *testing.T) {
    // Test cases
    tests := []struct {
        radius, height, expected float64
    }{
        {1.0, 1.0, 3.141592653589793},  // Volume of a cylinder with radius 1 and height 1
        {0.0, 1.0, 0.0},               // Volume should be 0 if radius is 0
        {1.0, 0.0, 0.0},               // Volume should be 0 if height is 0
        // Additional test cases can be added here
    }

    for _, test := range tests {
        if got := cylinderVolume(test.radius, test.height); got != test.expected {
            t.Errorf("cylinderVolume(%f, %f) = %f; want %f", test.radius, test.height, got, test.expected)
        }
    }
}

// TestSimulateJarFilling tests the simulateJarFilling function for basic logic
func TestSimulateJarFilling(t *testing.T) {
    // Assume a jar volume and coin volume for testing
    jarVolume := 100.0    // arbitrary volume of the jar
    coinVolume := 10.0    // arbitrary volume of a coin
    packingDensity := 0.5 // 50% packing density

    // Expect that approximately half the jar volume can be filled with coins
    expectedMin := int((jarVolume * packingDensity) / coinVolume) - 1
    expectedMax := int((jarVolume * packingDensity) / coinVolume) + 1

    got := simulateJarFilling(jarVolume, coinVolume, packingDensity)
    if got < expectedMin || got > expectedMax {
        t.Errorf("simulateJarFilling(%f, %f, %f) = %d; want between %d and %d", jarVolume, coinVolume, packingDensity, got, expectedMin, expectedMax)
    }
}

// TestMonteCarloSimulation tests the monteCarloSimulation function
func TestMonteCarloSimulation(t *testing.T) {
    // This test is more about checking that the function runs and provides a consistent result
    jarVolume := 100.0    // arbitrary volume of the jar
    coinVolume := 10.0    // arbitrary volume of a coin
    packingDensity := 0.5 // 50% packing density
    simulations := 10     // a small number of simulations for the test

    got := monteCarloSimulation(simulations, jarVolume, coinVolume, packingDensity)
    if got <= 0 {
        t.Errorf("monteCarloSimulation(%d, %f, %f, %f) = %d; want a positive number", simulations, jarVolume, coinVolume, packingDensity, got)
    }
}
