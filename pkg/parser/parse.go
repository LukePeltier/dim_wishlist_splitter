package parser

import "strings"

var UsedTags map[string]bool = make(map[string]bool, 0)

type Item struct {
	Code string
}

func NewItem(code string) *Item {
	i := Item{Code: code}
	return &i
}

type Block struct {
	Items  []*Item
	Notes  string
	Tags   []string
	Author string
}

func NewBlock() *Block {
	b := Block{Items: make([]*Item, 0), Notes: "", Tags: make([]string, 0)}
	return &b
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func (b *Block) AddToNotes(noteLine string) {
	b.Notes = b.Notes + "\n" + noteLine
	extras := strings.SplitAfter(noteLine, "|")
	if len(extras) > 0 {
		for _, extra := range extras {
			if strings.HasPrefix(extra, "tags:") {
				tagsString, _ := strings.CutPrefix(extra, "tags:")
				tagsSlice := strings.FieldsFunc(tagsString, Split)
				for _, tagStr := range tagsSlice {
					cleanTagStr := strings.ToLower(strings.Trim(tagStr, " |"))
					b.AddToTags(cleanTagStr)
					UsedTags[cleanTagStr] = true
				}
			}
		}
	}
}

func (b *Block) AddToItems(item *Item) {
	b.Items = append(b.Items, item)
}

func (b *Block) AddToTags(tag string) {
	b.Tags = append(b.Tags, tag)
}

func (b *Block) SetAuthor(author string) {
	b.Author = author
}
