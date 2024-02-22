package utils


import (
	"fmt"
)


type Graph struct {
	Adj map[int][]int;
}

func (g * Graph) dfsInternal(v int, visited map[int]bool) {
	visited[v] = true;
	fmt.Printf("%d  ", v);
	for _, neighbor := range g.Adj[v] {
		if !visited[neighbor] {
			g.dfsInternal(neighbor, visited);
		}
	}
}

func (g *Graph) DFS(v int) {
	visited := make(map[int]bool);
	for vertex, _ := range g.Adj {
		visited[vertex] = false;
	}
	
	g.dfsInternal(v, visited);
	fmt.Println();
}

func (g *Graph) BFS(v int) {
	visited := make(map[int]bool);
	
	visited[v] = true;
	queue := &Queue{};
	queue.PushBack(v);

	for ; !queue.Empty(); {
		s := queue.Front();

		fmt.Printf("%d  ", s);

		queue.PopFront();

		for _, neighbor := range g.Adj[s] {
			if !visited[neighbor] {
				visited[neighbor] = true;
				queue.PushBack(neighbor);
			}
		}
	}
	fmt.Println();


}

func (g *Graph) AddEdge(v int, w int) {
	g.Adj[v] = append(g.Adj[v], w);
}