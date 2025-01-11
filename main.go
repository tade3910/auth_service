package main

import recipescraper "github.com/tade3910/recipe_parser/Recipe_Scraper"

func main() {
	scaper := &recipescraper.RecipeScraper{}
	// scaper.Scrape("https://www.budgetbytes.com/creamy-garlic-chicken/")
	scaper.Scrape("https://www.allrecipes.com/recipe/269500/creamy-garlic-pasta/")
	// scaper.Scrape("https://thecozycook.com/easy-lasagna-recipe/")
}
