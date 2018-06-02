package factory_pattern
import (
	"fmt"
	// 頭にドットを付けると package 名を省略出来る
	. "github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/factory_pattern/animal"
	. "github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/factory_pattern/food"
)

func main(){
	// petType, foodType が引数で与えられたとします
	petType := "cat"
	foodType := "fish"

	// a(animal) f(food) の変数を初期化
	// golang 文化では変数名を短くしがち (package 名との重複等の兼ね合い？)
	var a InterfaceAnimal
	var f InterfaceFood

	// pwrType によって Animal,　Food の instane を取得します
	switch petType {
	case "dog":
		// petType == "dog" の場合、Dogクラス の instance を取得
		a = Dog{Animal: Animal{Name: "dog"}}

		// petType == "dog" の場合の foodType ごとの適量な Foodクラス の instance を取得します
		switch foodType {
		case "meat":
			// 100g のお肉
			f = Meat{Food: Food{Volume: 100}}
		case "fish":
			// 40g の魚
			f = Fish{Food: Food{Volume: 40}}
		default:
			// 適切な Case がない場合、なんと 300g の泥……。
			f = Mud{Food: Food{Volume: 300}}
		}
	case "cat":
		// petType == "cat" の場合、Catクラス の instance を取得
		a = Cat{Animal: Animal{Name: "cat"}}

		// petType == "cat" の場合の foodType ごとの適量な Foodクラス の instance を取得します
		switch foodType {
		case "meat":
			// 20g のお肉
			f = Meat{Food: Food{Volume: 20}}
		case "fish":
			// 80g の魚
			f = Fish{Food: Food{Volume: 80}}
		default:
			// 適切な Case がない場合、なんと 100g の泥……。
			f = Mud{Food: Food{Volume: 100}}
		}
	default:
		// 適切な Case がない場合、なんとモンスターが……
		a = Monster{Animal: Animal{Name: "monster"}}

		switch foodType {
		case "meat":
			f = Meat{Food: Food{Volume: 0}}
		case "fish":
			f = Fish{Food: Food{Volume: 0}}
		default:
			f = Mud{Food: Food{Volume: 1000}}
		}
	}

	// Animal の名前を取得します
	fmt.Println(a.GetName())

	// 鳴き声です
	fmt.Println(a.Say())

	// 今日の食べ物です
	fmt.Println(f.Get())
}
