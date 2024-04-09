package functions

import (
	"CE366VerticalStress/types"
	"fmt"
	"math"
)

func CalculateAnglesFromCoordinates(B, y, z float64) (float64, float64) {
	var alpha = CalculateAlpha(y, z, B)
	var bPrime = CalculateBetaPrime(y, z, B)
	var beta = CalculateBeta(alpha, bPrime)
	return alpha, beta
}

func CalculateSigmaVFromAngles(Q, alpha, beta float64) float64 {
	return (Q / math.Pi) * (alpha + (math.Sin(alpha) * math.Cos(2*beta)))
}

func CalculateSigmaVFromCoordinates(Q, B, y, z float64) float64 {
	alpha, beta := CalculateAnglesFromCoordinates(B, y, z)
	return CalculateSigmaVFromAngles(Q, alpha, beta)
}

func CalculateBetaPrime(y, z, B float64) float64 {
	return math.Atan((y - (0.5 * B)) / z)
}
func CalculateAlpha(y float64, z float64, B float64) float64 {
	return math.Atan((y+(0.5*B))/z) - math.Atan((y-(0.5*B))/z)
}
func CalculateBeta(alpha, betaPrime float64) float64 {
	return betaPrime + (alpha / 2)
}

func CheckDeviation(goal, sigmaV float64) (bool, types.Direction) {
	fmt.Printf("#checkdeviation->goal and sigmaV  is %v, %v\n\n", goal, sigmaV)
	deviation := goal - sigmaV // this will be positive if the point is further down, so direction should be changed to upward in case
	if deviation <= 0.0 {
		fmt.Printf("#checkdeviation->less than 0 is %v\n\n", deviation)
		d := math.Abs(deviation)
		if d <= types.AllowedDeviation {
			return true, types.STOP
		}
		return false, types.UP
	} else if deviation > 0.0 {
		fmt.Printf("#checkdeviation->greater than 0 is %v\n\n", deviation)
		if deviation <= types.AllowedDeviation {
			return true, types.STOP
		}
		return false, types.DOWN
	}
	fmt.Print(fmt.Errorf("unexpected error in deviation check \n"))
	return true, types.STOP
}

func CheckIfItsOnBottom(B, y, z float64) bool {
	//arbitrary check
	if z > 0.5 {
		dev := y - (B / 2)
		if math.Abs(dev) < 0.1 {
			return true
		}
	}
	return false
}
func IsFluctuating(trace []types.Direction) bool {
	if len(trace) <= 1 {
		return false
	}
	directionBeforeLast := trace[len(trace)-2]
	lastDirection := trace[len(trace)-1]
	if lastDirection == directionBeforeLast {
		return false
	}
	return true
}
func NextY(y *float64, direction types.Direction, trace []types.Direction) {
	if *y <= 0.0001 {
		*y = *y * 1.2
		fmt.Println("#2->0 next direction not applicable")
		return
	}
	fluctuation := IsFluctuating(trace)
	if direction == types.UP {
		fmt.Println("#2->1 next direction UP")
		if fluctuation {
			fmt.Println("#2->1 fluctuating")
			*y = *y * 1.05
			return
		} else {
			*y = *y * 1.09
		}
	} else if direction == types.DOWN {
		fmt.Println("#2->2 next direction DOWN")
		if fluctuation {
			fmt.Println("#2->2 fluctuating")
			*y = *y * 0.98
			return
		} else {
			*y = *y * 0.91
		}
	}

}

/*
\begin{equation}
\frac{Q}{\pi} \left( \text{atan} \left( \frac{y + \frac{B}{2}}{z} \right) - \text{atan} \left( \frac{y - \frac{B}{2}}{z} \right) + \sin \left( \text{atan} \left( \frac{y + \frac{B}{2}}{z} \right) - \text{atan} \left( \frac{y - \frac{B}{2}}{z} \right) \right) \cdot \cos \left( 2 \cdot \text{atan} \left( \frac{y - \frac{B}{2}}{z} \right) + \frac{\text{atan} \left( \frac{y + \frac{B}{2}}{z} \right) - \text{atan} \left( \frac{y - \frac{B}{2}}{z} \right)}{2} \right) \right)
\end{equation}
*/
