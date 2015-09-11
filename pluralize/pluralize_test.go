package pluralize_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/stringfy/pluralize"
	"github.com/stretchr/testify/assert"
)

func Test_Plural(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		count     int
		singular  string
		addPlural bool
		plural    string
		out       string
	}{
		{
			count:    1,
			singular: "ball",
			out:      "1 ball",
		},
		{
			count:     2,
			singular:  "ball",
			addPlural: true,
			plural:    "ballys",
			out:       "2 ballys",
		},
		{
			count:    2,
			singular: "boat",
			out:      "2 boats",
		},
		{
			count:    2,
			singular: "axis",
			out:      "2 axes",
		},
		{
			count:    2,
			singular: "Testis",
			out:      "2 testes",
		},
		{
			count:    2,
			singular: "octopus",
			out:      "2 octopi",
		},
		{
			count:    2,
			singular: "virus",
			out:      "2 viri",
		},
		{
			count:    2,
			singular: "viri",
			out:      "2 viri",
		},
		{
			count:    2,
			singular: "octopi",
			out:      "2 octopi",
		},
		{
			count:    2,
			singular: "bus",
			out:      "2 buses",
		},
		{
			count:    2,
			singular: "buffalo",
			out:      "2 buffaloes",
		},
		{
			count:    2,
			singular: "tomato",
			out:      "2 tomatoes",
		},
		{
			count:    2,
			singular: "zoocytium",
			out:      "2 zoocytia",
		},
		{
			count:    2,
			singular: "zoocytia",
			out:      "2 zoocytia",
		},
		{
			count:    2,
			singular: "analysis",
			out:      "2 analyses",
		},
		{
			count:    2,
			singular: "wolf",
			out:      "2 wolves",
		},
		{
			count:    2,
			singular: "thief",
			out:      "2 thieves",
		},
		{
			count:    2,
			singular: "hive",
			out:      "2 hives",
		},
		{
			count:    2,
			singular: "quality",
			out:      "2 qualities",
		},
		{
			count:    2,
			singular: "fox",
			out:      "2 foxes",
		},
		{
			count:    2,
			singular: "dash",
			out:      "2 dashes",
		},
		{
			count:    2,
			singular: "pass",
			out:      "2 passes",
		},
		{
			count:    2,
			singular: "vouch",
			out:      "2 vouches",
		},
		{
			count:    2,
			singular: "index",
			out:      "2 indices",
		},
		{
			count:    2,
			singular: "matrix",
			out:      "2 matrices",
		},
		{
			count:    2,
			singular: "quiz",
			out:      "2 quizzes",
		},
		{
			count:    2,
			singular: "ox",
			out:      "2 oxen",
		},
		{
			count:    2,
			singular: "oxen",
			out:      "2 oxen",
		},
		{
			count:    2,
			singular: "person",
			out:      "2 people",
		},
		{
			count:    2,
			singular: "man",
			out:      "2 men",
		},
		{
			count:    2,
			singular: "child",
			out:      "2 children",
		},
		{
			count:    2,
			singular: "sex",
			out:      "2 sexes",
		},
		{
			count:    2,
			singular: "move",
			out:      "2 moves",
		},
		{
			count:    2,
			singular: "zombie",
			out:      "2 zombies",
		},
		{
			count:    2,
			singular: "equipment",
			out:      "2 equipment",
		},
		{
			count:    2,
			singular: "information",
			out:      "2 information",
		},
		{
			count:    2,
			singular: "rice",
			out:      "2 rice",
		},
		{
			count:    2,
			singular: "money",
			out:      "2 money",
		},
		{
			count:    2,
			singular: "species",
			out:      "2 species",
		},
		{
			count:    2,
			singular: "series",
			out:      "2 series",
		},
		{
			count:    2,
			singular: "fish",
			out:      "2 fish",
		},
		{
			count:    2,
			singular: "sheep",
			out:      "2 sheep",
		},
		{
			count:    2,
			singular: "jeans",
			out:      "2 jeans",
		},
		{
			count:    2,
			singular: "police",
			out:      "2 police",
		},
		{
			count:    2,
			singular: "flower",
			out:      "2 flowers",
		},
	}

	for _, t := range tests {
		p := pluralize.New()
		if t.addPlural {
			p.Options(pluralize.AddPlural(t.plural))
		}
		pf := p.Perform(t.count, t.singular)
		a.Equal(t.out, pf)
	}
}

func ExampleNew() {
	p1 := pluralize.New()
	fmt.Println(p1.Perform(2, "octopus"))

	p2 := pluralize.New()
	fmt.Println(p2.Perform(2, "post"))

	p3 := pluralize.New()
	fmt.Println(p3.Perform(2, "jeans"))

	//Custom plural
	pc := pluralize.New()
	pc.Options(pluralize.AddPlural("boatys"))
	fmt.Println(pc.Perform(2, "boat"))
	// Output:
	// 2 octopi
	// 2 posts
	// 2 jeans
	// 2 boatys
}
