package cars
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var a = successRate/100
    return float64(productionRate)*a
}
// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var b = CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(b)/60
    
}
const singlecar = 10000
const tencar = 95000
// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
    oneCarCost := uint(singlecar * (carsCount%10))
    tenCarCost := uint(tencar * (carsCount/10))
    return oneCarCost + tenCarCost
}
