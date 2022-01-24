package main

import (
	"fmt"
	"html/template"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	//exVariable()
	//multiVariable()
	//DataType()
	//exOperator()
	//exCondition()
	//forLoop()
	//exArray()
	//exSlice()
	//exArray()
	//exMap()
	//exFunction()
	//listHobbies("eat", "play", "sleep")
	//exClosure()
	//exParam()
	//exPointer()
	//exStruct()
	//exInterface()
	//exReflect
	//exGoroutine()
	//exStringFormat()
	//exTime()
	//exWebServer()
	//exWebServer2()
	exURLParsing()
}

func exVariable() {
	var firstName string = "Arif"
	var middleName string
	middleName = "Random"
	lastName := "Lesmana"

	// _ Untuk menampung value yang tidak terpakai
	animal, _ := "bear", 21

	// new is use for pointer variable (?). the default value will following the data type.
	// and the result of the pointer variable is a address with hexadecimal code.
	// if u wanna know the real values, must be deference using asterisk (*)
	pointer := new(string)

	fmt.Println(firstName, middleName, lastName)
	fmt.Printf("Hello %s %s ! \n", firstName, lastName)
	fmt.Printf("This is %s here \n", animal)
	fmt.Println(pointer)
	fmt.Println(*pointer)
}

func multiVariable() {
	// 1st way
	var satu, dua, tiga int
	satu, dua, tiga = 1, 2, 3

	//2nd way
	//var satu, dua, tiga int = 1,2,3

	//3d way. tapi := hanya bisa digunakan dalam function
	/*satu, dua, tiga := 1,2,3
	//bisa digunakan untuk beberapa variabel dengan tipe data yang berbeda2
	one, dia, mungkin := 1, "Kamu", true */

	fmt.Println(satu, dua, tiga)
}

func DataType() {
	var (
		a     uint = 100
		b     int  = -2011
		c          = 2.14
		exist bool = true
	)
	const phi = 3.14
	fmt.Printf("uint use for positive number only %d \n", a)
	fmt.Printf("int use for positive and negative number %d \n", b)
	fmt.Printf("percentage f use for get float value %f\n", c)
	fmt.Printf("to get bool value use t %t\n", exist)
	fmt.Printf("Phi is valid : %.2f \n", phi) // .2 dua angka dibelakang koma
}

func exOperator() {
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Printf("Is it true the result is 2 ? %t \n", isEqual)

	var (
		left  = true
		right = false
	)
	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var rightReverse = !right
	fmt.Printf("!right = (%t)\n", rightReverse)
}

func exCondition() {
	point := rand.Int()

	if point == 10 {
		fmt.Printf("Nilai anda sempurna : %d", point)
	} else if point < 10 && point >= 5 {
		fmt.Printf("Anda lulus dengan nilai %d", point)
	} else {
		fmt.Printf("Anda tidak lulus dengan nilai %d", point)
	}

	switch point {
	case 1, 2, 3, 4, 5:
		fmt.Println("Failed")

	case 6:
		fmt.Println("Pass")

	case 10:
		fmt.Println("Perfect")

	default:
		fmt.Println("Take the pass first !")

	}

	switch {
	//penggunaan case dengan syarat (if-else)
	case point == 10:
		fmt.Println("Youre perfect")
	case point < 10 && point >= 5:
		fmt.Println("Youre passed")
	default:
		fmt.Println("Try again")
	}
}

func forLoop() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
}

func exArray() {
	var animals [3]string
	//var animals[...]string -> panjangnya akan menyesuaikan isi yang ada pada array
	animals[0] = "Shark"
	animals[1] = "Dog"
	animals[2] = "Cat"

	fmt.Println(animals[0], animals[1], animals[2])
	fmt.Println("Array range", len(animals))
	fmt.Println("List of animal : ", animals)

	//multidimension
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//array with for
	for i := 0; i < len(animals); i++ {
		fmt.Printf("%d animal is : %s \n", i, animals[i])
	}

	//array with range (?)
	for _, animal := range animals {
		fmt.Printf("nama buah : %s\n", animal)
	}

	//deklarasi sekaligus panjang data dengan tipe data
	var fruits = make([]string, 2)
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fmt.Println(fruits)
	fmt.Println(fruits[0:1])
}

