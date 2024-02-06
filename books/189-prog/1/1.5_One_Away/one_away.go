package main

import (
	"fmt"
)

func checkInsertOrRemove(str1 string, str2 string) bool {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1;
	}
	if (len(str1) - len(str2)) > 1 {
		return false;
	}

	insertSeen := false;
	var j int = 0;
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[j] {
			if insertSeen {
				return false;
			}
			insertSeen = true;
		} else {
			j ++;
			if j >= len(str2) {
				break;
			}
		}
	}
	return true;
}

func checkReplace(str1 string, str2 string) bool {
	replaceSeen := false;
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if replaceSeen {
				return false;
			}
			replaceSeen = true;
		}
	}
	return true;
}

func main() {
	str1 := string("");
	str2 := string("");
	fmt.Println("Enter the first string:")
	fmt.Scan(&str1);
	fmt.Println("Now please enter the second:");
	fmt.Scan(&str2);

	if str1 == str2 {
		fmt.Printf("true\n");
	} else if len(str1) == len(str2) {
		if checkReplace(str1, str2) {
			fmt.Printf("true\n");
		} else {
			fmt.Printf("false\n"); 
		}
	} else {
		if checkInsertOrRemove(str1, str2) {
			fmt.Printf("true\n");
		} else {
			fmt.Printf("false\n"); 
		}
	}
}
