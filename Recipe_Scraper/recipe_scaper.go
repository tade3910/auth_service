package recipescraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type RecipeScraper struct {
}

func (scraper *RecipeScraper) Scrape(url string) {
	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch the URL: %s, Status Code: %d", url, resp.StatusCode)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	scraper.debugNode(doc, 0)
	fmt.Println("----Ingredients found----")
	allIngredients := parseIngredients(doc)
	for i, ingredients := range allIngredients {
		fmt.Printf("----List %d----\n", i+1)
		for index, ingredient := range ingredients {
			fmt.Printf("%d. %s\n", index+1, ingredient)
		}
		fmt.Println("---------------")
	}
}

func getListNode(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && (n.Data == "ul" || n.Data == "ol") {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childRes := getListNode(c)
		if childRes != nil {
			return childRes
		}
	}
	return nil
}

func getIngredient(n *html.Node) string {
	ingredient := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			ingredient += c.Data
		}
		ingredient += getIngredient(c)
	}
	return strings.TrimSpace(ingredient)
}

func getIngredients(n *html.Node) []string {
	ingredients := []string{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "li" {
			ingredient := getIngredient(c)
			if len(ingredient) > 0 {
				ingredients = append(ingredients, ingredient)
			}
		}
	}
	if len(ingredients) == 0 {
		return nil
	}
	return ingredients
}

func getIngredientListNodes(n *html.Node, ingredientListNodes *[]*html.Node) bool {
	if n.Type == html.TextNode && strings.Compare(strings.ToLower(n.Data), "ingredients") == 0 {
		log.Print("Found potential ingredients")
		return true
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if getIngredientListNodes(c, ingredientListNodes) {
			// check if c contains list
			listNode := getListNode(c)
			if listNode != nil {
				*ingredientListNodes = append(*ingredientListNodes, listNode)
				return false
			}
			return true
		}
	}
	return false
}

func parseIngredients(n *html.Node) [][]string {
	ingredientListNodes := []*html.Node{}
	getIngredientListNodes(n, &ingredientListNodes)
	if len(ingredientListNodes) == 0 {
		log.Fatal("Could not find ingredients")
	}
	// For now I'm only going to handle the first one
	ingredientLists := [][]string{}
	for _, node := range ingredientListNodes {
		ingredientLists = append(ingredientLists, getIngredients(node))
	}
	return ingredientLists
}

// Recursive function to print HTML nodes with indentation for better readability
func (scraper *RecipeScraper) debugNode(n *html.Node, indentLevel int) {
	// Add indentation for pretty printing
	indent := fmt.Sprintf("%*s", indentLevel*2, "")

	// Print the opening tag and its attributes (if any)
	if n.Type == html.ElementNode {
		fmt.Printf("%s<%s", indent, n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}
		fmt.Println(">") // Opening tag
	}

	// Print text node content (if any)
	if n.Type == html.TextNode {
		fmt.Printf("%s%s\n", indent, n.Data)
	}

	// Recursively print child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		scraper.debugNode(c, indentLevel+1) // Increase indentation for nested elements
	}

	// Print the closing tag (if it is an element)
	if n.Type == html.ElementNode {
		fmt.Printf("%s</%s>\n", indent, n.Data)
	}
}
