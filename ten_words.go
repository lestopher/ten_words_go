package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strings"
    "sort"
)

const (
    DECLARATION_OF_INDEPENDENCE string = "data/declaration_of_independence.txt"
    CONSTITUTION                string = "data/constitution.txt"
)

var BLACKLIST = [...]string{"a", "an", "the", "of", "to", "and", "for", "our", "in", "has"}

//////////// WordCount /////////////
type WordCount struct {
    word string
    count int
}

type By func(wc1, wc2 *WordCount) bool

func (by By) Sort(wordCountSlice []WordCount) {
    wcs := &wordCountSorter{
        wordCountSlice: wordCountSlice,
        by: by,
    }
    sort.Sort(wcs)
}

type wordCountSorter struct {
    wordCountSlice []WordCount
    by func(wc1, wc2 *WordCount) bool
}

func (s *wordCountSorter) Len() int {
    return len(s.wordCountSlice)
}

func (s *wordCountSorter) Swap(i, j int) {
    s.wordCountSlice[i], s.wordCountSlice[j] = s.wordCountSlice[j], s.wordCountSlice[i]
}

func (s *wordCountSorter) Less(i, j int) bool {
    return s.by(&s.planets[i], &s.planets[j])
}
//************ WordCount *************//

// Gets rid of punctuation and numbers then lower cases everything
func sanitize(document string) string {
    regex := regexp.MustCompile("[,:;.&\n\r\"'()-0-9]")
    safe := regex.ReplaceAllString(document, " ")
    safe = strings.ToLower(safe)
    return safe
}

func count_words(words []string) map[string]int {
    wm := make(map[string]int)
    for _, value := range words {
        _, ok := wm[value]
        if ok {
            wm[value] += 1
        } else {
            wm[value] = 1
        }
    }
    return wm
}

func ordered_by_value(wm map[]) {

}

func main() {
    document, err := ioutil.ReadFile(CONSTITUTION)
    if err != nil {
        fmt.Println("Err is ", err)
    }
    // Makes a key-value of string:int
    wordMap := make(map[string]int)
    strBuffer := sanitize(string(document))
    words := strings.Fields(strBuffer)

    wordMap = count_words(words)


    for k, v := range wordMap {
        fmt.Printf("%s: %d\n", k, v)
    }
}
