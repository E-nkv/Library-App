package main

import "fmt"

var allowedTagNames = []string{
	"Action",
	"Adventure",
	"Comedy",
	"Drama",
	"Fantasy",
	"Historical Fiction",
	"Horror",
	"Mystery",
	"Romance",
	"Science Fiction",
	"Thriller",
	"Young Adult",
	"Dystopian",
	"Non-Fiction",
	"Biography",
	"Self-Help",
	"Cookbook",
	"Graphic Novel",
	"Poetry",
	"Classic Literature",
	"Children's Literature",
	"Literary Fiction",
	"Suspense",
	"Urban Fantasy",
	"Magical Realism",
	"Crime Fiction",
	"Family Saga",
	"Western",
	"Sports Fiction",
	"Travel Literature",
	"Historical Romance",
	"Cyberpunk",
	"Steampunk",
	"Paranormal Romance",
	"New Adult",
	"Short Stories",
	"Anthology",
	"Memoir",
	"Essays",
	"Philosophy",
	"Religion & Spirituality",
	"Health & Wellness",
	"Politics & Current Events",
	"Science & Nature",
	"True Crime",
	"Humor",
	"Satire",
	"Fairy Tales & Folklore",
	"Coming-of-Age",
	"Epic Fantasy",
	"Space Opera",
	"Alternate History",
	"Military Fiction",
	"Psychological Thriller",
	"Gothic Fiction",
	"Chick Lit",
	"Romantic Comedy",
	"Dark Fantasy",
	"Cozy Mystery",
	"Historical Mystery",
	"Medical Thriller",
	"Legal Thriller",
	"Political Thriller",
	"Animal Fiction",
	"Inspirational Fiction",
	"Women’s Fiction",
	"LGBTQ+ Fiction",
	"Futuristic Fiction",
	"Survival Fiction",
	"Post-Apocalyptic",
	"Time Travel",
	"Space Exploration",
	"Cyber Crime",
	"Teen Fiction",
	"Supernatural",
	"Adventure Romance",
	"Art & Photography",
	"Business & Economics",
	"Education",
	"Parenting",
	"Nature Writing",
	"Environmental Fiction",
	"Social Issues",
	"Classic Mystery",
	"Historical Adventure",
	"Fantasy Adventure",
	"Science Fantasy",
	"Sci-Fi Romance",
	"Dramedy",
	"Sword and Sorcery",
	"Narrative Non-Fiction",
	"Cultural Studies",
	"Graphic Memoir",
	"Travelogue",
	"Tales of the Unexpected",
	"Sagas & Epics",
}
var totalAllowedTags = len(allowedTagNames)

func (s S) seedTags(N int) error {
	if N > totalAllowedTags {
		return fmt.Errorf("invalid amount of tags: got %d, and the max is %d", N, totalAllowedTags)
	}
	_, err := s.DB.Exec("DELETE FROM tags")
	if err != nil {
		return err
	}

	smtp, err := s.DB.Prepare("INSERT INTO tags VALUES(DEFAULT, $1)")
	if err != nil {
		return err
	}
	for _, tagName := range allowedTagNames {
		if _, err := smtp.Exec(tagName); err != nil {
			return err
		}
	}
	return nil
}
