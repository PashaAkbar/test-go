package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveCryptarithm(input string) string {
    // Pisahkan input menjadi bagian-bagian
    parts := strings.Fields(input)
    
    // Kumpulkan semua huruf unik
    letters := make(map[rune]bool)
    for _, part := range parts {
        for _, char := range part {
            if char >= 'A' && char <= 'Z' {
                letters[char] = true
            }
        }
    }
    
    // Buat slice dari huruf-huruf unik
    uniqueLetters := make([]rune, 0, len(letters))
    for letter := range letters {
        uniqueLetters = append(uniqueLetters, letter)
    }
    
    // Coba semua kemungkinan angka
    return backtrack(parts, uniqueLetters, make(map[rune]int), 0)
}

func backtrack(parts []string, letters []rune, assigned map[rune]int, index int) string {
    if index == len(letters) {
        if isValid(parts, assigned) {
            return formatSolution(parts, assigned)
        }
        return ""
    }
    
    used := make([]bool, 10)
    for _, v := range assigned {
        used[v] = true
    }
    
    for i := 0; i <= 9; i++ {
        if !used[i] {
            assigned[letters[index]] = i
            if result := backtrack(parts, letters, assigned, index+1); result != "" {
                return result
            }
            delete(assigned, letters[index])
        }
    }
    
    return ""
}

func isValid(parts []string, assigned map[rune]int) bool {
    values := make([]int, len(parts))
    for i, part := range parts {
        if part == "=" || part == "+" || part == "-" {
            continue
        }
        value := 0
        for _, char := range part {
            if char >= 'A' && char <= 'Z' {
                value = value*10 + assigned[char]
            }
        }
        values[i] = value
    }
    
    if parts[1] == "+" {
        return values[0] + values[2] == values[4]
    } else if parts[1] == "-" {
        return values[0] - values[2] == values[4]
    }
    return false
}

func formatSolution(parts []string, assigned map[rune]int) string {
    result := make([]string, len(parts))
    for i, part := range parts {
        if part == "=" || part == "+" || part == "-" {
            result[i] = part
            continue
        }
        var value strings.Builder
        for _, char := range part {
            if char >= 'A' && char <= 'Z' {
                value.WriteString(strconv.Itoa(assigned[char]))
            }
        }
        result[i] = value.String()
    }
    return strings.Join(result, " ")
}

func main() {
    inputs := []string{
        "II + II = JIU",
        "ABD - AD = DKL",
    }
    
    for _, input := range inputs {
        fmt.Printf("Input: %s\n", input)
        fmt.Printf("Output: %s\n\n", solveCryptarithm(input))
    }
}