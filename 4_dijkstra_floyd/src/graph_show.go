package main

import (
	"fmt"
	"github.com/emicklei/dot"
	// "github.com/wcharczuk/go-chart"
	"os"
	"os/exec"
	"strconv"
)

type edge struct {
	node  string
	label string
}
type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(from, to, label string) {
	g.nodes[from] = append(g.nodes[from], edge{node: to, label: label})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (e *edge) String() string {
	return fmt.Sprintf("%v", e.node)
}

func (g *graph) String() string {
	out := `digraph finite_state_machine {
		rankdir=LR;
		size="8,5"
		node [shape = circle];`
	for k := range g.nodes {
		for _, v := range g.getEdges(k) {
			out += fmt.Sprintf("\t%s -> %s\t[ label = \"%s\" ];\n", k, v.node, v.label)
		}
	}
	out += "}"
	return out
}

func toDotFile(graph [][]int, fileName string) {
	g := makeDotFormat(graph)

	visualFolderName := "visualisation"
	dotsFolderName := "dot"

	makeFolder(visualFolderName)
	makeFolder(visualFolderName + "/" + dotsFolderName)

	file, fileErr := os.Create("../" + visualFolderName + "/" + dotsFolderName + "/" + fileName)
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	fmt.Fprintf(file, "%v\n", g)
}

func makeFolder(folderName string) string {
	path := "/home/mira/Code/APA/4_dijkstra_floyd/"
	if err := os.Mkdir(path+folderName, os.ModePerm); err != nil && !os.IsExist(err) {
		// todo: catch err
	}
	return folderName
}

func openDot(name string) {
	dir1 := "/home/mira/Code/APA/4_dijkstra_floyd/visualisation/dot/" + name

	cmd := exec.Command("open", dir1)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", 556)
	}
}

func makeDotFormat(graph [][]int) *dot.Graph {
	g := dot.NewGraph(dot.Undirected)
	var F, T dot.Node
	var W string

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph); j++ {
			if graph[i][j] == 0 {
				continue
			}
			F = g.Node(strconv.Itoa(i)).Attr("color", "#84a598")
			T = g.Node(strconv.Itoa(j))
			W = strconv.Itoa(graph[i][j])
			if len(g.FindEdges(T, F)) > 0 {
				continue
			}
			g.Edge(F, T, W).Attr("color", "#d3869b")
		}
	}
	return g
}
