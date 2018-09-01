# ドメイン駆動設計
ドメイン駆動設計とは、Eric Evans による 2003年 出版の著書 『Domain-driven design』によって提案されたソフトウェアの設計手法です。
通称 DDD。

2003年からジワジワ浸透し、難しい概念などと言われつつも現場で揉まれてベストプラクティスの一つとして定着した感があります。

このドメイン駆動設計、とっつきにくい部分も確かにあります。
そこで、今回はドメイン駆動設計を掻い摘んで実践向きに説明しようと思います！

## ドメイン駆動設計は何が良いのか？
ドメイン駆動設計を使う事のメリットは何でしょうか？
システム開発というのは重労働ですが、それ以上に難解なのはシステム運用と言って過言ではありません。
度重なる追加回収で草臥れてしまったシステムは枚挙に暇がありません。

そしてドメイン駆動設計 最大のメリットはシステム運用に柔軟さと堅牢さをもたらすことです。

## ドメイン駆動設計ってそもそも何なのか？
ドメイン駆動設計のメリットを見る前に、ドメイン駆動設計とは何かを見ていきましょう。

設計手法にも様々な物があります。
データフローを中心に置くケースや、ワークフローを中心とするケース、オブジェクトを中心とするケース。

ドメイン駆動設計では、業務領域(ドメイン) にクローズアップし、業務領域に関連する物をドメインモデルとして抽出・オブジェクト化し、ビジネスロジックなどの処理と隔離させます。

システムは時に業務要件とは掛け離れた書かれ方をすることがあります。
それをあくまで業務要件を中心とし、実際に業務で使われる用語を使いシステム利用者、システム開発者で共通認識を持つというのが一つのポイントになります。

そして、ドメインモデルは次の要素によって構成されます。

- エンティティ (参照オブジェクト): ドメインモデル内のオブジェクトであり、その属性によってではなく、連続性と識別性によって定義される。
- 値オブジェクト: 事物の特性を記述するオブジェクトである。特に識別する情報はなく、通例、読み出し専用のオブジェクトであり、Flyweight パターンを用いて共有できる。
- サービス: 操作がオブジェクトに属さない場合に、問題の自然な解決策として、操作をサービスとして実現することができる。サービスの概念は、GRASPにおいて"純粋人工物"と呼ばれるものである。
- リポジトリ:ドメインオブジェクトを取得するメソッドは、記憶域の実装を簡単に切り替えられるようにするため、専門のリポジトリオブジェクトに処理を委譲するべきである。
- ファクトリー : ドメインオブジェクトを生成するメソッドは、実装を簡単に切り替えられるようにするため、専門のファクトリーオブジェクトに処理を委譲するべきである。


難しい！！

## 簡単アバウト ドメイン駆動設計
このようにドメイン駆動設計は一般的に難しい概念と言われます。
なので、噛み砕いてみましょう。

### ドメインってなに？
一言でいうと、システム利用者の業務内容。
例えば、配送システムの場合、出荷センターを中心とした荷物の配送を行う各業務がドメインと言えます。

### ドメインモデルってなに？
ドメイン(業務領域)内の登場人物や、処理内容といった物。

例えば、配送システムの場合……
- 登場人物
    - 荷物
    - トラック
    - 工場の各装置
    - 作業員
    - 配送先
    - 仕入先
- 処理内容
    - 出荷
    - 仕入
    - 顧客管理

それらの登場人物を、システム利用者の呼称ベースでモデル化したものと言えるでしょう。

### エンティティってなに？
大雑把に一言でいうと、DB に紐づく各テーブル毎に特性の異なるオブジェクト。
上記の荷物は、荷物DB から 荷物ID で一意に選択されたレコードは、荷物エンティティというオブジェクトとして振る舞う。
そして、荷物エンティティと、トラックエンティティでは、それぞれ違うメソッドを保有する

### 値オブジェクトってなに？
所謂オブジェクトのメンバ変数……なのですが、ドメイン駆動設計では、メンバ変数にもプリミティブな型ではなく、そのエンティティに相応しいオブジェクトとして定義します。
例えば、、、

- 普通のオプジェクト設計 (https://paiza.io/projects/1932AxnNYwfKc9y4dATFTA)
    ```java
    class Glass {
        // 種別
        private String type;
        // グラスのサイズ
        private Number size;
        // 残容量
        private Number amount;
    
        Glass(Number size) {
            this.size = size;
            this.amount = 0;
        }
        // 注ぐ
        public void pour(Number pourAmount, String type) {
            this.type = type;
            this.amount = this.amount.intValue() + pourAmount.intValue();
        }
    
        // 飲む
        public void drink(Number drinkAmount) {
            this.amount = this.amount.intValue() - drinkAmount.intValue();
        }

        // 確認
        public Number check() {
            return this.amount;
        }
    }
    ```
- ドメイン駆動設計(https://paiza.io/projects/yVFstKSL3XUDUIwynHT1Hw)
    ```java
    class Glass {
        // 種別
        private String type;
        // グラスのサイズ
        private Number size;
        // 残容量
        private GlassAmount amount;
    
        Glass(Number size) {
            this.size = size;
            this.amount = new GlassAmount();
        }
        
        // 注ぐ
        public void pour(Number pourAmount, String type) {
            this.type = type;
            this.amount.add(pourAmount.intValue());
        }
    
        // 飲む
        public void drink(Number drinkAmount) {
            this.amount.subtract(drinkAmount.intValue());
        }
        // 確認
        public Number check() {
            return this.amount.check();
        }
    }
    class GlassAmount {
        private Number amount;
        
        GlassAmount() {
            this.amount = 0;
        }
        
        public void add(Number amount) {
            this.amount = this.amount.intValue() + amount.intValue();
        }
        
        public void subtract(Number amount) {
            this.amount = this.amount.intValue() - amount.intValue();
        }
        
        public Number check() {
            return this.amount;
        }
    }

    ```

このメリットはなにかというと、処理の置き換えをオブジェクトの置き換えで対応出来るという点になります。

### サービスってなに？
所謂ビジネスロジック。
次の 2種類 のサービスがある

- ドメインサービス
    - エンティティや、値オブジェクトに持たせたくないドメイン関連のビジネスロジック
- アプリケーションサービス
    - 上記ドメインサービスや、ドメインモデルを操作するためのサービス


### リポジトリってなに？
データ永続化を行うオブジェクト。
DAO に近いが、DB に関わらず広範囲のデータを永続化させる
リポジトリは差し替え可能になっていたりする

### ファクトリってなに？
ファクトリーパターン です。
オブジェクトを生成する際に、ファクトリーオブジェクトに委譲することで、色々安定する

## つまりドメイン駆動設計って何なのか？
簡単ではない存在、それがドメイン駆動設計です。
この次の章で詳しくみていきましょう。
