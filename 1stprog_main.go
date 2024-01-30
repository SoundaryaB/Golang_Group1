package main

import "fmt"

func main() {
	fmt.Println("EXERCISE 1")
	fmt.Println("********************************")
	mapNode := map[string][]string{"SR Family": {"7750", "7450", "7950", "7705hm"},
		"MPR Family": {"MPR-9500", "MPRe"}}

	// 1. Print family -> nodes
	fmt.Println("Map with Key as 'Family Name' and Values as 'Node Types' is displayed:")
	for keys, values := range mapNode {
		fmt.Printf("%s->%v\n", keys, values)
	}

	// 2. Add a new Family
	mapNode["PSS Family"] = []string{"1830-PSS", "1830-OCS"}
	fmt.Println("********************************")
	fmt.Println("After adding a new Family, the updated map is:")
	for keys, values := range mapNode {
		fmt.Printf("%s->%v\n", keys, values)
	}

	//3. Check if a family contains a node and add if it doesn't
	if !hasNode(mapNode["SR Family"], "7250") {
		mapNode["SR Family"] = append(mapNode["SR Family"], "7250")
	}

	fmt.Println("********************************")
	fmt.Println("After checking if value 7250 exists, the updated family is:")
	fmt.Printf("SR Family-> %v\n", mapNode["SR Family"])
	fmt.Println("********************************")

	// 4. Delete a deprecated node
	fmt.Println("After deleting 1830 PSS Family, the updated map is:")
	delete(mapNode, "PSS Family")
	for keys, values := range mapNode {
		fmt.Printf("%s->%v\n", keys, values)
	}
}

func hasNode(slice []string, value string) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}
