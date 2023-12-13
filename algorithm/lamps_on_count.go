package algorithm

import "fmt"

func ToggleLamps(n int) []bool {
    lamps := make([]bool, n+1)

    for trip := 1; trip <= n; trip++ {
        for switchNumber := trip; switchNumber <= n; switchNumber += trip {
            lamps[switchNumber] = !lamps[switchNumber]
        }
    }

    return lamps
}

func CountOnLamps(lamps []bool) int {
    count := 0
    for _, lampState := range lamps {
        if lampState {
            count++
        }
    }
    return count
}

func LampsOnCount() {
    n := 100
    lamps := ToggleLamps(n)
    result := CountOnLamps(lamps)

    fmt.Printf("After 100 trips, the number of lamps turned on is: %d\n", result)
}
