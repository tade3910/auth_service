package main

import recipescraper "github.com/tade3910/recipe_parser/Recipe_Scraper"

func main() {
	scaper := &recipescraper.RecipeScraper{}
	scaper.Scrape("https://www.budgetbytes.com/creamy-garlic-chicken/")
	scaper.Scrape("https://www.allrecipes.com/recipe/269500/creamy-garlic-pasta/")
	scaper.Scrape("https://thecozycook.com/easy-lasagna-recipe/")
	scaper.Scrape("https://www.allrecipes.com/chef-johns-jollof-rice-recipe-7499757")
	scaper.Scrape("https://sweetandsavorymeals.com/jerk-chicken-recipe/")
	scaper.Scrape("https://simplehomeedit.com/recipe/baked-garlic-chicken-breast/")

}
