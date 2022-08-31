package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	line := "слиток тяпка листок пятак столик пятка я в поо"
	lLine := strings.ToLower(line)
	sLine := strings.Split(lLine, " ")
	anagrams, err := MakeMap(sLine)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(anagrams)
}

// MakeMap ...makes anagrams
func MakeMap(s []string) (map[string][]string, error) {
	parse := make(map[string][]string)
	for _, v := range s {
		word := []rune(v)
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		sWord := string(word)
		parse[sWord] = append(parse[sWord], v)
	}

	resMap := make(map[string][]string)
	for _, v := range parse {
		if len(v) > 1 {
			resMap[v[0]] = v
			sort.Strings(v)
		}
	}
	fmt.Println(resMap)
	return resMap, nil
}
