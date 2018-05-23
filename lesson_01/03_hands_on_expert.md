# Golang ハンズオン ★上級編★
勉強会に手応えがない、、、
そういうマニアックな方向けの上級編です。
Let's Try!

## Webサーバ APIサーバ 拡張 - Web サーバから API を叩いて値を返す -
こちらは余裕のある人は試してみましょう！
勉強会枠内で出来たら素晴らしい！

1. server.go を開きます
1. import に "html/template"、"io" を追加します
    ```diff
      import (
    +     "html/template"
    +     "io"
          "net/http"
          "github.com/labstack/echo"
      )
    ```
1. Template を読み込めるように Struct を追加します
    ```diff
          "github.com/labstack/echo"
      )
    + type TemplateRenderer struct {
    +     templates *template.Template
    + }
    ```
1. TemplateRender にメソッドを追加します
    ```diff
      type TemplateRenderer struct {
          templates *template.Template
      }
    + func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    +     if viewContext, isMap := data.(map[string]interface{}); isMap {
    +         viewContext["reverse"] = c.Echo().Reverse
    +     }
    + 
    +     return t.templates.ExecuteTemplate(w, name, data)
    + }
    ```
1. main 関数に Template を読み込む処理を追加します
    ```diff
      func main() {
          e := echo.New()
    +     renderer := &TemplateRenderer{
    +         templates: template.Must(template.ParseGlob("public/views/*.html")),
    +     }
    +     e.Renderer = renderer
    ```
1. Template 用のディレクトリを作成します
    - プロジェクトルートに public/views ディレクトリを作りましょう
1. public/views に api-test.html を作りましょう
    ```html
    <html>
      <head>
        <script>
          function callAPI() {
            const ip = document.querySelector('#send-ip') || 'localhost:1323';
            const request = document.querySelector('#request');  
            const display = document.querySelector('#response > p');
            
            const url = 'http://' + ip.value + '/api/users/request';
            fetch(url, {
              method: 'POST',
              mode: 'cors',
              headers: {
                'content-type': 'application/json'
              },
              body: JSON.stringify({name: request.value})
            })
              .then(response => {
                return response.json();
              })
              .then(json => {
                if (json.code) {
                  throw Error('エラーコード: ' + json.code + ', エラーメッセージ: ' + json.message);
                };
                const result = "相手の名前: " + json.my_name + ", 貴方の名前: " + (json.your_name || "hogehoge丸");
                            
                display.innerHTML = result;
              })
              .catch(error => {
                alert(error)
              })
          }
        </script>
      </head>
      <body>
        <h1>{{index . "content-name"}}</h1>
      
        <h2>送信先IP、名前(ニックネームも可) を入力して下さい</h2>
        <label>送信先IP</label>
        <input type="text" id="send-ip" value="localhost:1323">
        <label>名前</label>
        <input type="text" id="request">
        <button type="submit" onClick="callAPI()">submit</button>
     
        <h2>レスポンス表示欄</h2>
        <div id="response">
          <p></p>
        </div>
      </body>
    </html>
    ```
1. server.go にルーティングを追加します
    - "cntent-name": に何か名前を付けてみましょう！
    - response.MyName の欄に自分の名前を書いてみましょう！
        ```diff
              e.GET("/api/users/:name", func(c echo.Context) error {
                  name := c.Param("name")
                  return c.JSON(http.StatusOK, Response{Name: name})
              })
        +     e.GET("/api-test", func(c echo.Context) error {
        +         return c.Render(http.StatusOK, "api-test.html", map[string]interface{}{
        +             "content-name": ここにコンテンツの名前を適当につけてみよう！,
        +         })
        +     })
        +     e.POST("/api/users/request", func(c echo.Context) error {
        +         request := new(AnswerRequest)
        +         if err := c.Bind(request); err == nil {
        +             return c.JSON(http.StatusBadRequest, ErrorResponse{100, "エラーです"})
        +         }
        +         response := new(AnswerResponse)
        +         response.MyName = ここに名前をを入力してみよう！
        +         response.YourName = request.Name
        +       
        +         ip := net.ParseIP(c.RealIP()).To4()
        +       
        +         log.Infof("貴方のサーバにアクセスがありました: {name: %s, IPAddr: %s", request.Name, ip)
        +       
        +         return c.JSON(http.StatusOK, response)
        +     })
        ```
1.  実は上記コードを追加しただけではエラーが発生します
    - import に何かを追加します。試してみましょう！
        - Log.Infof は、"github.com/labstack/gommon/log" を import するので注意！
1. API Request, Response 用の Struct を作ろう
    - リクエストもレスポンス同様に Struct が必要になります。
    - `json:"name"` というのは、タグと呼ばれる物で値の振る舞いを変更することができます
        - この場合は、json のリクエスト、レスポンス時の名前を変更できます。
            ```diff
                   Response struct {
                       Name string
                   }
            +     AnswerRequest struct {
            +         Name string `json:"name" query:"name"`
            +     }
            +     AnswerResponse struct {
            +         MyName string `json:"my_name"`
            +         YourName string `json:"your_name"`
            +     }
            
            ```
1. 上記コードだけでは動きません！
    - 何かの Response が不足している？
1. 各ファイルを保存します
1. Web サーバを再起動しましょう
1. http://localhost:1323/api-test をブラウザで開いてみましょう
1. 名前を入力して『submit』ボタンをクリック
1. 名前が表示されれば成功です！

## 更に拡張してみよう
現時点では API は同一のホスト (同じドメインのサイト) のみ対応しています。
他の Webサーバ にも対応するよう変更してみましょう

1. CORS の設定を追加する
    1. server.go に ライブラリ "github.com/labstack/echo/middleware" を追加する
    1. dep ensure を実行してライブラリを追加する
    1. CORS を全てのホストを受けられ可能に設定する
        ```diff
              e := echo.New()
        
        +     e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        +         AllowOrigins: []string{"*"},
        +         AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
        +     }))
        
        ```
1. ファイルを保存して、サーバを再起動しましょう
1. http://localhost:1323/api-test をブラウザで開いてみましょう
1. 送信先IP を変更して実行してみましょう！
    - 送信先IP は wifi を通じて探して見て下さい

