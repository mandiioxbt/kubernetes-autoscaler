package main
import "fmt"
type Config struct { Min, Max int; TargetGPU float64 }
func main() {
    c := Config{Min: 1, Max: 16, TargetGPU: 70.0}
    fmt.Printf("Autoscaler: %d-%d replicas, target %.0f%% GPU\n", c.Min, c.Max, c.TargetGPU)
}
