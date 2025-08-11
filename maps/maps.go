package main

import "fmt"
func main(){
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Twitter": "https://www.twitter.com",	

}
fmt.Print(websites);
fmt.Print(websites["Google"]); // Accessing a value by key
// Adding a new key-value pair
websites["LinkedIn"] = "https://www.linkedin.com"
fmt.Print(websites);
// Deleting a key-value pair
delete(websites, "Twitter")
fmt.Print(websites);
// Checking if a key exists
if value, exists := websites["Facebook"]; exists {
	fmt.Println("Facebook exists:", value)
} else {
	fmt.Println("Facebook does not exist")			
}
// Iterating over the map
for key, value := range websites {
	fmt.Printf("%s: %s\n", key, value)
}
// Length of the map
fmt.Println("Number of websites:", len(websites))
// Clearing the map
websites = make(map[string]string) // Reinitialize the map
fmt.Println("Cleared map:", websites)
// Example of a nested map
	nestedMap := map[string]map[string]string{
		"Google": {
			"URL": "https://www.google.com",
			"Category": "Search Engine",
		},
		"Facebook": {
			"URL": "https://www.facebook.com",
			"Category": "Social Media",
		},
	}
	fmt.Println("Nested Map:", nestedMap)
	// Accessing nested map values
	if googleInfo, exists := nestedMap["Google"]; exists {
		fmt.Println("Google URL:", googleInfo["URL"])
		fmt.Println("Google Category:", googleInfo["Category"])
	} else {
		fmt.Println("Google does not exist in the nested map")
	}
	// Adding a new entry to the nested map
	nestedMap["Twitter"] = map[string]string{
		"URL":     "https://www.twitter.com",
		"Category": "Social Media",
	}	
}
// maps/maps.go	

