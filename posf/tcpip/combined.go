package main

//
//func readFile(fileName string) []byte {
//	data, err := ioutil.ReadFile(fileName)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if len(data) == 0 {
//		fmt.Println("Exit (data = 0)")
//		return nil
//	}
//	return data
//}
//
//func dataUnmarshal(data []byte) []Combined {
//	d := []Combined{}
//	err := json.Unmarshal(data, &d)
//	if err != nil {
//		fmt.Println("error: ", err.Error())
//	}
//	return d
//}
