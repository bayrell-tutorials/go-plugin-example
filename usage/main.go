package main

import (
    "fmt"
)

func main() {
	
    pluginPath := "../library/libsupermath.so"
    libSuperMath, err := LoadSuperMath(pluginPath)
    if err != nil {
        fmt.Println("Ошибка при загрузке плагина:", err)
        return
    }
	
    a := 5
    b := 7
    addResult := libSuperMath.Add(a, b)
    subResult := libSuperMath.Sub(a, b)
    multResult := libSuperMath.Mult(a, b)

    fmt.Printf("Результаты операций для %d и %d:\n", a, b)
    fmt.Printf("Сложение: %d\n", addResult)
    fmt.Printf("Вычитание: %d\n", subResult)
    fmt.Printf("Умножение: %d\n", multResult)
}
