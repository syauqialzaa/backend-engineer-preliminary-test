package algorithm

import "fmt"

func ValleyCounter(n int, s string) int {
    level := 0    // Current altitude level
    valleys := 0  // Number of valleys

    for _, step := range s {
        if step == 'U' {
            level++
        } else if step == 'D' {
            level--
            // Check if we are coming up from a valley
            if level == 0 {
                valleys++
            }
        }
    }

    return valleys
}

func CountingValleys() {
    var n int
    fmt.Print("Enter the number of steps: ")
    fmt.Scan(&n)

    var s string
    fmt.Print("Enter the path (U for uphill, D for downhill): ")
    fmt.Scan(&s)

    result := ValleyCounter(n, s)
    fmt.Println("Number of valleys Gary walked through:", result)
}
