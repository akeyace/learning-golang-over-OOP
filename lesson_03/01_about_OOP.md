# 関数型プログラミングとは
関数型プログラミングは幾つかの基本的な概念によって成り立っています。  
これらの概念一つ一つは簡単な物から難しい物まで含み、一筋縄ではいきません。  
そこで実際にプログラミングをしながら、『どうして関数型が良いのか？』を見ていきましょう。  
  
気軽にプログラミングをするためにこちらのサイトを使いましょう。
  
[Paiza.io](https://paiza.io/)

## 関数型プログラミング基礎
関数型プログラミングは広義では、高階関数 (関数を引数に取ったり、戻り値にしたりする)の利用と言えるかもしれません。  
map関数や、filter関数は、Java のストリームAPI でもお馴染みの物となりました。  
では、もっと踏み込んだ本質的な関数型プログラミングとはどのようなもものでしょうか？  
見ていきましょう

## 関数型プログラミングを構成する要素
関数型プログラミングはどのような概念、どのような要素によって成り立っているのでしょうか？  
それは次のような要素によって成り立っています。  

- 参照透過性 (Referential transparency)
- 関数合成 (Function composition)
- カリー化 (Currying)
- 部分適用 (Partial application)
- 不変性 (Immutable)
- モナド (Monad)

### 参照透過性
- 同じ入力に対して、同じ作用、同じ出力を持つこと
    - メンバ変数の値によって振る舞いが変わるケース
        - 例: publishFlg が false の場合は、保存処理のみ、true の場合は、公開処理も実行される
            ```java
            public Boolean save() {
                Boolean result = this.db.save();
                if (result && this.publishFlg) {
                    this.publishArticle();
                }
                return result
            }
            ```
    - 他の処理に依存するケース
        - 例: 他の処理に依存する場合
            ```java
            public String getFilePath(String fileName) {
                String path = new File(".").getAbsoluteFile().getParent();
                return path + fileName;
            }
            ```
        - 例: 参照透過性を持たせる場合
            ```java
            public String getFilePath(String path, String fileName) {
                return path + fileName;
            }

            // 条件によって値が変動する処理を引数にする
            getFilePath(new File(".").getAbsoluteFile().getParent(), fileName);
            ```

#### 具体的に何が良いの？
- より疎結合な設計を行いやすい
    - コンポーネント化しやすく、処理の使い回しがしやすくなる
- 状態と、処理を分離しやすくなる
- 引数によって必ず同じ値が返却されるので、テスタブルなコードを作りやすい

## 関数合成
- 複数の関数を組み合わせて一つの処理に纏めること
- 引数に関数を持つ関数を利用する
- 関数合成しない例
    ```java
    int result = 0;
    for(int e : values) {
      if(e > 3 && e % 2 == 0) {
        result = e * 2;
        break;
      }
    }
    ```
- ラムダ式による関数合成を行う例
    ```java
    int result = values.stream()
        .filter(e -> e > 3)
        .filter(e -> e % 2 == 0)
        .map(e -> e * 2)
        .findFirst()
        .orElse(0);
    ```
※参照元: [Java8 のイディオム - IBM developerWorks](https://www.ibm.com/developerworks/jp/java/library/j-java8idioms6/index.html)

#### 具体的に何が良いの？
- メンテナンス性の高いコードを作りやすい
- リーダブルなコードを作りやすい

### カリー化と部分適用
- 部分適用は、関数を部分的に、段階的に適用できるという物
- カリー化は部分適用を適用しやすくする物
- こちらの記事がわかりやすかったです！
    - [カリー化と部分適用（JavaScriptとHaskell）](https://qiita.com/7shi/items/a0143daac77a205e7962)

#### 具体的に何が良いの？
- 関数のカスタマイズが楽に！
    - Template Method パターンのように部分与える引数の置き換えによる修正が楽になる

### 不変性
- オブジェクトのメンバ変数というのは管理しづらい側面があり、時にバグを引き起こすトリガーになります
- そこで、オブジェクトの値は全て定数として扱うというのが不変性のポイント
- これはインスタンス生成時に値が決定され、そのオブジェクトの生存期間 (Life time) は値が変更されない事を意味し、それにより『うっかり』値を上書きする事によるバグを防ぎます

不変性の例
```java
import java.util.*;

final class Person {
    // final なメンバ変数 name はコンストラクタで値が設定された以降は値が変わらない
    public final String firstName;
    public final String lastName;
    
    // コンストラクタで name の値が設定される
    public Person(String firstName, String lastName) {
        this.firstName = firstName;
        this.lastName = lastName;
    }
    
    // name の値を取り出す事は可能
    public String getName() {
        return this.firstName + " " + this.lastName;
    }
    
    // 名前を変更する際は新しいオブジェクトを返す
    public final Person changeName(String lastName) {
        return new Person(this.firstName, lastName);
    }
}

public class Main {
    public static void main(String[] args) throws Exception {
        // Kazuko Sato さんという人がいたとする
        Person kazuko = new Person("Kazuko", "Sato");
        
        // Kazuko Sato
        System.out.println(kazuko.getName());
        
        // オブジェクトのハッシュコードを保持しておく
        int beforeObjectHashCode = kazuko.hashCode();
        
        // 名字が Takahashi さんに変わる
        kazuko = kazuko.changeName("Takahashi");
        
        // Kazuko Takahashi
        System.out.println(kazuko.getName());
        
        // オブジェクトのハッシュコードを保持しておく
        int afterObjectHashCode = kazuko.hashCode();
        
        // 最初のハッシュコードと比較すると、異なるオブジェクトであることがわかる
        System.out.println(beforeObjectHashCode);
        System.out.println(afterObjectHashCode);
        System.out.println(beforeObjectHashCode == afterObjectHashCode);
        
    }
}

```

#### 具体的に何が良いの？
- オブジェクトの変更が明確になる
    - 上記の例では、changeName メソッド実行時のみ新しいインスタンスが生成されます
    - バッドパターンでは、他のオブジェクトによって、いつ、どこで値が変更されるかが分からなくなり、データフローが不透明になる事があります
- 元々のオブジェクトに影響を与えずに新しいインスタンスを生成出来る
    - 上記の例では、名字が変わりましたが、以前の名字の時の情報を残しておきたい場合従来では、Clonable インタフェースを実装して、clone メソッドを実装するなどの手間が掛かります


### モナド
- 関数型プラグラミングを語る上で外せない概念にして、御しがたい存在、それがモナドです
- 参考URL
    -[モナド (プログラミング)](https://ja.wikipedia.org/wiki/モナド_(プログラミング))
    -[モナドを理解する - 迷える者への手引き] (https://www.infoq.com/jp/articles/Understanding-Monads-guide-for-perplexed)
    -[モナドはポケモン。数学が出てこないモナド入門](https://qiita.com/hiruberuto/items/8bbc0343bf794c368287)

#### 具体的に何が良いの？
- 語れる方がいらっしゃいましたら、宜しくお願い致します。


## どのように使われているの？
- map関数、filter関数などは Java を始め様々な言語で実装されており、慣れると非常に使いやすい
- 関数を引数にとる、コールバック処理などもお馴染み
-


## 関数型プログラミングまとめ
関数型プログラミングは、難しそうな概念ではあるものの、使えたら、便利そうな感じがしてきませんか？  
実際に利用してみると使いやすく、特に map関数、filter関数などは手放せなくなること請け合いです。  
  
食べず嫌いの方がいらっしゃいましたら、是非 Try してみましょう！
