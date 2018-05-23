# Golang 環境を構築して Web サーバ, API サーバを立てて動かしてみよう！
Golang を学ぶには、実際に手を動かすのが一番の近道です。
そこで、第一回では環境を構築し、Web サーバを動かしてみましょう！

## Golang 環境構築
勉強会冒頭で、必要なアプリケーションをダウンロードしたかと思います。
そのアプリケーションをインストールし、各種設定を行っていきます。

### GOPATH の設定
Golang では、環境変数 GOPATH の設定が非常に重要になります。
これは、Golang の思想であるライブラリ指向が関連しているともいえます。
百聞は一見にしかず。
設定していきましょう

1. Golang 開発用のディレクトリを作成する
    - Mac の場合
      1. /Users/xxx/go を作成するため、ターミナルを開き次のコマンドを入力
          ```bash
          # ~/ は、現在のユーザ (current user) のディレクトリを表す
          mkdir ~/go
          ```
          - 手作業で GUI で入力するのも可
    - Windows の場合
      1. ﻿C:\Users\xxx\go を作成するため、コマンドプロンプトを開き次のコマンドを入力
          ```bash
           mkdir %USERPROFILE%\go
           ```
          - 手作業で GUI で入力するのも可
1. Golang をインストールする
    - インストール先のディレクトリ等はデフォルトのままで OK
1. Golang インストールの確認
    - Mac の場合
        1. ターミナルを開き、次のコマンドを入力
            ```bash
            go env
            ```
    - Windows の場合
        1. コマンドプロンプトを開き、次のコマンドを入力
            ```bash
            go env
            ```
1. 環境変数の一覧が表示されたら成功です
1. 環境変数の設定を行います
    - Mac の場合
        1. .bashrc を開きます
            ```bash
            vi ~/.bashrc
            ```
        1. 下記を追記します
            ```bash
            export GOPATH=$HOME/go
            export PATH=$PATH:$GOPATH/bin
            ```
        1. 設定を読み込みます
            ```bash
            source ~/.bashrc
            ```
    - Windows の場合
        - デフォルトで良い感じに設定される様子？
        - 自身で設定する場合
            1. Windows メニュー から、『コンピュータ』を右クリック → 『プロパティ』をクリック
            1. 『システムの詳細設定』をクリック
            1. 『環境変数』ボタンをクリック
            1. 下記を追記します
                - 変数名: GOPATH
                - 変数値: %USERPROFILE%\go
1. GOPATH の確認
    - Mac の場合
        1. ターミナルを開き、次のコマンドを入力
            ```bash
            go env
            ```
    - Windows の場合
        1. コマンドプロンプトを開き、次のコマンドを入力
            ```
            go env
            ```
1. GOPATH に設定した情報が反映されていれば OK です

### IDE のインストール
1. Eclipse をインストールする
1. Eclipse を実行する
1. ワークスペースを聞かれるので GOPATH ディレクトリの src を入力し、『起動』ボタンを押します
    - Mac の場合: /Users/xxx/go/src
    - Windows の場合: C:\Users\xxx\go\src
1. メニューバー『ヘルプ』 → 『新規ソフトウェアのインストール…』をクリック
1. インストールダイアログが開くので、『追加ボタン』をクリック
1. リポジトリーの追加ダイアログが開くので、下記を入力して『OK』ボタンをクリック
    - 名前: GoClipse
    - ロケーション: http://goclipse.github.io/releases/
1. リポジトリーの追加ダイアログが開くので、下記を入力して『OK』ボタンをクリック
    - 名前: TM Terminal
    - ロケーション: http://download.eclipse.org/tm/terminal/marketplace
1. インストールダイアログに戻り、リストに『GoClipse』、『TM Terminal 4.1 Main Features』が表示されるので、チェックボックスをチェックして『次へ』ボタンをクリック
1. インストール詳細と表示されるので、『次へ』ボタンをクリック
1. ライセンスのレビューと表示されるので、『使用条件の条項に同意します』チェックボックスをチェックして『完了』ボタンをクリック
1. GoClipse、TM Terminal のインストールが開始され、『これらの証明書を信頼しますか？』とダイアログが表示されるので、チェックボックスをチェックして『選択を受け入れる』ボタンをクリック
1. ソフトウェア更新ダイアログが表示されるので『いますぐ再起動』ボタンをクリック
1. これでインストール完了！ IDE の準備完了です！

