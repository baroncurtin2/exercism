package tree

import "fmt"

const rootId int = 0

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	positions := buildPositions(records)
	root := buildNodes(positions, records)

	return root, nil

}

func buildPositions(records []Record) []int {
	positions := make([]int, len(records))

	for i, r := range records {
		if r.ID < rootId || r.ID >= len(records) {
			panic(fmt.Sprintf("out of bounds record id %d", r.ID))
		}
	}

	return positions
}

func buildNodes(positions []int, records []Record) *Node {
	recordLen := len(records)
	nodes := make([]Node, recordLen)

	for i := range positions {
		r := records[positions[i]]

		if r.ID != i {
			panic(fmt.Sprintf("non-contiguous node %d, (want %d", r.ID, i))
		}

		validParentForChild := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootId)

		if !validParentForChild {
			panic(fmt.Sprintf("node %d has impossible parent %d", r.ID, r.Parent))
		}

		nodes[i].ID = i

		if i != rootID {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}

	return &nodes[0]
}

