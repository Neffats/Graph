package main

import (
	"fmt"

	"github.com/Neffats/graph"
)

func main() {
	g := graph.NewNetwork()
	fillGraph(g)
	fmt.Printf("%s", g)
}

func fillGraph(g *graph.Network) {
    nA := graph.NewNode("A")
    nB := graph.NewNode("B")
    nC := graph.NewNode("C")
    nD := graph.NewNode("D")
    nE := graph.NewNode("E")
    nF := graph.NewNode("F")
    g.AddNode(nA)
    g.AddNode(nB)
    g.AddNode(nC)
    g.AddNode(nD)
    g.AddNode(nE)
    g.AddNode(nF)

    g.AddEdge(nA, nB)
    g.AddEdge(nA, nC)
    g.AddEdge(nB, nE)
    g.AddEdge(nC, nE)
    g.AddEdge(nE, nF)
    g.AddEdge(nD, nA)
}