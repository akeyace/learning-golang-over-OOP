# デザインパターン ファクトリーパターンを学ぼう
デザインパターンは中々覚えるのが大変です。
しかし、手を動かして実践してみると意外に理解は進みます。
そこで今回は実際に手を動かしてファクトリーパターンを見てましょう！

環境構築が大変なのでオンラインエディタを使って学習を勧めていきます。

こちらのサイトを使いましょう。
[Paiza.io](https://paiza.io/)

## Factory パターン概要
Factory パターンは厳密には、Factory Method パターン、Abstract Factory パターンがあります。
それら2つを纏めて Factory パターンと言われる事もあります。

それぞれの内容を確認したあと、実装を見ていきましょう

### Factory Method パターン概要
- 目的
    - オブジェクトを生成する時のインタフェースだけを想定して、実際にどのクラスをインスタンス化するかはサブクラスが決めるようにする
    Factory Method パターンはインスタンス化をサブクラスに任せる
- 別名
    - Virtual Constructor
- 動機(意訳)
    - 構築するシステムが大きくなると、それにつれ沢山の部品(クラスと、その子孫クラス達)を作る事になる
    すると、正しくオブジェクトのインスタンスを得るという所にバグが生まれやすくなる
- 適用可能性
    - クラスが、生成しなければならないオブジェクトのクラスを事前に知ることが出来ない場合
    - サブクラス化により、生成するオブジェクトを特定化する場合
    - クラスが責任を幾つかのサブクラスの中の1つに委譲するときに、どのサブクラスに委譲するのかに関する知識を局所化したい場合
- クラス図
    ![factory method by wikipedia](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/Factory_Method_UML_class_diagram.svg/1000px-Factory_Method_UML_class_diagram.svg.png)
        - Wikipedia より
- 構成要素
    - Product クラス
    - ConcreteProduct クラス
    - Creator クラス
    - ConcreteCreator クラス
- 結果
   - サブクラスに手がかりを提供する
   - パラレルなクラス階層をつなぐ
   
## Abstract Factory パターン概要
- 目的
    - 互いに関連したり依存し合うオブジェクト群を、その具象クラスを明確にせずに生成するためのインタフェースを提供する
- 別名
    - Kit
- 動機(意訳)
    - 例えば外観を変えるテーマがある時に、テーマを変える事でそれに伴い生成されるオブジェクトのセットが切り替わることがある
    その時に適切なオブジェクトのインスタンス化を行いたい
- 適用可能性
    - システムを部品の生成、組み合わせ、表現の方法から独立にすべき場合
    - 部品の集合が複数存在して、その中の1つを選んでシステムを構築する場合
    - 一群の関連する部品を常に使用しなければならないように設計する場合
    - 部品のクラスライブラリを提供する際に、インタフェースだけを公開して、実装は非公開にしたい場合
- クラス図
    ![abstract factory by wikipedia](https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Abstract_Factory_UML_class_diagram.svg/1300px-Abstract_Factory_UML_class_diagram.svg.png)
- 構成要素
    - AbstractFactory クラス
    - ConcreteFactory クラス
    - AbstractProduct クラス
    - ConcreteProduct クラス
    - Client クラス
- 結果
   - 具象クラスを局所化する
   - 部品の集合を用意に変更できるようになる
   - 部品間の無矛盾性を促進する
   - (副作用) 新たな種類の部品に対応することが困難である

## Factory パターンを実装する
百聞は一見にしかず。
それでは実際に実装してみましょう！

まずは、Factory パターンを適用しない場合のソースコードを見てみましょう。

実装は、解りやすいよう Java で実装しています。

### Factory パターン未実装の場合
- Main.java
    ```java
    public class Main {
        public static void main(String[] args) throws Exception {
            // petType, foodType が引数で与えられたとします
            String petType = "dog";
            String foodType = "meat";
            
            // animal food の変数を初期化
            Animal animal = null;
            Food food = null;
            
            // pwrType によって Animal,　Food の instane を取得します
            switch(petType) {
            case "dog":
                // petType == "dog" の場合、Dogクラス の instance を取得
                animal = new Dog("dog");
                
                // petType == "dog" の場合の foodType ごとの適量な Foodクラス の instance を取得します
                switch(foodType) {
                case "meat":
                    // 100g のお肉
                    food = new Meat(100);
                    break;
                case "fish":
                    // 40g の魚
                    food = new Fish(40);
                    break;
                default:
                    // 適切な Case がない場合、なんと 300g の泥……。
                    food = new Mud(300);
                }
                break;
            case "cat":
                // petType == "cat" の場合、Catクラス の instance を取得
                animal = new Cat("cat");
                
                // petType == "cat" の場合の foodType ごとの適量な Foodクラス の instance を取得します
                switch(foodType) {
                case "meat":
                    // 20g のお肉
                    food = new Meat(20);
                    break;
                case "fish":
                    // 80g の魚
                    food = new Fish(80);
                    break;
                default:
                    // 適切な Case がない場合、なんと 100g の泥……。
                    food = new Mud(100);
                }
                break;
            default:
                // 適切な Case がない場合、なんとモンスターが……
                animal = new Monster("monster");
                switch(foodType) {
                case "meat":
                    food = new Meat(0);
                    break;
                case "fish":
                    food = new Fish(0);
                    break;
                default:
                    food = new Mud(1000);
                }
            }
            
            // Animal の名前を取得します
            System.out.println(animal.getName());
            
            // 鳴き声です
            System.out.println(animal.say());
            
            // 今日の食べ物です
            System.out.println(food.get());
        }
    }
    ```
 - Animal.java
    ```java
    abstract class Animal {
        private String name;
        
        Animal(String name) {
            this.name = name;
        }
        
        public String getName() {
            return "I'm " + this.name;
        }

        abstract public String say();
    }
    
    class Dog extends Animal {
        Dog(String name) {
            super(name);
        }
        
        public String say() {
            return "ワン";
        }
    }
    
    class Cat extends Animal {
        Cat(String name) {
            super(name);
        }
        
        public String say() {
            return "ニャー";
        }
    }
    ```
 - Food.java
    ```java
    abstract class Food {
        Number volume;
        
        Food(Number volume) {
            this.volume = volume;
        }
        
        abstract String get();
    }
    
    class Meat extends Food {
        Meat(Number volume) {
            super(volume);
        }
        
        String get() {
            return "this is a meat! volume: " + this.volume.toString() + "g";
        }
    }
    
    class Fish extends Food {
        Fish(Number volume) {
            super(volume);
        }
        
        String get() {
            return "this is a fish! volume: " + this.volume.toString() + "g";
        }
    }
    
    class Mud extends Food {
        Mud(Number volume) {
            super(volume);
        }
    
        String get() {
            return "this is a mud.... volume: " + this.volume.toString() + "g";
        }
    }
    ```
上記コードは Paiza.io に書いたので実際に動かしてみましょう。
[Factory パターン未実装コード - Paiza.io](https://paiza.io/projects/rPCxneJTTGIkNLTIm5ERLg)

### Factory パターン実装の場合
上記コードを Factory パターンに変えてみましょう
Paiza.io を使って下記のコードを追加してみましょう。
- AnimalCreator.java
    ```java
    abstract class AnimalCreator {
        abstract Animal factoryMethod();
    }
    
    class DogCreator extends AnimalCreator {
        Animal factoryMethod() {
            return new Dog("dog");
        }
    }
    
    class CatCreator extends AnimalCreator {
        Animal factoryMethod() {
            return new Cat("cat");
        }
    }
    ```
 - FoodCreator.java
    ```java
    abstract class FoodCreator {
        Number volume;
        
        FoodCreator(Number volume) {
            this.volume = volume;
        }
        
        abstract Food factoryMethod();
    }
    
    class MeatCreator extends FoodCreator {
        MeatCreator(Number volume) {
            super(volume);
        }
        
        Food factoryMethod() {
            return new Meat(volume);
        }
    }
    
    class FishCreator extends FoodCreator {
        FishCreator(Number volume) {
            super(volume);
        }
        
        Food factoryMethod() {
            return new Fish(volume);
        }
    }
    ```
- Factory.java
    ```java
    abstract class Factory {
        abstract Animal animalFactoryMethod();
        abstract Food meatFactoryMethod();
        abstract Food fishFactoryMethod();
        
        Food foodFuctoryMethod(FoodType foodType) throws Exception {
            switch(foodType) {
                case Meat:
                    return this.meatFactoryMethod();
                case Fish:
                    return this.fishFactoryMethod();
                default:
                    throw new Exception("error");
            }
        }
    }
    
    class DogFactory extends Factory {
        Animal animalFactoryMethod() {
            return new DogCreator().factoryMethod();
        }
        
        Food meatFactoryMethod() {
            return new MeatCreator(100).factoryMethod();
        }
        
        Food fishFactoryMethod() {
            return new FishCreator(40).factoryMethod();
        }
    }
    
    class CatFactory extends Factory {
        Animal animalFactoryMethod() {
            return new CatCreator().factoryMethod();
        }
        
        Food meatFactoryMethod() {
            return new MeatCreator(20).factoryMethod();
        }
        
        Food fishFactoryMethod() {
            return new FishCreator(80).factoryMethod();
        }
    }
    
    enum FoodType {
        Meat,
        Fish
    }
    ```
 
 これで Factory パターンを実装出来ました。
 仕上げに Main.java を修正しましょう。

 - Main.java
    ```java
    import java.io.BufferedReader;
    import java.io.InputStreamReader;
    
    public class Main {
        public static void main(String[] args) throws Exception {
            // factory, foodType が引数で与えられたとします
            Factory factory = new DogFactory();
            // Factory factory = new CatFactory();
            FoodType foodType = FoodType.Meat;
            // FoodType foodType = FoodType.Fish;
            
            // factory から Animal型 の instance を取得します
            Animal animal = factory.animalFactoryMethod();
            
            // factory から Food型 の instance を取得します。
            Food food = factory.foodFuctoryMethod(foodType);
            
            // Animal の名前を取得します
            System.out.println(animal.getName());
            
            // 鳴き声です
            System.out.println(animal.say());
            
            // 今日の食べ物です
            System.out.println(food.get());
        }
    }
    ```

随分と Main.java がスッキリしました！
実行して動きを確認してみましょう！

[Factory パターン未実装コード - Paiza.io](https://paiza.io/projects/vtixW3gkFQ-w8002CGPUOA)

### Factory パターンの効果
では、Factory パターンを適用する事でどのような効力があるのか見ていきましょう！

1. コードがスッキリする
    - instance 生成処理を Factoryクラスの中に隠蔽することでメインの処理が見やすくなります
1. if/switch 文 を書かなくて良くなる
    - if/switch 文は時にバグの温床になる事があります
        - 今回の例では、petType, foodType の文字列に間違いがあるとモンスター等予期せぬ事が……
1. 新しいペットを追加した時にメイン処理への影響が少ない
    - Factory クラスの中に隠蔽しているので、それ以外のコードの修正を最小限にすることができる
1. 新しいペット追加時に、実行時エラーをなくすことができる
    - if/switch 文ではなく、クラスの追加なので間違いがある場合コンパイルエラーになる

## Factory パターンまとめ
Factory パターンはデザインパターンの基本的なパターンの一つです。
繰り返し書く事が学習の近道！
是非実践で取り入れて行きましょう！
