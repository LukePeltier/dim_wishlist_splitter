package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LukePeltier/dim_wishlist_splitter/pkg/parser"
)

func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}

func main() {
	voltronFile, err := os.Open("dim-wish-list-sources/voltron.txt")
	if err != nil {
		fmt.Printf("Error opening voltron: %v\n", err)
	}

	fileScanner := bufio.NewScanner(voltronFile)

	itemBlocks := make([]*parser.Block, 0)

	var currentBlock *parser.Block

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "//notes:") {
			currentBlock = parser.NewBlock()
			currentBlock.AddToNotes(line)
			itemBlocks = append(itemBlocks, currentBlock)
		} else if strings.HasPrefix(line, "//") {
			currentBlock = nil
		} else if strings.HasPrefix(line, "dimwishlist:") {
			if currentBlock == nil {
				continue
			}
			item := parser.NewItem(line)
			currentBlock.AddToItems(item)
		}
	}

	var count uint32 = 0

	fmt.Println("title:Filtered Voltron")
	for _, block := range itemBlocks {
		include := false

		if (!contains(block.Tags, "controller") || (contains(block.Tags, "m+kb") || contains(block.Tags, "mkb") || contains(block.Tags, "mnk"))) && (!contains(block.Tags, "pvp") || contains(block.Tags, "pve")) {
			include = true
		}
		if strings.Contains(strings.ToLower(block.Notes), "pvp") && !strings.Contains(strings.ToLower(block.Notes), "pve") {
			include = false
		}
		if include {
			count++
			fmt.Println(block.Notes)
			for _, item := range block.Items {
				fmt.Println(item.Code)
			}

			fmt.Println("")
		}
	}
	fmt.Println(count)
}
