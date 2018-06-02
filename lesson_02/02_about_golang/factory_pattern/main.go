package main
import (
	"fmt"
	"github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang_factory/animal"
	"github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang_factory/food"
	"github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang_factory/factory"
)

func main(){
	// petType, foodType が引数で与えられたとします
	petType := animal.Dog{}
	foodType := food.Meat{}

	// factory の instance を取得します
	fa := factory.NewFactory(petType, foodType)

	// factory から Animal型 の instance を取得します
	a := fa.AnimalFactoryMethod()

	// factory から Food型 の instance を取得します
	f := fa.FoodFactoryMethod()

	// Animal の名前を取得します
	fmt.Println(a.GetName())

	// 鳴き声です
	fmt.Println(a.Say())

	// 今日の食べ物です
	fmt.Println(f.Get())
}