func exSlice() {
	//Slice little bit same with array. slice doesnt need declare the range of the array
	var fruits = []string{"Apple", "Banana", "Grape", "Honeydew"}
	var (
		aFruits  = fruits[0:3]
		bFruits  = fruits[1:4]
		aaFruits = aFruits[0:1]
		bbFruits = bFruits[1:2]
	)
	var cFruits = append(bbFruits, "Strawberry")

	fmt.Println(len(fruits)) // 4
	fmt.Println(cap(fruits)) // 4
	fmt.Println(fruits[0:1]) // [apple]
	fmt.Println(aFruits)     // [apple Banana Grape]
	fmt.Println(bFruits)     // [Banana Grape Honeydew]
	fmt.Println(aaFruits)    // [Apple]
	fmt.Println(bbFruits)    // [Grape]
	fmt.Println(cFruits)     // [Grape Strawberry]

	// Example using copy
	a := make([]string, 3)
	b := []string{"Watermelon", "Pineapple", "Apple", "Mango"}
	//copy will insert their value to a
	n := copy(a, b)

	fmt.Println(a) // [Watermelon Pineapple Apple]
	fmt.Println(b) // [Watermelon Pineapple Apple Mango]
	fmt.Println(n) // 3

	x := []int{1, 3, 4, 5, 12}
	y := []int{2, 8, 5, 9}
	cp := copy(x, y)
	fmt.Println(x)  // [2 8 5 9 12]
	fmt.Println(y)  // [2 8 5 9]
	fmt.Println(cp) // 4

	//Slicing with 3 index. Third index is use for the capacity of the slice
	var mFruits = fruits[0:2:3]
	fmt.Println(mFruits)
	fmt.Println(len(mFruits))
	fmt.Println(cap(mFruits))

}

func exMap() {
	// Map is same as dictionary in python
	numbers := map[string]int{}
	// Map can be declare with another ways
	// 1. var numbers = make(map[string]int)
	// 2. var number = *new(map[string]int)

	numbers["satu"] = 1
	numbers["dua"] = 2
	fmt.Println("satu", numbers["satu"])
	fmt.Println("dua", numbers["dua"])
	fmt.Println(numbers)

	//set map value fwhen made it
	months := map[string]string{"1st Month": "January", "2nd Month": "February"}
	//can be set in vertical way too

	/*months := map[string]string{
	"1st Month": "January",
	"2nd Month": "February"}*/

	fmt.Println(months)

	for key, val := range months {
		fmt.Println(key, " \t:", val)
	}

	//delete a key and a value in map
	delete(numbers, "dua")
	fmt.Println(numbers)

	//Checking item with a key
	var value, isExist = months["1st Month"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item not exist")
	}

	var chikens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	// key can be different each other
	// var data = []map[string]string{
	// 	{"name": "chicken blue", "gender": "male", "color": "brown"},
	// 	{"address": "mangga street", "id": "k001"},
	// 	{"community": "chicken lovers"},
	// }

	for _, chiken := range chikens {
		fmt.Println(chiken)
	}
}

func ExFunction() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	//Sprintf is use for convert the output into string
	var avg = calculation(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata: %.3f /n", avg)
	fmt.Printf(msg)

	randomValue = randomWithRange(2, 100)
	fmt.Printf("Random number : %d \n", randomValue)

	devideNumber(10, 2)
	devideNumber(4, 0)
	devideNumber(10, 5)

	var diameter = 15
	var area, keliling = calculate(float64(diameter))

	fmt.Printf("area of circle : \t %f \n", area)
	fmt.Printf("circumference of circle : \t %f \n", keliling)

	fmt.Sprintf("Nilai rata2nya adalah %f", avg)

}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func devideNumber(f, s int) {
	if s == 0 {
		fmt.Println("Invalid devider %d cannot devide by %d", f, s)
		return
	}

	result := f / s
	fmt.Printf("%d / %d = %d \n", f, s, result)
}

func calculate(d float64) (area float64, keliling float64) {
	area = math.Pi * math.Pow(d/2, 2)
	keliling = math.Pi * d
	return area, keliling
}

// ... is Variadic. use for unlimited parameter
func calculation(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func listHobbies(name string, hobbies ...string) {
	var hobbyAsString = strings.Join(hobbies, ",")

	fmt.Printf("Hello my name is %s \n", name)
	fmt.Printf("My hobbies are: %s \n", hobbyAsString)

}

func exClosure() {
	//MinMax case
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}

		}
		return min, max
	}
	var numbers = []int{1, 2, 3, 4, 5, 6, 67, 1, 7, 8, 10}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v \n min : %v \n max : %v \n", numbers, min, max)

	// Immediately-Invoked Function Expression (IIFE) Closure
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(5)
	fmt.Println("List of number : ", numbers)
	fmt.Println("List of filtered number : ", newNumbers)

}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func exParam() {
	var (
		data          = []string{"john", "michael", "hanabi"}
		dataContainsO = filter(data, func(each string) bool {
			return strings.Contains(each, "a")
		})
		dataLenght5 = filter(data, func(each string) bool {
			return len(each) == 6
		})
	)
	fmt.Printf("Main data : %v \n", data)
	fmt.Printf("Filtered data, data contains A : %v \n", dataContainsO)
	fmt.Printf("filter data lenght : %v \n", dataLenght5)
}

