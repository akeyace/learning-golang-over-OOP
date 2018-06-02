# Golang で実践向け (コマンドライン/.exe) ツールを作ってみよう！

Golang を学ぶには、実際に手を動かすのが一番の近道です。
そこで、第二回では Windows でも Mac でも Linux でもどこでも動作するコマンドラインツールを作ってみましょう！

最初に コマンドラインツールの概要、Golang の強力なクロスコンパイラについて説明します。
その後にハンズオンでコマンドラインツールを作っていきます！

## Golang 環境構築がまだの方
Golang 環境構築がまだの方は、次のページを見て Golang 環境を構築してみましょう！
[Golang 環境を構築して Web サーバ, API サーバを立てて動かしてみよう](https://github.com/akeyace/learning-golang-over-OOP/blob/master/lesson_01/03_hands_on.md)

不明な事は私や、周りの人に聞いてください！

## コマンドラインツールとは
コマンドラインツールは、コマンドプロンプト(Windows)、ターミナル(MacOSX、Linux等) から操作を行うツールの事を指します。
略称は CUI(Character User Interface) や、CLI(Command Line Interface) があります。
それに対して、ディスプレイ上をマウスで操作するツールを GUI(Graphical User Interface) と言います。 

特に、Linux を普段使う人には馴染み深い物ですが、普段 Windwos で GUI を使う人には敷居が高いように思えます。
しかし、逆に考えれば Golang で簡単にコマンドラインツールを作ってコマンドプロンプトから複雑な操作が出来るようになれば、とてもハッカーぽくて素敵だとは思いませんか？

### コマンドラインツールはどういうものがあるか？
コマンドプロンプトや、ターミナルから動作させるコマンドは全てコマンドラインツールです。
下記の様な物があります。
- cd
    - change directory コマンド
    - ディレクトリの変更を行う
- mkdir
    - make directory コマンド
    - コマンドラインから新しいディレクトリを作成することが出来る
- rm
    - remove コマンド
    - ファイル・ディレクトリを削除する恐ろしいコマンド
    - 特にオプション r (recursive: 再帰), f (force: 強制) を付けると破壊力抜群
- dep
    - dependency(依存) management コマンド
    - ライブラリの依存関係を管理してくれて便利

### コマンドラインツールの主な構成
コマンドラインは主に次の要素から構成されています。
- コマンドラインの使い方説明
    - どのようなオプションを使うのかを教えてくれる
- オプション
    - コマンドラインに引数として付加情報を与える事で様々な動作を行う
    - 例: dep ensure -add github.com/spf13/viper
        - 第一引数: ensure
        - dep コマンドの ensure 機能を使う
    - 第二引数: -add
        - ensure に対して 追加 を行う事を表している
    - 第三引数: github.com/spf13/viper
        - ensure -add に対して追加するライブラリを表している
- 設定ファイル
    - config ファイルとも言う
    - きめ細やかな設定を指定する事ができる

これでコマンドラインツールの概要は掴めたかと思います！

## Golang の強力なクロスコンパイラ
『クロスコンパイラ』とは、様々な環境にコンパイルするという意味です。
Golang は極々簡単に Windows, MacOSX, Linux と複数の環境にシングルバイナリ(単一ファイル) でコンパイルする事が可能です。
職場で、Windows, MacOSX, Linux が混在しているとツールを作るのにそれぞれ専用のアプリを作る必要がありますが、Golang なら一つのソースコードから、各環境で動作するアプリを作成できます。

### Golang クロスココンパイルの方法
Golang では作成したソースコードを実行するのに次のコマンドが用意されています。
|コマンド|概要|
|:--|:--|
|go run main.go|コンパイルしたファイルを即時実行する|
|go build .|コンパイルしてシングルバイナリの実行ファイルを出力する|
|go install .|コンパイルしたファイルを環境変数 $GOBIN で指定された場所にシングルバイナリで出力する。このコマンドを実行するとコマンドプロンプト、ターミナルからコマンドを実行出来るようになる|

では、クロスコンパイルはどうするのでしょうか？
なんと、あっと驚く方法で簡単にクロスコンパイル出来ちゃいます

- Window 32bit 用実行ファイルを出力する
    ```bash
    GOOS=windows GOARCH=386 go build .
    ```
- Window 64bit 用実行ファイルを出力する
    ```bash
    GOOS=windows GOARCH=amd64 go build .
    ```
- MacOSX 用実行ファイルを出力する
    ```bash
    GOOS=darwin GOARCH=amd64 go build .
    ```
- Linux 用実行ファイルを出力する
    ```bash
    GOOS=linux GARCH=amd64 go build .
    ```

なんと簡単な！

上記コマンドは次を表しています。
- GOOS
    - クロスコンパイルする OS
- GOARCH
    - クロスコンパイルする CPU アーキテクチャ

GOOS, GOARCH を付けない場合は、Golang をインストールしたパソコンに適切に合わせたコンパイルをします。
また、CPU アーキテクチャが同じ場合は GOOS のみを設定してもクロスコンパイル可能です

- 自分の環境に実行ファイルを出力する
    ```bash
    go build .
    ```
- CPU が同じ環境の Linux 用実行ファイルを出力する
    ```bash
    GOOS=linux go build .
    ```
- CPU が同じ環境の Linux 用実行ファイルを出力する
    ```bash
    GOOS=linux go build .
    ```

自分自身の環境は、go env コマンドを使う事で確認する事ができます。

## Golang でコマンドラインツールを作ろう！
さあ、いよいよお待ちかねコマンドラインツールの作成です。
ハンズオンなので、是非実際に作成してみましょう！

### コマンドライン作成用のライブラリを取り込む
Golang はライブラリ指向の言語です。
既にある便利なライブラリをどんどん活用しましょう！

1. 次のディレクトリを作成する
    - /Users/xxx/go/src/local/go-original-cli
1. Eclipse を起動し、上記で作成したディレクトリでプロジェクトを作成する
1. main.go ファイルを作成する
    ```go
        package main
    
    import (
        "log"
        "os"
    
        "github.com/urfave/cli"
    )
    
    func main() {
        err := cli.NewApp().Run(os.Args)
        if err != nil {
            log.Fatal(err)
        }
    }
    ```
1. コマンドラインで下記を実行する
    ```bash
    dep init
    ```
1. vendor, Gopkg.toml, Gopkg.lock が作成され。ライブラリが取り込まれます。
1. 早速コンパイルしてみましょう
    ```bash
    go build .
    ```
1. すると、ディレクトリ名(今回はgo-original-cli) のファイルがコンパイルされます
1. 実行してみましょう
    ```bash
    ./go-original-cli
 j   ```
1. なんとコマンドラインツールっぽいものが出来てしまいました！

### コマンドラインツールの設定
それでは、よりコマンドラインツールらしく設定を行っていきましょう。
設定方法は簡単です！

- main.go
    ```diff
      func main() {
    -     err := cli.NewApp().Run(os.Args)
    app := cli.NewApp()
    app.Name = "HSCLT: Hyper Scraping Command Line Tool!"
    app.Author = "Your Name"
    app.Email = "your email"
    app.Version = "0.1.0"
    app.Copyright = "Copyright(c) 2018 " + app.Author
    app.Usage = "you'll scraping easy!"

    err := app.Run(os.Args)
          if err != nil {
              log.Fatal(err)
          }
    ```
        - app := cli.NewApp() で コマンドラインツールのポインタを取得
    - app.Name 等で名称の設定等を行っています
    - Author, Email 等自身の設定を書いてみましょう！
- go build . でコンパイルして実行してみましょう！
    ```bash
    # help が表示される
    ./go-original-cli

    # これでも help が表示される
    ./go-original-cli --help

    # version 情報が表示される
    ./go-original-cli --version
    ```

### コマンドラインツールのコマンド実装
では、コマンドラインツールを作っていきましょう！

- main.go
    ```diff
      app.Usage = "you'll scraping easy!"

    + // app.Commands で新しくコマンドを追加できます
    + app.Commands = []cli.Command{
    +     {
    +         // コマンド名
    +         Name:    "get",
    +         // 短縮コマンド名
    +         Aliases: []string{"g"},
    +         // 説明文
    +         Usage:   "scraping website! ",
    +         // コマンド実行時の処理
    +         Action: func(c *cli.Context) error {
    +             // c.Args() で引数取得
    +             if website := c.Args().First(); website == "" {
    +                 // 引数がなかった場合はエラー
    +                 return cli.NewExitError("get command have to arg website. please read help by --help", 1)
    +             }
                  // 処理結果を出力
    +             fmt.Printf("scaping %s\n", c.Args().First())
    +             return nil
    +         },
    +     },
    + }
      
      err := app.Run(os.Args)
    ```

コンパイルして実行してみましょう！
すると、Help に追加したコマンドが表示されるようになりました。
さらに

```bash
# get の後に google.com と書くと、scaping google.com と表示されました
./go-original-cli get googel.com

# 引数がないとエラーが表示されます
./go-original-cli get
```

### コマンドラインツールをインストールする
コンパイルし、動作確認が出来たらインストールしてみましょう！

- install コマンドを実行
    ```bash
    go install .
    ```

なんとこれだけで OK です。
go install すると 実行ファイルが $GOBIN の中に取り込まれます。
実行してみましょう！

- 実行する
    ```bash
    go-original-cli
    ```

コマンドラインツールらしくなりました！
あとは、Action に処理を書いていけば好きな処理を簡単に呼びだす事が可能になります。

## Golang で実践向け (コマンドライン/.exe) ツールを作ってみよう！
コマンドラインツールを作るにも、ライブラリを利用することでいとも簡単に出来てしまいました。
後はアイディア次第で何でも出来る状態が整いました！
大事なのは兎に角作ること。
作れば作っただけ経験が蓄積されていきます！