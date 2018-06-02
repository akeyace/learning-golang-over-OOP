# Golang のデザインパターン
デザインパターンに強い効力がある事がわかりました。
Golang のプログラミングにも使いたい所です。
ところがどっこい、Golang では幾つかのデザインパターンが言語仕様として取り込まれています。

では、Golang ではどう書くのでしょうか？
実際に書いていきましょう！

## Golang で Factory パターン
Golang では Factory パターンが言語仕様に取り込まれ、随所で書いていく事になります。
どのような言語仕様なのか見ていきましょう

### Golang で Factory パターン未実装の場合
下記のコードを IDE で実際に書いてみましょう！
- main.go
    ```go
    package main
    import (
        "fmt"
        // 頭にドットを付けると package 名を省略出来る
        . "github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/animal"
        . "github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/food"
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
    ```
- animal/animal.go
    ```go
    package animal
    
    type (
        InterfaceAnimal interface {
            GetName() string
            Say() string
        }
    
        Animal struct {
            Name string
        }
    
        Dog struct {
            Animal
        }
    
        Cat struct {
            Animal
        }
    
        Monster struct {
            Animal
        }
    )
    
    func (a Animal) GetName() string {
        return "I'm " + a.Name
    }
    
    func (d Dog) Say() string {
        return "ワン"
    }
    
    func (c Cat) Say() string {
        return "ニャン"
    }
    
    func (m Monster) Say() string {
        return "ぎゃー"
    }
    ```
- food/food.go
    ```go
    package food
    
    // strings converter ライブラリ
    import "strconv"
    
    type (
        InterfaceFood interface {
            Get() string
        }
    
        Food struct {
            Volume int
        }
    
        Meat struct {
            Food
        }
    
        Fish struct {
            Food
        }
    
        Mud struct {
            Food
        }
    )
    
    func (m Meat) Get() string {
  	    // strconv.Itoa で  int型 から string型に変換する
        return "this is a meat! volume: " + strconv.Itoa(m.Volume) + "g"
    }
    
    func (f Fish) Get() string {
        return "this is a fish! volume: " + strconv.Itoa(f.Volume) + "g"
    }
    
    func (m Mud) Get() string {
        return "this is a mud.... volume: " + strconv.Itoa(m.Volume) + "g"
    }
    ```

上記コードは GitHub に上がっています。
[Golang Factory パターン未実装コード - GitHub](https://github.com/akeyace/learning-golang-over-OOP/tree/lesson_2_1-before-factory-pattern/lesson_02/02_about_golang/factory_pattern)

### Golang で Factory パターン実装の場合
では、Factory パターンを適用してみましょう。
Golang では次のように書きます。

- animal/animal.go
    ```diff
          Monster struct {
              Animal
          }
      )
    + 
    + func NewDog() Dog {
    +     return Dog{Animal{"dog"}}
    + }
    + 
    + func NewCat() Cat {
    +     return Cat{Animal{"cat"}}
    + }
    + 
    + func NewMonster() Monster {
    +     return Monster{Animal{"monster"}}
    + }
    ```
- food/food.go
    - import 追加
        ```diff
        - import "strconv"
        + import (
        +     "strconv"
        +     "github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/factory_pattern/animal"
        + )
        ```
    - Factory メソッド追加
        ```diff Mud struct { Food
              }
          )
        
        + func NewMeat(a animal.InterfaceAnimal) Meat {
        +     switch a.(type) {
        +     case animal.Dog:
        +         return Meat{Food{100}}
        +     case animal.Cat:
        +         return Meat{Food{40}}
        +     default:
        +         panic("error")
        +     }
        + }
        + 
        + func NewFish(a animal.InterfaceAnimal) Fish {
        +     switch a.(type) {
        +     case animal.Dog:
        +         return Fish{Food{100}}
        +     case animal.Cat:
        +         return Fish{Food{40}}
        +     default:
        +         panic("error")
        +     }
        + }
        ```

Golang の場合、Factory Method パターンまでで十分で、より複雑な場合に Abstract Factory パターンを使う事になります。
ですが、ここは先程の Java と合わせて Abstract Factory パターンを実装してみましょう。

- factory/factory.go
    ```go
    package factory
    
    import (
        "github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/factory_pattern/animal"
        "github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/factory_pattern/food"
    )
    
    type Factory struct {
        a animal.InterfaceAnimal
        f food.InterfaceFood
    }
    
    func (f Factory) AnimalFactoryMethod() animal.InterfaceAnimal {
        return f.a
    }
    
    func (f Factory) FoodFactoryMethod() food.InterfaceFood {
        return f.f
    }
    
    func NewFactory(a animal.InterfaceAnimal, f food.InterfaceFood) Factory {
        var ai animal.InterfaceAnimal
        var fi food.InterfaceFood
    
        switch a.(type) {
        case animal.Dog:
            ai = animal.NewDog()
        case animal.Cat:
            ai = animal.NewCat()
        default:
            panic("error")
        }
    
        switch f.(type) {
        case food.Meat:
            fi = food.NewMeat(ai)
        case food.Fish:
            fi = food.NewFish(ai)
        default:
            panic("error")
        }
        return Factory{ai, fi}
    }
    ```    

Factory パターンを適用したら、main.go を修正しましょう。

- main.go
    ```go
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
    ```
 
恐ろしいほどスッキリしました。

修正して動作を確認してみましょう。

[Golang Factory パターン実装コード - GitHub](https://github.com/akeyace/learning-golang-over-OOP/tree/lesson_2_1-after-factory-pattern/lesson_02/02_about_golang/factory_pattern)

## Golang の Factory パターン
- NewXXX はコンストラクタ関数と言われ、その名の通り他の言語のコンストラクタに該当します
- しかし、コンストラクタでありながら実態は関数なので Factory Method パターンのような動作をします
- 複雑な処理の場合は Abstract Factory パターンを使い呼び出すコンストラクタ関数を変更させることができます
- まさにナチュラルボーンデザインパターン！

## その他 Golang のデザインパターン
- Template Method パターン
    - 上記コードで InterFaceAnimal#GetName() をオーバーライドするかのような動きは Template Method パターンを実装していると言えます。
    - Golnag の Interface による実現はダックタイピングのため、自然と Template Method パターンになります
- Strategy パターン
    - 上記コードで Mixin による移譲と、Interface によるポリモーフィズムは Strategy パターンを実装していると言えます。
    - 特に Golang では継承がないため、DRY 原則に則ると自然と Strategy パターンになります
- Singleton パターン
    - 言語仕様として関数の外に var Utils = &Utils{} と初期化したり、init 関数を作る事で簡単に Singleton を作成可能
    - Golang は並列処理が前提の言語なので、そのあたりは扱いやすく作られています
- Iterator パターン
    - Golang では言語仕様には取り込まれておらず、関数のコールバック等により実装することが可能です。

## Golang のデザインパターンまとめ
Golang のオブジェクト指向の扱いやすさは、デザインパターンを言語仕様に取り込んでいる辺りにあるかと思われます。
特にデザインパターンを意識しなくても、自然にデザインパターンになるという面白さがあります。
Golang を上手く使うことは、難解なデザインパターンの習得の近道ではないかと思われます。

他に 「これデザインパターンが言語仕様で実装されてる！」と気づいた方は是非お知らせください！

Golang 流オブジェクト指向に慣れてデザインパターンをモノにしましょう！
