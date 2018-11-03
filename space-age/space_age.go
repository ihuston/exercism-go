// Provides Age using a variety of planets
package space

type Planet string

var conversion = map[Planet]float64{
    "Earth": 31557600,
    "Mercury": 31557600*0.2408467,
    "Venus": 31557600*0.61519726,
    "Mars": 31557600*1.8808158,
    "Jupiter": 31557600*11.862615,
    "Saturn": 31557600*29.447498,
    "Uranus": 31557600*84.016846,
    "Neptune": 31557600*164.79132,
}

// Age converts seconds to a year on a named planet
func Age(seconds float64, planet Planet) float64 {
    return seconds/conversion[planet]
}

