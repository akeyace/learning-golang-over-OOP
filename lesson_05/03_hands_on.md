# Golang と js フレームワークで リッチ Web アプリケーションを作ろう！

こちらをインストール！
https://github.com/olebedev/go-starter-kit
1. GOPATH 配下に下記のディレクトリを作成する
    - ~/go/src/github.com/xxx/go-js-project
2. dep をインストールする
    ```bash
    # xxx は、自身のgithub アカウントだとなおよし
    cd ~/go/src/github.com/xxx/go-js-project
    go get -u github.com/golang/dep/cmd/dep
    ```
3. go-starter-kit をインストールする
    ```bash
    dep github.com/olebedev/go-starter-kit
    ```