### IDE に GOPATH を設定する
1. 上記 go env コマンドで、GOROOT と、GOPATH の path を
覚えておく
   - GOROOT="/usr/local/go" ← の /user/local/go が path
   - GOPATH="/User/ak/go" ← の /User/ak/go が path
1. IDE のメニューバーから『Eclipse』 → 『環境設定』をクリック
1. 一覧から 『Go』を選択
1. 下記を入力して『適用』ボタンをクリック
    - ディレクトリー欄: 上記 GOROOT の path を入力
    - 『GOPATH 環境変数と同じ値を使用する』チェックボックスのチェックを外し、上記 GOPATH の path を入力
1. 一覧から『Go』をクリックしサブメニューを開き『ツール』を選択
1. gocode, guru, godef の『ダウンロード』ボタンをそれぞれクリックする
    - ダウンロードボタンをクリックすると Download gocode 等のダイアログがそれぞれ開かれるので 『OK』ボタンをクリックする
1. 『適用して閉じる』ボタンをクリックして環境設定を閉じる

これで設定完了！

### Golang インストールまとめ
Golang 開発環境構築は他の言語と比べると比較的簡単ですが、唯一のハマりどころが GOPATH。
これは一つの約束事で、GOPATH で指定されたディレクトリの下にある src ディレクトリで開発を行っていく事になります。
※direnv 等のツールにより、GOPATHを個別に設定することも出来ますが、まずはスンダードにいきましょう

## Golang を動かしてみよう
では、お待ちかね、Golang を動かしましょう！

と、その前に、、、Golang はライブラリ指向だと伝えました。その流儀に従ってディレクトリを作成しましょう

### 開発用ディレクトリを作成する
1. GOPATH の src ディレクトリの下に次のディレクトリを作成します
    - github のアカウントが有る場合
        - GOPATH/src/github.com/xxx
            - xxx は自身のアカウント名
    - github アカウントが無い場合
        - GOPATH/src/local

### 最初のプロジェクト作成
では、IDE でプロジェクトを作成してみましょう！
1. IDE のメニューバーから『ファイル』 → 『新規』 → 『プロジェクト』をクリック
1. 新規プロジェクトダイアログが開くので、『Go』 → 『Goプロジェクト』を選択し、『次へ』をクリック
1. 新規 Go プロジェクトダイアログが開くので、下記を入力して『』ボタンをクリック
    - プロジェクト名: go-first-project!
    - 『デフォルト・ロケーションを使用。』チェックボックスのチェックを外す
    - ディレクトリー: 上記で作成した開発用ディレクトリの下の階層に 『/go-first-project』を追記したもの
        - 例: /Users/xxx/go/src/github.com/xxx/go-first-project
        - 例: C:\Users\xxx\go\src\local\go-first-project
            - go-first-project ディレクトリは自動的に作成されるので、事前に作成しなくても大丈夫
1. 下記画像の用になれば OK

### Golang で最初のプログラミング！
お約束のハローワールドを書いてみましょう

1. IDE のプロジェクトエクスプローラーを開く
1. ルート(この場合は go-first-project) を選択し右クリック → 『新規』 → 『Go ファイル』をクリック
1. 下記を入力して『完了』ボタンをクリック
    - ソース・ファイル: test.go
1. test.go が開かれるので次を記述
    - package 名を main に変更する
        ```diff
        - package go-first-project
        + package main
        ```
    - import の中に "fmt" を記述
        ```diff
          import (
        +     "fmt"
          )
        ```
    - func main を作る
        ```go
        func main() {
      
        }
        ```
    - Hello World!! を出力する
        ```diff
          func main() {
        +     fmt.Println("Hello World!")
          }
        ```
    - 最終的に次のようになります
        ```go
        package main
        
        import (
            "fmt"
        )
        
        func main() {
            fmt.Println("Hello World")
        }
        ```
1. ファイルを保存する
1. エディタ上で右クリック → 『実行』 → 『1 Go アプリケーション』をクリック
1. コンソールに Hello World! と表示されたら成功です

### Golang で Web ページを作る
では、今回の目的である Web サーバ、API サーバを作っていきましょう！
とはいえ、一から作るというのも近代的ではありません。
ここは、Golang のライブラリ指向を十分に活用しましょう

