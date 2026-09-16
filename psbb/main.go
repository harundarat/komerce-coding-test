package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculateMinBuses calculates the minimum number of buses required based on PSBB rules.
func calculateMinBuses(totalFamilies int, memberCountsStr []string) (totalBuses int, err error) {
	if len(memberCountsStr) != totalFamilies {
		return 0, fmt.Errorf("Input must be equal with count of family")
	}

	// familySizeCounts stores the frequency of families based on their member size (index 1-4).
	familySizeCounts := make([]int, 5)
	for _, memberCountStr := range memberCountsStr {
		memberCount, err := strconv.Atoi(memberCountStr)
		if err != nil || memberCount < 1 || memberCount > 4 {
			return 0, fmt.Errorf("Invalid family member count")
		}
		familySizeCounts[memberCount]++

	}

	// 1. Families with 4 members (1 full bus)
	totalBuses += familySizeCounts[4]

	// 2. Families with 3 members (max combined with 1 family of 1 member)
	totalBuses += familySizeCounts[3]
	familySizeCounts[1] -= familySizeCounts[3]
	if familySizeCounts[1] < 0 {
		familySizeCounts[1] = 0
	}

	// 3. Families with 2 members paired together (max 2 families per bus)
	totalBuses += familySizeCounts[2] / 2

	// If there is 1 remaining family of 2 members, allocate 1 bus for them.
	// The remaining seats can only be filled by max 1 family of 1 member.
	if familySizeCounts[2]%2 != 0 {
		totalBuses++

		familySizeCounts[1]--
		if familySizeCounts[1] < 0 {
			familySizeCounts[1] = 0
		}
	}

	// 4. Remaining families with 1 member (max combined 2 families per bus)
	if familySizeCounts[1] > 0 {
		totalBuses += (familySizeCounts[1] + 1) / 2
	}

	return totalBuses, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input the number of families: ")
	inputTotalFamilies, _ := reader.ReadString('\n')
	totalFamilies, err := strconv.Atoi(strings.TrimSpace(inputTotalFamilies))
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Print("Input the number of members in the family (separated by a space): ")
	inputMemberCounts, _ := reader.ReadString('\n')
	memberCountsStr := strings.Fields(inputMemberCounts)

	totalBuses, err := calculateMinBuses(totalFamilies, memberCountsStr)
	if err != nil {
		if err.Error() == "Input must be equal with count of family" {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Error: ", err)
		}

		return
	}

	fmt.Printf("Minimum bus required is: %d\n", totalBuses)
}