func exPointer() {
	// using *(arterisk) for pointer data type.
	// we can get pointer value from a variable with & (amperstand). This method called refencing
	// and we also can get true pointer value with *. this method called deferencing
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)     : ", numberA)
	fmt.Println("numberA (address)   : ", &numberA)
	fmt.Println("numberB (value)     : ", *numberB)
	fmt.Println("numberA (address)   : ", numberB)
}

// Make a struct with a type. we must definite the key first with the type also
type students struct {
	name string
	age  int
}

type class struct {
	students
	major string
	grade int
	age   int
}

func exStruct() {
	// var s1 students
	// s1.name = "Arif Lesmana"
	// s1.grade = 10

	// There os third way to using the struct
	//1st way
	var s1 = students{}
	s1.name = "Arif Lesmana"
	s1.age = 10

	var s2 = students{"David", 20}
	var s3 = students{name: "Teguh"}

	//is okay to flip the value. it's okay
	var s4 = students{name: "Jason", age: 13}
	var s5 = students{name: "Michael", age: 100}

	var s6 = class{}
	s6.name = "Johan"
	s6.age = 23
	s6.major = "Science"
	s6.grade = 3

	var s7 = class{}
	s6.name = "Johan"
	s6.age = 23
	s6.major = "Science"
	s6.grade = 3
	s6.students.age = 15

	//Combine struct with slice

	var allStudents = []students{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	fmt.Printf("My name is %s \t", s1.name)
	fmt.Printf("Age : %d \n", s1.age)

	fmt.Printf("My name is %s \t", s2.name)
	fmt.Printf("Age : %d \n", s2.age)

	fmt.Println("My name is \t", s3.name)
	fmt.Println("Age : \n", s3.age)

	fmt.Printf("My name is %s \t", s4.name)
	fmt.Printf("Age : %d \n", s4.age)

	fmt.Printf("My name is %s \t", s5.name)
	fmt.Printf("Age : %d \n", s5.age)

	fmt.Printf("My name is %s \t", s5.name)
	fmt.Printf("Age : %d \n", s5.age)

	fmt.Println("My name is \t", s6.name)
	fmt.Println("Age : \n", s6.age)
	fmt.Println("Major : \t", s6.major)
	fmt.Println("Grade : \n", s6.grade)

	fmt.Println("My name is \t", s7.name)
	fmt.Println("Age : \n", s7.age)
	fmt.Println("Major : \t", s7.major)
	fmt.Println("Grade : \n", s7.grade)
	fmt.Println("Major Age : \n", s7.students.age)

	for _, student := range allStudents {
		fmt.Printf("My name : %s \t My age : %v", student.name, student.age)
	}
}

type passangers struct {
	name         string
	from         string
	destination  string
	chair_number int
}

func (p passangers) sayHello() {
	fmt.Println("Hello, my name is %s", p.name)
}

func (p passangers) getNameAt(i int) string {
	return strings.Split(p.name, "")[i-1]
}

func exMethod() {
	var s1 = passangers{"Anton", "Jakarta", "Singapore", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("Nama panggilan : ", name)
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Phi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type berhitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	panjang float64
}

func (k kubus) luas() float64 {
	return math.Pow(k.panjang, 2) * 6
}

func (k kubus) keliling() float64 {
	return k.panjang * 12
}

func (k kubus) volume() float64 {
	return math.Pow(k.panjang, 3)
}

func exInterface() {
	var bangunDatar hitung
	var bangunRuang berhitung

	bangunDatar = persegi{10.0}
	fmt.Println("Luas persegi : ", bangunDatar.luas())
	fmt.Println("Keliling persegi : ", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Printf("Luas lingkaran : %.2f \n", bangunDatar.luas())
	fmt.Println("Keliling lingkaran : ", bangunDatar.keliling())
	fmt.Println("Jari-jari lingkaran : ", bangunDatar.(lingkaran).jariJari())

	bangunRuang = kubus{7.0}
	fmt.Println("Keliling kubus : ", bangunRuang.keliling())
	fmt.Println("Luas kubus : ", bangunRuang.luas())
	fmt.Println("Volume kubus : ", bangunRuang.volume())

}

func exReflect() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func exGoroutine() {
	runtime.GOMAXPROCS(2)

	go print(5, "Apa kabar")
	print(5, "apa kabar ")

	var input string
	fmt.Scanln(&input)

}

func exStringFormat() {
	type student struct {
		name        string
		height      float64
		age         int32
		isGraduated bool
		hobbies     []string
	}

	var data = student{
		name:        "wick",
		height:      182.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}

	fmt.Printf("%b\n", data.age) // %b Mengubah data numerik menjadi string biner

	fmt.Printf("%c\n", 1234) // %c mengubah data numerik menjadi kode unicode string
	fmt.Printf("%c\n", 12)

	fmt.Printf("%d\n", data.age) // %d mengubah data numerik menjadi bentuk string

	fmt.Printf("%e\n", data.height) // %e atau %E memformat data desimal menjadi numerik standart Scientific notation
	fmt.Printf("%E\n", data.height)

	fmt.Printf("%.9f\n", data.height) // %f mengambil data desimal dan bisa menentukan lebarnya
	// 182.500000000
	fmt.Printf("%.2f\n", data.height)
	// 182.50
	fmt.Printf("%.f\n", data.height)
	// 182

	fmt.Printf("%e\n", 0.123123123123) // %g sama dengan %f namun dengan %g lebarnya tidak bisa ditentukan
	// 1.231231e-01
	fmt.Printf("%f\n", 0.123123123123)
	// 0.123123
	fmt.Printf("%g\n", 0.123123123123)
	// 0.123123123123

	fmt.Printf("%o\n", data.age) // %o memformat data numerik menjadi string numerik berbasis 8 oktal

	fmt.Printf("%p\n", &data.name) //%p memformat nilai pointer dituliskan dengan numberik berbasis 16 dengan prefix 0x
	// 0x2081be0c0

	fmt.Printf("%q\n", `" name \ height "`) //%q digunakan untuk escape string
	// "\" name \\ height \""

	fmt.Printf("%s\n", data.name) //%s mengambil nilai string
	// wick

	fmt.Printf("%t\n", data.isGraduated) //%t mengambil tipe data
	// false
	fmt.Printf("%T\n", data.name)
	// string
	fmt.Printf("%T\n", data.height)
	// float64
	fmt.Printf("%T\n", data.age)
	// int32
	fmt.Printf("%T\n", data.isGraduated)
	// bool
	fmt.Printf("%T\n", data.hobbies)
	// []string

	fmt.Printf("%v\n", data) //%v mengembalikan string nilai aslinya
	// {wick 182.5 26 false [eating sleeping]}

	fmt.Printf("%#v\n", data) //%#v untuk format struct, mengembalikan nama dan property sesuai dengan struktur struct
	// main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:[]string{"eating", "sleeping"}}

	fmt.Printf("%x\n", data.age) //%x mengubah data numerik menjadi numerik hexadesimal
	// 1a

	var d = data.name

	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	// 7769636b
	fmt.Printf("%x\n", d)
	// 7769636b

	fmt.Printf("%%\n") //%% untuk menampilkan %
	// %
}

func exTime() {
	time1 := time.Now()
	fmt.Printf("Time 1 : %v \n", time1)

	time2 := time.Date(2022, 05, 21, 12, 34, 10, 0, time.UTC)
	fmt.Printf("Date Time : %v \n", time2)

}

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "apa kabar!")
// }

// func exWebServer() {
// 	//http.HandleFunc() Berfungsi untuk routing aplikasi web. yang artinya menentukan aksi ketika url tsb diakses oleh user
// 	// Ada 2 param yang harus diisi. yang pertama rute yang diinginkan (/) dan yang kedua adalah aksi saat rute itu diakses
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Halo!")
// 	})
// 	http.HandleFunc("/index", index)
// 	//http.listenAndServe adigunakan untuk menghidupkan server sekaligus menjalankan aplikasi menggunakan server tsb.

// 	fmt.Println("Starting web server at http://localhost:8080/")
// 	http.ListenAndServe(":8080", nil)
// }

func exWebServer2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name": "Arif LEsmana",
			"Msg":  "Have a niceday",
		}
		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func exURLParsing() {
	var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("url:%s\n", urlString)

	fmt.Printf("Protocol : %s\n", u.Scheme)
	fmt.Printf("host : %s\n", u.Host)
	fmt.Printf("Path : %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("Name: %s, Age: %s\n", name, age)

}
