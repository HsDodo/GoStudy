package html

type Node struct {
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
)

type Attribute struct {
	Key, Val string
}
