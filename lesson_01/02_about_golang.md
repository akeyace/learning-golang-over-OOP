# Golang 流オブジェクト指向
Golang （Go言語) とは何なのか？
そして、Golang 流のオブジェクト指向とは何なのか？
これから新しい世界を見ていきましょう。

## Golang とは
Google 社が開発したプログラミング言語
2012 年にバージョン 1.0 がリリースされた新しいプログラミング言語
処理速度も、コンパイル速度共に高速で、並列処理が得意という特徴があります
他のモダン言語と違い、開発のしやすさや、メンテナンスのしやすさにフォーカスされています
オブジェクト指向の流れを組んでいますが、既存のオブジェクト指向言語と大きな違いがあります

## Golang のオブジェクト指向
Golang のオブジェクト指向を理解する為、再度再度オブジェクト指向について考えてみましょう。
オブジェクト指向といっても、実は様々な方法論があります。

- クラスベース方式
  - 代表的なプログラミング言語: Java, C#, Python 等々
  - 特徴: Class や継承など『オブジェクト指向のおさらいとアンチパターン』で説明した特徴を持ち、最も一般的なオブジェクト指向と言えます
- プロトタイプベース方式
   - 代表的なプログラミング言語: Javascript(新しいバージョン ES6 以降はクラスベース方式も可能)
   - 特徴: プロトタイプベースでは、オブジェクトを Class よりも高い自由度で扱う事ができる。例えば、既に存在するオブジェクトに対して後付で変更を加えると言った事が可能。プラモデルで例えると、プラモデルの部品と設計書がある事は共通するが、柔らかい素材で出来ていて、部品を曲げて変形させるといった事が可能。設計書には部品の曲げ方について記述する事になる
- Mixin方式
  - 代表的なプログラミング言語: Python, C# 等
  - 特徴: Mixin とは、Class の継承や、Interface の実現に似た機能で、他のオブジェクトの特徴を副作用なく取り込む事が出来る機能の事。Python, C# 等では、クラスベース方式と混在して Mixin 方式を採用している。プラモデルでいうと、共通部品同士を取り込んで (Mix して！) 別の部品を作る事が出来る

### オブジェクト指向の本質
ここでオブジェクト指向の本質に立ち返ってみます。

『オブジェクト指向のおさらいとアンチパターン』の最初で書いた事を引用します。

> プログラミングのソースコードをオブジェクト(物体)とみなし、オブジェクト間のメッセージング(処理の呼び出し等)により記述する開発手法。

大事な事は、ソースコードを書きやすくする為の手法であり、オブジェクト間でメッセージをやり取りするという事です。

現在の大多数のクラスベース方式のオブジェクト指向言語は、開発効率向上の為、より厳密でバグが少なくするという方向で進化しているように思われます。
そのデメリットとして複雑化しているとも言えます。

### Golang  のオブジェクト指向
オブジェクト指向にも幾つもの方法があることがわかりました。
それでは、Golang のオブジェクト指向はどのような物でしょうか？

それは、Mixin方式をベースとしてクラスベース、プロトタイプベースを取り込んだ独自のオブジェクト指向と言えるものです。

独自と聞くと嫌な予感がする方も多いかと思います。
(独自方式というものは、大体大変な事態に陥りやすい)

では、どのような特徴があるのか見ていきましょう。

### Golang オブジェクト指向の特徴
- Class がない
  - Class がないので継承もない
- 代わりに Struct (構造体) と言われるオブジェクト利用する
  - Struct は自身のプロパティ、メソッドを持つ事ができる
- 継承がない代わりに Mixin によって他の Struct を再利用する事ができる
- Interface はある
  - Interface は Java のように厳密なものではなく、ダックタイピングと言われる、動的 (自動的) な判別により、ポリモーフィズムを行っている

では、具体的な例を見てみましょう
下記は Java の場合と Golang の場合で同じ結果になる処理です。

- Java の場合
    ```Java
    class Animal {
        private String name;
    
        public Animal(String name) {
            this.name = name;
        }
    
        public String getName() {
            return "I'm " + this.name;
        }
    }
    
    class Dog extends Animal {
        private String name;
        public Dog(String name) {
            super(name);
        }
    }
    
    public class Main {
        public static void main(String[] args) throws Exception {
            Animal animal = new Animal("dog");
    
            // 出力: I'm Dog
            System.out.println(animal.getName());
        
            Dog dog = new Dog("dog");
            
            // 出力: I'm Dog
            System.out.println(dog.getName());
        }
    }
    ```

- Golang の場合
    ```go
    package main
    import "fmt"
    
    type (
        Animal struct {
            name string
        }
        Dog struct {
            Animal
        }
    )
    
    func (a *Animal) GetName() string {
        return "I’m " + a.name
    }
    
    func main() {
        animal := Animal{name: "dog"}
    
        // 出力: I'm Dog
        fmt.Println(animal.GetName())
    
        dog := new(Dog)
        dog.name = "dog"
    
        // 出力: I'm Dog
        fmt.Println(dog.GetName())
    }
    ```

オブジェクト指向の方法に違いがあり、細かい所では書き方が異なりますが、同じ様な書き方ができます。

### Golang 流オブジェクト指向のメリット
では、この違いの何が嬉しいのでしょうか？

