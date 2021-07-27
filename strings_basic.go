package main
import "fmt"
import "unicode"
import "strings"
import "strconv"


func test_strconv(){
    i, _ := strconv.Atoi("42")
    fmt.Println("Result strconv.Atoi %d\n", i)
    
    i2 := 10
	s := strconv.Itoa(i2)
	fmt.Printf("%T, %v\n", s, s)
    
    // FormatFloat(f float64, fmt byte, prec, bitSize int) -> string
    v := 3.1415926535
	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
    
    fs := "3.1415926535"
    fn32, _ := strconv.ParseFloat(fs, 32)
    fn64, _ := strconv.ParseFloat(fs, 64)
    fmt.Printf("%T, %v\n", fn32, fn32)
    fmt.Printf("%T, %v\n", fn64, fn64)
    
    val := "3.1415926535"
    if s, err := strconv.ParseFloat(val, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat(val, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}

func test_string_manipulation(){
    //params to replace: s, old, new string, n int
    fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) 
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
    
    fmt.Println("Result strings.ToLower", strings.ToLower("WOW"))
    fmt.Println("Result strings.ToUpper", strings.ToUpper("pow"))
    fmt.Println("Result strings.TrimSpace", strings.TrimSpace("  whitespaces-removed   "))
    
}


func test_compare() {
    compare_1 := strings.Compare("a", "a")
    compare_2 := strings.Compare("a", "b")
    compare_3 := strings.Compare("b", "a")
    
    fmt.Printf("Result Compared_1 %d\n", compare_1)
    fmt.Printf("Result Compared_2 %d\n", compare_2)
    fmt.Printf("Result Compared_3 %d\n", compare_3)
    
    fmt.Printf("Result Contains %t\n", strings.Contains("hello", "o"))
    fmt.Printf("Result HasPrefix %t\n", strings.HasPrefix("hello", "hel"))
    fmt.Printf("Result Index %d\n", strings.Index("asdfgh", "d"))
}


func test_unicode() {
    digit := rune('1')
    space := rune(' ')
    letter := rune('a')
    lower := rune('G')
    punct := rune('?')
    
    is_digit := unicode.IsDigit(digit)
    is_space := unicode.IsSpace(space)
    is_letter := unicode.IsLetter(letter)
    is_lower := unicode.IsLower(lower)
    is_punct := unicode.IsPunct(punct)
    
    fmt.Printf("Result is_digit %t\n", is_digit)
    fmt.Printf("Result is_space %t\n", is_space) 
    fmt.Printf("Result is_letter %t\n", is_letter) 
    fmt.Printf("Result is_lower %t\n", is_lower) 
    fmt.Printf("Result is_punct %t\n", is_punct)
    
    upper_letter := unicode.ToUpper(letter)
    fmt.Printf("Result ToUpper %#U\n", upper_letter)
    
    lower_letter := unicode.ToLower(lower)
    fmt.Printf("Result ToLower %#U\n", lower_letter)
    
}

func main() {
    test_unicode()
    test_compare()
    test_string_manipulation()
    test_strconv()
}
