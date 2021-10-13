package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

func main() {

	data := map[string]float64{"a": 0.1, "e": 0.5, "f": 0.6, "d": 0.4, "b": 0.2, "c": 0.3}

	fmt.Println(setData(data))

}

func setData(data map[string]float64) string {
	type temp struct {
		k string
		v float64
	}
	td := make([]temp, 0, len(data))
	for key, val := range data {
		t := temp{key, val}
		td = append(td, t)
	}

	// сортировка по убыванию
	sort.SliceStable(td, func(i, j int) bool {
		return td[i].v*100 > td[j].v*100
	})

	mas := make([]string, 0, len(data))
	for _, t := range td {
		d := fmt.Sprintf("['%s',%2.2f]", t.k, t.v)
		mas = append(mas, d)
	}

	js, _ := json.Marshal(mas)
	js = bytes.ReplaceAll(js, []byte(`"`), []byte(``))

	return "[['Комбинация', 'Вероятность']," + string(js)[1:]
}
