package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "sort"
    "strings"
)

const (
    DECLARATION_OF_INDEPENDENCE string = "data/declaration_of_independence.txt"
    CONSTITUTION                string = "data/constitution.txt"
)

var BLACKLIST map[string]struct{} = map[string]struct{}{
    "a":   struct{}{},
    "an":  struct{}{},
    "the": struct{}{},
    "of":  struct{}{},
    "to":  struct{}{},
    "and": struct{}{},
    "for": struct{}{},
    "our": struct{}{},
    "in":  struct{}{},
    "has": struct{}{},
}

//////////// WordCount /////////////
type Word struct {
    value string
    count int
}

type WordCount struct {
    words       []Word
    wordTracker map[string]int
    length      int
    by          func(w1, w2 *Word) bool
}

/**
 * Returns a bool value and the slice index
 * @param word string
 * @return (bool, int)
 */
func (wc *WordCount) exists(word string) (bool, int) {
    index, ok := wc.wordTracker[word]
    if ok {
        return true, index
    }
    return false, -1
}

type By func(wc1, wc2 *Word) bool

func (by By) Sort(words []Word) {
    wcs := &WordCount{
        words: words,
        by:    by,
    }
    sort.Sort(wcs)
}

func (s *WordCount) Len() int {
    return len(s.words)
}

func (s *WordCount) Swap(i, j int) {
    s.words[i], s.words[j] = s.words[j], s.words[i]
}

func (s *WordCount) Less(i, j int) bool {
    return s.by(&s.words[i], &s.words[j])
}

//************ WordCount *************//

// Gets rid of punctuation and numbers then lower cases everything
func sanitize(document string) string {
    regex := regexp.MustCompile("[,:;.&\n\r\"'()-0-9]")
    safe := regex.ReplaceAllString(document, " ")
    safe = strings.ToLower(safe)
    return safe
}

func count_words(words []string) *WordCount {
    wc := new(WordCount)
    if wc.words == nil {
        wc.words = make([]Word, 1)
        wc.wordTracker = make(map[string]int)
    }
    for _, value := range words {
        exists, index := wc.exists(value)
        if exists {
            wc.words[index].count += 1
        } else {
            _, blacklisted := BLACKLIST[value]
            if !blacklisted && len(value) > 2 {
                word := Word{value, 1}
                wc.words = append(wc.words, word)
                wc.wordTracker[value] = len(wc.words) - 1
            }
        }
    }
    return wc
}

func main() {
    document, err := ioutil.ReadFile(CONSTITUTION)
    if err != nil {
        fmt.Println("Err is ", err)
    }
    strBuffer := sanitize(string(document))
    words := strings.Fields(strBuffer)
    wc := count_words(words)

    for _, v := range wc.words {
        fmt.Printf("%s: %d\n", v.value, v.count)
    }
}
