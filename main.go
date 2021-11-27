package main

import (
    "fmt"
)

type Strategy interface {
    ExecuteStrategy()
}

type StrategyA struct {
    Name string
}

func (s *StrategyA) ExecuteStrategy() {
    fmt.Printf("%s実行\n", s.Name)
}

func NewStrategyA(name string) Strategy {
    return &StrategyA{
        Name: name,
    }
}

type StrategyB struct {
    Name string
}

func (s *StrategyB) ExecuteStrategy() {
    fmt.Printf("%s実行\n", s.Name)
}

func NewStrategyB(name string) Strategy {
    return &StrategyB{
        Name: name,
    }
}

type Player struct {
    Strategy Strategy
}

func main() {
    sa := NewStrategyA("StrategyA")
    p1 := &Player{
        Strategy: sa,
    }

    p1.Strategy.ExecuteStrategy()

    sb := NewStrategyB("StrategyB")
    p2 := &Player{
        Strategy: sb,
    }

    p2.Strategy.ExecuteStrategy()
}

