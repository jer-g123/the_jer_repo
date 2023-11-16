package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{
		"var1: CoLorA",
		"var2: ColOrB",
		"var3: COlorc",
	}

	substr := []string{
		"Var1",
		"vAr2",
		"vaR3",
	}

	SortStr(str, substr)
}

func SortStr(str, substr []string) {
	originalCase := make(map[string]string)

	// Store the original case of substrings
	for _, sub := range substr {
		originalCase[strings.ToLower(sub)] = sub
	}

	for _, s := range str {
		s_low := strings.ToLower(s)

		// Initialize a variable to store the reconstructed string
		var reconstructed string

		for _, sub := range substr {
			sub_low := strings.ToLower(sub)
			if strings.Contains(s_low, sub_low) {
				parts := strings.SplitN(s_low, sub_low, 2)
				if len(parts) == 2 {
					reconstructed = fmt.Sprintf("%s%s%s", parts[0], originalCase[sub_low], parts[1])
					break
				}
			}
		}

		// Print the reconstructed string if it's not empty
		if reconstructed != "" {
			fmt.Println(reconstructed)
		} else {
			fmt.Printf("Could not find the tag: %v\n", s_low)
		}
	}
}



import statements: Import necessary packages (fmt and strings).
main function: Define the main function, create str and substr slices with example data, and call the SortStr function with these slices.
SortStr function declaration: Start the SortStr function. Initialize a map (originalCase) to store the original case of substrings.

Iterate over str and substr:
for _, s := range str: Loop through each string in str.
for _, sub := range substr: For each string in substr, convert it to lowercase (sub_low) and check if it's contained in the lowercase version of the current string (s). If it is, store the original case of the substring in the originalCase map.
/////This loop checks if the lowercase version of the current substring (sub_low) is contained in the lowercase version of the current string (s). If it is, it stores the original case of the substring (sub) in the originalCase map.
///////I appreciate your diligence, and I hope this clarifies the intention of that part of the code.

Continue with the SortStr function:
for _, s := range str: Loop through each string in str.
s_low := strings.ToLower(s): Convert the current string to lowercase and store it in s_low.
found := false: Initialize a flag (found) to check if any substring is found in the current string.

Nested loop over substr:
for _, sub := range substr: Loop through each substring in substr.
sub_low := strings.ToLower(sub): Convert the current substring to lowercase.
if strings.Contains(s_low, sub_low) {: Check if the lowercase substring is contained in the lowercase version of the current string.
parts := strings.SplitN(s_low, sub_low, 2): Split the lowercase string into two parts using the lowercase substring as the delimiter. The 2 in SplitN ensures that only the first occurrence is split.
if len(parts) == 2 {: Check if the split resulted in two parts.
rest := originalCase[sub_low] + parts[1]: Revert the second part to its original case by using the originalCase map.
fmt.Printf("%s%s%s\n", parts[0], sub_low, rest): Print the reconstructed string.
found = true: Set the found flag to true since a match was found.
break: Break out of the loop since a match was found.


Check if no substring was found in the current string:
if !found {: If the found flag is still false, print a message indicating that the tag was not found in the current string.
This code essentially iterates over each string in str, checks for the presence of substrings in a case-insensitive manner, and prints the reconstructed string with the original case of the substrings. If a substring is not found, it prints a message indicating that the tag was not found in the current string. The originalCase map is crucial for maintaining the original case of the substrings.



