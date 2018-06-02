# Golang ハンズオン ★上級編★
勉強会に手応えがない、、、
そういうマニアックな方向けの上級編です。
Let's Try!

## Golang で実践向け (コマンドライン/.exe) ツールを作ってみよう！ 中身もきっちり実装しよう！
それではコマンドの中身を作りましょう。
『Action: func...』 の中に処理を記述していきます。
察しの良い方は既に理解されているやもしれません。
今回作るのはスクレイピングツール。
Website の中身を取得するツールです。

### スクレイピング 用の処理を実装
1. scraping ディレクトリを作成
1. scraping/main.go を作成
    ```go
    package scraping
    
    import (
        "fmt"
        "log"
        "net/http"
    
        "github.com/PuerkitoBio/goquery"
    )
    
    func Scrape() {
        // Request the HTML page.
        res, err := http.Get("http://metalsucks.net")
        if err != nil {
            log.Fatal(err)
        }
        defer res.Body.Close()
        if res.StatusCode != 200 {
            log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
        }
    
        // Load the HTML document
        doc, err := goquery.NewDocumentFromReader(res.Body)
        if err != nil {
            log.Fatal(err)
        }
    
        // Find the review items
        doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
            // For each item found, get the band and title
            band := s.Find("a").Text()
            title := s.Find("i").Text()
            fmt.Printf("Review %d: %s - %s\n", i, band, title)
        })
    }
    ```
- main.go を修正
    ```diff
      import (
          "log"
          "os"
          "fmt"
  
          "github.com/urfave/cli"
    +     "github.com/akeyace/learning-golang-over-OOP/lesson_02/03_hands_on/scraping"
      )
    ```
    
    ```diff
      Action: func(c *cli.Context) error {
          if website := c.Args().First(); website == "" {
              return cli.NewExitError("get command have to arg website. please read help by --help", 1)
          }
    +     scraping.Scrape()
          fmt.Printf("scaping %s\n", c.Args().First())
              return nil
          },
    ```
