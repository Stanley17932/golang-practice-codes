package main

import "fmt"

type basket []produce


type produce struct {
	name string
	flavour string
	kind string

}


func (p *basket) add_item(entry produce){
	*p = append(*p, entry)


	fmt.Printf("Entry %s created!\n", entry.name)
}

func (p basket) change_item(name string, entry produce) {
		for key, val := range p {
				if val.name == name {
						p[key] = entry
				}
		}
		fmt.Printf("Item %s changed!\n", name)
}

func (p basket) items(){
	fmt.Printf("There are %d number of items in the basket\n",
			len(p))
	for _, val := range p {
		fmt.Printf(`Name: %s\n
		Flavour: %s\n
		Kind: %s\n`,
				val.name,
				val.flavour,
				val.kind)
		fmt.Println("")
	}
}

func main(){
	basket := new(basket)

	basket.add_item(
			produce{
				name:"apple",
				flavour: `It's a little sour and bitter, but mostly
				sweet, not at all salty, very juicy in general.`,
				kind: "fruit",
			},
	)
	basket.change_item("carrot",
			produce{
				name: "cucumber",
				flavour: `Slightly bitter with a mild melon aroma
				and planty flavor.`,
				kind: "veggies",
			},
		)

		basket.items()
}