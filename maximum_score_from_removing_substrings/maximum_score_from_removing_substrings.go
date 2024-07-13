package main

func min(a,b int) int{
    if a>b {
        return b
    }
    return a
}

func maximumGain(s string, x int, y int) int {
    maxScore := 0
    aCount, bCount := 0, 0
    less := min(x,y)
    for i:=0; i<len(s); i++ {
        if string(s[i]) == "a"  {
            if bCount > 0 && y > x { // ba
                maxScore += y
                bCount--
            } else {
                aCount++
            }
        } else if string(s[i]) == "b" { 
            if aCount > 0  && x > y { // ab
                maxScore += x
                aCount--
            } else {
                bCount++
            }
        } else {
            maxScore += min(aCount, bCount) * less
            aCount = 0
            bCount = 0
        }
    }
    maxScore += min(aCount, bCount) * less
    return maxScore
}