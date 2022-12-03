package main

import "fmt"

func scambioV1(str1, str2 string) { //versione 1
	str1, str2 = str2, str1
}

func scambioV2(p_str1, p_str2 *string) { //versione 2
	*p_str1, *p_str2 = *p_str2, *p_str1
}

func main() {
	s1 := "ciao"
	s2 := "hello"

	fmt.Println(s1, s2)

	scambioV1(s1, s2)                        //versione 1
	fmt.Println("scambio V1:", s1, s2)

	scambioV2(&s1, &s2)                  //versione 2
	fmt.Println("scambio V2:", s1, s2)
}