#### パッケージマネージャをインストールする
近代の開発では切っても切り離せないパッケージマネージャ。
これは、ライブラリをいい感じに利用するための素晴らしいツールです。

Golang には、dep という純正のパッケージマネージャがあります。
dep を使う事で、プロジェクト毎に異なりバージョンのライブラリを導入することが可能になります。

早速インストールしましょう！

1. IDE のアイコンメニューから『ターミナル』をクリック
    - ※ターミナル、コマンドプロンプト操作に慣れている方は任意のツールを使って下さい
1. ターミナルの起動ダイアログが開くので下記を入力し『OK』ボタンをクリック
    - ターミナルの選択: Remote Terminal
    - ラジオボタンから『ローカル』を選択
    - エンコード: UTF-8
1. ターミナル上で次のコマンドを実行
    ```bash
    go get -u github.com/golang/dep/cmd/dep
    ```

#### Golang の Web フレームワークを導入する
フレームワークは、日本語では枠組み。
何かというと、基本的な機能が予め備わっているアプリケーションの土台です。
Web サイトを作るなら、Web の基本機能を提供してくれるフレームワークを使うのが一番の近道です！
早速 dep で導入しましょう！

1. ターミナル、コマンドプロンプトでプロジェクトのルートを開く
    - 例: cd ~/go/src/github.com/xxx/go-first-project
1. ターミナル、コマンドプロンプトで次のコマンドを実行する
    ```bash
    dep init
    dep ensure -add github.com/labstack/echo@^3.1
    ```
1. IDE のプロジェクト・エクスプローラーをウィンドウ内で右クリック → 『リフレッシュ』をクリック
1. すると、vendor, Gopkg.lock, Gopkg.toml が追加されました
    - vendor: ライブラリが保存される
    - Gopkg.lock: ライブラリの具体的なバージョン情報
    - Gopkg.toml: 利用するライブラリ
1. これで準備は完了です

#### Golang で web サーバ実行
では Web サーバを立ててみましょう

1. ルート(この場合は go-first-project) を選択し右クリック → 『新規』 → 『Go ファイル』をクリック
1. 下記を入力して『完了』ボタンをクリック
    - ソース・ファイル: server.go
