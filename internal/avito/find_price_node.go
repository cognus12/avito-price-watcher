package avito

import "strings"

func findPriceNode(content string) string {
	startIndex := strings.Index(content, "<span class=\"js-item-price\"")

	var node string

	node = content[startIndex:]

	closeIndex := strings.Index(node, "</span>")

	node = node[0 : closeIndex+7]

	return node
}
