package html

import (
	"regexp"
	"strings"

	"github.com/matDobek/gov--attendance-check/internal/predicates"
	"golang.org/x/net/html"
)

type element struct {
	tag   string
	class []string
	id    []string
}

func Extract(doc string, query string) ([]string, error) {
	nodes, err := Search(doc, query)
	if err != nil {
		return nil, err
	}

	result := []string{}
	for _, node := range nodes {
		children := childrenOf(node)
		childrenResult := doExtract(children)
		var filteredChildrenResult []string

		for _, child := range childrenResult {
			child = strings.Trim(child, " \n\t")
			if child == "" {
				continue
			}

			filteredChildrenResult = append(filteredChildrenResult, child)
		}

		result = append(result, strings.Join(filteredChildrenResult, " "))
	}

	return result, nil
}

func doExtract(nodes []*html.Node) []string {
	var result []string

	for _, node := range nodes {
		if node.FirstChild == nil {
			result = append(result, node.Data)
			continue
		}

		children := childrenOf(node)
		childrenResult := doExtract(children)
		result = append(result, childrenResult...)
	}

	return result
}

func childrenOf(parent *html.Node) []*html.Node {
	children := []*html.Node{}

	child := parent.FirstChild
	for {
		children = append(children, child)

		child = child.NextSibling

		if child == nil {
			break
		}
	}

	return children
}

func ExtractAttr(doc string, query string, attr string) ([]string, error) {
	nodes, err := Search(doc, query)
	if err != nil {
		return nil, err
	}

	result := []string{}
	for _, node := range nodes {
		for _, a := range node.Attr {
			if a.Key == attr {
				result = append(result, a.Val)
			}
		}
	}

	return result, nil
}

func Search(doc string, query string) ([]*html.Node, error) {
	reader := strings.NewReader(string(doc))
	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	elems := toElements(query)
	result := doSearch([]*html.Node{root}, elems)

	return result, nil
}

func toElements(query string) []element {
	var result []element

	query = strings.Trim(query, " \n\t")
	for _, s := range strings.Split(query, " ") {
		tag := regexp.MustCompile("^[a-zA-Z-_]+").FindString(s)
		id := regexp.MustCompile("\\#[a-zA-Z0-9-_]+").FindAllString(s, -1)
		class := regexp.MustCompile("\\.[a-zA-Z0-9-_]+").FindAllString(s, -1)

		// remove the leading '#' or '.'
		for i, v := range id {
			id[i] = v[1:]
		}
		for i, v := range class {
			class[i] = v[1:]
		}

		result = append(result, element{tag: tag, id: id, class: class})
	}

	return result
}

func doSearch(roots []*html.Node, queries []element) []*html.Node {
	nodes := []*html.Node{}

	for _, root := range roots {
		found := findNodes(root, queries[0])

		// do not search children, if that's the last element from the query ( return outermost element )
		// e.g for
		//		body
		//			div.first.container
		//				div.second.container
		//	query: "body .container"
		//	will return only {"div.first.container"}
		if len(queries) == 1 {
			nodes = append(nodes, found...)
			continue
		}

		var children []*html.Node
		for _, n := range found {
			if n.FirstChild == nil {
				continue
			}

			children = append(children, n.FirstChild)
		}

		foundFromChildren := doSearch(children, queries[1:])
		nodes = append(nodes, foundFromChildren...)
	}

	return nodes
}

func findNodes(node *html.Node, query element) []*html.Node {
	var found []*html.Node

	if node == nil {
		return []*html.Node{}
	}

	for {
		if isMatching(node, query) {
			found = append(found, node)
		}

		foundFromChildren := findNodes(node.FirstChild, query)
		found = append(found, foundFromChildren...)

		if node.NextSibling == nil {
			break
		}

		node = node.NextSibling
	}

	return found
}

func isMatching(n *html.Node, query element) bool {
	attrs := make(map[string]string)
	for _, a := range n.Attr {
		attrs[a.Key] = a.Val
	}

	result := true

	if len(query.tag) > 0 {
		result = result && (n.Data == query.tag)
	}

	if len(query.id) > 0 {
		ids := strings.Split(attrs["id"], " ")

		for _, v := range query.id {
			if predicates.Contains(ids, v) {
				continue
			}

			result = false
			break
		}
	}

	if len(query.class) > 0 {
		classes := strings.Split(attrs["class"], " ")

		for _, v := range query.class {
			if predicates.Contains(classes, v) {
				continue
			}

			result = false
			break
		}
	}

	return result
}