1. server.go が開かれるので次を記述
    - package 名を main に変更する
        ```diff
        - package go-first-project
        + package main
        ```
    - import の中に "net/http", "github.com/labstack/echo" を記述
        ```diff
          import (
              "net/http"
        +     "github.com/labstack/echo"
          )
        ```
    - func main を作る
        ```go
        func main() {
      
        }
        ```
    - echo フレームワークの機能をポインタとして呼び出す
        ```diff
          func main() {
        +     e := echo.New()
          }
        ```
    - ルーティング(http://xxx.yyy/zzz の zzz にアクセスした際にどのようなページを表示するか)
        ```diff
          func main() {
              e := echo.New()
        +     e.GET("/", func(c echo.Context) error {
        +         return c.String(http.StatusOK, "Hello, World!")
        +     })
          }
        ```
    - サーバ起動処理と、ロギング処理を記述
        ```diff
          func main() {
              e := echo.New()
              e.GET("/", func(c echo.Context) error {
                    return c.String(http.StatusOK, "Hello, World!")
              })
        +     e.Logger.Fatal(e.Start(":1323"))
          }
        ```
    - 最終的に次のようになります
        ```go
        package main
        
        import (
            "net/http"
            "github.com/labstack/echo"
        )
        
        func main() {
            e := echo.New()
            e.GET("/", func(c echo.Context) error {
                return c.String(http.StatusOK, "Hello, World!")
            })
            e.Logger.Fatal(e.Start(":1323"))
        }
        ```
1. ファイルを保存する
1. ターミナル、コマンドプロンプトで次のコマンドを入力
    ```bash
    go run server.go
    ```
1. これで port 番号1323 の Webサーバ が立ち上がります
1. http://localhost:1323 をブラウザで開きます
1. ブラウザに Hello World! と表示されたら成功です！
1. http://localhost/1323/test をブラウザで開いてみましょう
1. すると、{"message":"Not Found"} と表示されます
1. IDE のターミナルを見るとエラーメッセージが表示されていることが確認できます
1. 先程の server.go を修正してみましょう
    - http://localhost:1323/test に対応するルーティングを書いてみましょう
        ```diff
            func main() {
                e := echo.New()
                e.GET("/", func(c echo.Context) error {
                    return c.String(http.StatusOK, "Hello, World!")
                })
        +       e.GET("/test", func(c echo.Context) error {
        +           return c.String(http.StatusOK, "testを表示するよ！")
        +       })
                e.Logger.Fatal(e.Start(":1323"))
            }
        ```
1. terminal で Ctrl+C と入力して Webサーバを落とします
1. 再度 Webサーバ を立ち上げましょう
      ```bash
      go run server.go
      ```
1. http://localhost:1323/test をブラウザで開いてみましょう

#### API サーバを立ち上げる
今度は APIサーバとして動作させてみましょう

1. API の レスポンス用の Struct を定義する
    - Golang では API のレスポンス毎に Struct を定義します
        - Name を返す API になｒます
        - レスポンスで使う Struct は Public の必要があるので頭文字は大文字になります
            ```diff
              import (
                  "net/http"
                  "github.com/labstack/echo"
              )
            + 
            + type (
            +     Response struct {
            +         Name string
            +     }
            + )
            ```
    - API 用のルーティングを作ります
        - API らしく、/api から開始して、Path パラメータ (URL の Path をパラメータとして利用する places/osaka/no/1001 みたいな物) を使ってみましょう
            ```diff
                  e.GET("/", func(c echo.Context) error {
                      return c.String(http.StatusOK, "Hello, World!")
                  })
                  e.GET("/test", func(c echo.Context) error {
                      return c.String(http.StatusOK, "testを表示するよ！")
                  })
            +     e.GET("/api/users/:name", func(c echo.Context) error {
            +         name := c.Param("name")
            +         return c.JSON(http.StatusOK, Response{Name: name})
            +     })
                  e.Logger.Fatal(e.Start(":1323"))
              }
            ```
    - Path パラメータを取り出してみます
        ```diff
              e.GET("/", func(c echo.Context) error {
                  return c.String(http.StatusOK, "Hello, World!")
              })
              e.GET("/test", func(c echo.Context) error {
                  return c.String(http.StatusOK, "testを表示するよ！")
              })
              e.GET("/api/users/:name", func(c echo.Context) error {
        +         name := c.Param("name")
              })
              e.Logger.Fatal(e.Start(":1323"))
          }
        ```
    - API なのでレスポンスを返す時は JSON 形式にしてみましょう
        - c.JSON とすることで、JSON 型のレスポンスを返すことができます
        - 第一引数: http.StatusOK
        - 第二引数: この API のレスポンスになります
            ```diff
                  e.GET("/", func(c echo.Context) error {
                      return c.String(http.StatusOK, "Hello, World!")
                  })
                  e.GET("/test", func(c echo.Context) error {
                      return c.String(http.StatusOK, "testを表示するよ！")
                  })
                  e.GET("/api/users/:name", func(c echo.Context) error {
                      name := c.Param("name")
            +         return c.JSON(http.StatusOK, Response{Name: name})
                  })
                  e.Logger.Fatal(e.Start(":1323"))
              }
            ```
    - 最終的にこのようになります
        ```go
        package main
        
        import (
            "net/http"
            "github.com/labstack/echo"
        )
        
        type (
            Response struct {
                Name string
            }
        )
        
        func main() {
            e := echo.New()
            e.GET("/", func(c echo.Context) error {
                return c.String(http.StatusOK, "Hello, World!")
            })
          
            e.GET("/test", func(c echo.Context) error {
                return c.String(http.StatusOK, "testを表示するよ！")     
            })
          
            e.GET("/api/users/:name", func(c echo.Context) error {
                name := c.Param("name")
                return c.JSON(http.StatusOK, Response{Name: name})
            })
            e.Logger.Fatal(e.Start(":1323"))
        }
        ```
1. IDE のターミナルを開き Ctrl+C で Webサーバを落とします
1. 再度 Webサーバ を立ち上げましょう
    ```bash
    go run server.go
    ```
1. http://localhost:1323/api/users/xxx をブラウザで開いてみましょう
    - xxx の部分は好きな文字列を入力してみて下さい
1. すると JSON 形式を返す API サーバが出来ました！