- Mixin が使えるため、サブクラス地獄というものが発生しづらい
  - Java では上手く設計しないと、孫、ひ孫、玄孫、、とそれぞれが密接な関係 (密結合という。反対は疎結合といい、疎結合の方が問題が発生しづらい) になり、修正が難しくなることがあります。
  - Mixin の利用により、移譲というオブジェクト指向の書き方がし易い

コンポーネント単位と言うものがあり、大体のオブジェクト指向言語では、package と Class が最小単位になりますが、Golang の場合は package のみがコンポーネント単位になります。

これにより、ドメイン駆動開発がし易いという特徴があります。

今は解らなくても心配無用です。
これから勉強会の中で少しずつ説明していきます！

### Golang 流オブジェクト指向のデメリット
大多数のオブジェクト指向では、ポリモーフィズムを行うため、ジェネリクス (総称) といわれる機能が備わっています。

ジェネリクスとは、『型』が異なっていても同じような扱いを行う事が可能になる機能ですが、Golang にはありません。
その為、ジェネリクスのあるオブジェクト指向と同じような設計が出来ず、書き方を Golang 流に合わせる必要があります。

また、Golang は関数型プログラミングが可能な言語ですが、ジェネリクスがないため、関数型プログラミングを行う際に若干の不便があります。

### Golang の基本構文
それでは、Golang の基本構文を見ていきましょう。

- package の指定と、ライブラリの読み込み
    ```go
    // package は最上段に書きます。 その時にコメントは無視されます。
    package main
    
    // import 文により、他のライブラリを読み込みます。
    import “fmt”
    ```
- struct の書き方
    ```go
     // type 文によって Struct を定義します
     // Animal と頭文字を大文字で書いた場合は public にります
     type Animal struct {
         name string
     }
     
     // dog と頭文字を小文字で書いた場合は、private となり、他の package から触る事ができなくなります (カプセル化)
     type dog struct {
         Animal
     }
     
     // type 文は () で囲む事によって一纏めにすることが出来ます
     type (
         Animal struct {
             name string
         }
         Dog struct {
             Animal
         }
     )
    ```
- メソッドの書き方
    ```go
    // Golang の特殊な書き方で、(a *Animal) の部分の右側をレシーバと言い、このメソッドを所有する Struct を指定しています
    // GetName と頭文字が大文字の場合は public になり、小文字で getName とすると private になります
    // string の部分は戻り値の型が書かれます。この場合は string (文字列型) を返しています
    func (a *Animal) GetName() string {
        // return の後に戻り値を記述します
        // “” で囲んだ文字は文字列型になります
        // a.name は、レシーバの左側にある文字 a は、レシーバの参照 (instance のようなもの。他の言語でいう this 相当) を表しています
        return "I’m " + a.name
    }
    ```
- global な関数
    ```go
    // レシーバの記述が無い場合は、この Package 内で利用可能な global な関数として扱われます
    // GetName のあとの() に書かれた 『name string』は、この関数内で利用する name という変数が string 型であることを表します
    func GetName(name string) string {
        return "I’m " + name
    }
    
    func main() {
        name := GetName("dog")
    
        fmt.Println(name)
    }
    ```
- 処理の書き方
    ```go
    // main() はプログラムが実行された際に最初に呼ばれる関数になります
    func main() {
        // Struct Animal の変数を代入しています
        // animal は、変数相当の物になります
        // 『:=』は、型推論を行う時の書き方です
        // 型推論をしない場合 var animal Animal = Animal(name: “dog”) となります
        // Animal{} は、Struct から変数を得る時の書き方になります
        // name: “dog” は、Animal のプロパティ name に dog という文字列を代入しています
        animal := Animal{name: "dog"}
    
        // fmt.Println は、文字列を標準出力に表示する方法
        // animal の GetName メソッドを実行しています
        fmt.Println(animal.GetName())
    
        // new により Struct Dog のポインタを代入しています
        dog := new(Dog)
        // dog にMixin された Animal の name プロパティに文字列 dog を代入しています
        dog.name = "dog"
    
         // dog に MIxin された Animal の GetName メソッドを実行しています
         fmt.Println(dog.GetName())
    ```

これが Golang の基本的な構文になります。

### Golang その他特徴
Golang はシンプルさを追求した言語で、言語仕様が少なく、比較的早くに習得することが可能です。
とはいえ、他のオブジェクト指向言語との違いから、Java 等になれた技術者は最初は上手く扱えない所もあります。

また、Golang の思想として、言語仕様をシンプルにする事で、言語自体のアップデートは少なめです。
その分、各種ライブラリを充足させていこうという理念があるようです。
これは、言語仕様自体をライブラリの追加によりカスタマイズ可能であるという思想の現れでもあります。
その為、ライブラリ自体を開発しやすく、誰でも簡単に公開出来る仕組みも言語仕様として組み込まれています。
(この辺りは node.js 等とは好対照と言えるでしょう) 

### Golang まとめ
Golang の世界の入り口如何でしたでしょうか？

Golang  のオブジェクト指向は最初はややこしいと思われる所もあります。
しかし、Golang のオブジェクト指向は、より本質的なオブジェクト指向に立ち返っているとも言えます。

Golang を触りたくて仕方がないのではないでしょうか？
その辺りの思想/開発のコツをハンズオンや、その他学習を通してこれから学んで参ります。
