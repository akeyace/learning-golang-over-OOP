# Golang で低レイヤー処理を扱う
ブロックチェーンを学び、早速実装したい気持ちになったかと思います。
しかし、ここは一呼吸置いて実際にブロックチェーンを作成する前に、
Golang での低レイヤー処理の扱い方を学びましょう。

## なぜ低レイヤー処理を学ぶのか？
ブロックチェーンでは、暗号化技術など byte単位の処理を扱います。
良いものを作る際には原理を知っておいて損はありません。
また、Golang の文化では、低レイヤー処理を利用する際に既存のライブラリで取り扱わず、ちょっとした処理は自身で実装するようです。

ちょっと不便な所ではありますが、郷に入っては郷に従え(Go だけに)！
どのように byte を扱うか見ていきましょう

## 準備
いつものように https://paiza.io を利用してコードを実行してみましょう！

## Golang での byte の扱い
どうやって byte を扱うか見ていきましょう

### シングルクォーテーションの利用
Golang では string型 を扱う際にダブルクォーテーションを使いました。

```go
package main

import "fmt"

func main(){
    variable := "string型です"
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

感の良い方は気づきましたね！？

Golang で bit を扱う際は シングルクォーテーションを使います。

```go
package main

import "fmt"

func main(){
    variable := 'a'
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

a に該当する文字コード 97 と型が int32 と出力されました。

int32 って byte?

これは、文字に該当するコードを返すため、例えば次の様にすると、、、


```go
package main

import "fmt"

func main(){
    variable := '字'
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

文字コード 23383 が出力されました。
これは byte ではなく、文字コードです。

シングルクォーテーションを使う際は、一文字しか入力されない事に気をつけましょう。

```go
package main

import "fmt"

func main(){
    variable := '文字' //エラーが返る
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

### byte型 の扱い
では byte を扱っていきましょう

```go
package main

import "fmt"

func main(){
    variable := byte('a')
    // variable := byte("a") ダブルクォーテーションではエラーが返る
    // variable := byte('ab') シングルクォーテーションでは一文字しか扱えない
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

97 と出力され、型は uint8 でした。
この uint8 が Golang における byte を表す型です。

byte を扱う際には便利な方法があります。


```go
package main

import "fmt"

func main(){
    variable := []byte("abc")
    // variable := []byte('a') 配列にする場合はシングルクォーテーションではエラーが返ります。
    // variable := []byte('abc') シングルクォーテーションでは一文字しか扱えない
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

[97 98 99] と配列が出力され、型は []uint8 でした。

この様にして byte 単位の処理を行います。

### bit の扱い
では更にディープに bit を扱う場合はどうなるでしょうか？

前提条件として、皆様ご存知の通り、bit は 2進数を表し、byte は 8桁の bit により、0〜255の値を取ります。

Golang では bit演算子と言われる演算子が組み込まれています。

```go
package main

import "fmt"

func main(){
    variable := 1 << 8
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

256と出力されました。
これは 2進数で 1 を 8桁分 左にずらしたことを意味します。
つまり、、、
1 0000 0000 
という状態になった事を表します。

```go
package main

import "fmt"

func main(){
    variable := 1024 >> 4
    fmt.Println(variable)
    fmt.Printf("型は%T", variable)
}
```

こんどは 64 と出力されました。
これは 1024 を 2進数で表すと……
100 0000 0000
と、なるのを左に4つ動かすと……
100 0000
となり、これは10進数で 64 を表します。

### 事前宣言済み擬似定数 iota による bit 演算
iota は、同一 const 内で iota が登場する度に整数が加算される(インクリメント)性質を持っています。
bit を扱う際にはかなり便利？

```go
package main

import "fmt"

const (
    bit1 = 1 << iota
    bit2 = 1 << iota
    bit3 = 1 << iota
    bit4 = 1 << iota
)

func main(){
    fmt.Println(bit1)
    fmt.Println(bit2)
    fmt.Println(bit3)
    fmt.Println(bit4)
}
```

### 更に bit演算！
bit を扱う際には便利なライブラリが用意されています。

```go
package main

import (
    "fmt"
    "math/bits"
)

func main(){
    var variable uint8 = 12
    // 12 を 2進数で表す (%b は Binary: 2進数表現)
    fmt.Println(fmt.Sprintf("%b", variable))

    fmt.Println("8bit で表した場合に、左からいくつ 0 が続くか")
    fmt.Println(bits.LeadingZeros8(variable))

    fmt.Println("8bit で表した場合に、右からいくつ 0 が続くか")
    fmt.Println(bits.TrailingZeros8(variable))

    fmt.Println("引数の数を表現するのに必要な最小 bit数")
    fmt.Println(bits.Len8(variable))

    fmt.Println("８bit 表した場合に立っている bit数")
    fmt.Println(bits.OnesCount8(variable))

    fmt.Println("８bit 表した場合に bit の並びを逆順にする")
    fmt.Println(bits.Reverse8(variable))
}
```

なんと便利な bit演算子！

その他詳細はこちらのサイトが解りやすかったです。

[Go1.9で実装されたmath/bitsパッケージの関数一覧とその使用例](https://qiita.com/cia_rana/items/2df8fb14517ab74af4c7)

## Golang で暗号化処理を書く
では byte の扱いがわかったところで、暗号化処理を書いてみましょう。
今回は組み込まれている暗号化アルゴリズムから下記を実装してみましょう。

- パスワード認証
    - AESアルゴリズム で暗号化した物を Hash化する
    - Hash化された文字列 をAESアルゴリズム で復号化する

### 暗号化処理基本
暗号化処理においては、次の実装が必要になります。

- 暗号化アルゴリズムの指定
- ブロックモード指定
- byte のパディング指定
- byte のトリミング処理

順を追って実装していきましょう。
ライブラリは組み込まれている crypto/aes パッケージを利用するため、ブロック暗号のモード指定や、
パディング、トリムの便利な関数が用意されていますが、その他の暗号化アルゴリズムを利用する場合は、
自身で実装を行う必要があります。

※と、思っていたらその他暗号化アルゴリズムでも自身で実装しなくても、便利に使えるようになっている！？

そこで、今回は自身でそれらの実装も行ってしまいましょう。
ブロックモードはセキュアではありませんが、理解しやすい ECBモードを使います。


#### 構造体作成

```go
package main

import (
    "fmt"
)

// 暗号用構造体を定義
type Auth struct {
    SecretKey []byte
}

// 暗号化処理用のコンストラクタを定義
func NewAuth() *Auth {
    SECRET_KEY := []byte("secretGosecretGo")

    return &Auth{SecretKey: SECRET_KEY}
}

func main() {
    auth := NewAuth();
    // auth の中身
    fmt.Printf("%#v\n", auth)
}
```

これが暗号化処理を行う基本的な構造体定義と、構造体のポインタを返すコンストラクタになります。

#### パディング、トリミング処理実装

```go
// PKCS#5 方式のパディング処理 
// 暗号化を行う為に BlockSize の倍数の byte にする必要があるため、byte数を調整する
func (a *Auth) pKCS5Padding(ciphertext []byte, blockSize int) []byte {
    // パディングが必要な byte数 を算出する
    padding := blockSize - len(ciphertext)%blockSize

    // 必要な byte数分の スライス(配列)を生成
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)

    // パディングした物を返却する
    return append(ciphertext, padtext...)
}

// PKCS#5 方式のトリミング処理
// パディングしたままでは、復号に誤りが生じるため、追加した分の文字列を取り除く
func (a *Auth) pKCS5Trimming(encrypt []byte) []byte {
    // パディングした byte数 を取得する (空文字を除く: -1 がポイントでハマりどころ)
    padding := encrypt[len(encrypt)-1]

    // トリミングした物を返却する
    return encrypt[:len(encrypt)-int(padding)]
}
```

#### 暗号化と復号化

```go
func (a *Auth) encrypt(targetString string) []byte {
    // AES暗号のオブジェクトを取得する
    cipher, _ := aes.NewCipher(a.SecretKey)

    // 暗号化アルゴリズム毎に定義される BlockSize を取得する
    blockSize := cipher.BlockSize()

    // BlockSize が合うようにパディングを行う
    content := a.pKCS5Padding([]byte(targetString), blockSize)

    // 暗号化された文字列格納用の []byte型変数を用意する    
    result := make([]byte, len(content))

    // ECBモードで暗号化処理を行う
    for i := 0; i < len(content)/blockSize; i++ {
        currentIndex := blockSize * i
        // 文字列を暗号化していく
        cipher.Encrypt(result[currentIndex:], content[currentIndex:])
    }

    return result
}

func (a *Auth) decrypt(cryptedPassword []byte) string {
    // AES暗号のオブジェクトを取得する
    cipher, _ := aes.NewCipher(a.SecretKey)

    // 暗号化アルゴリズム毎に定義される BlockSize を取得する
    blockSize := cipher.BlockSize()

    // 復号化された文字列格納用の []byte型変数を用意する    
    result := make([]byte, len(cryptedPassword))

    // ECBモード で復号化処理を行う
    for i := 0; i < len(cryptedPassword)/blockSize; i++ {
        currentIndex := blockSize * i
        cipher.Decrypt(result[currentIndex:], cryptedPassword[currentIndex:])
    }

    // パディングされた byte を取り除いた物を返す
    return string(a.pKCS5Trimming(result))
```

これで基本的な処理を実装しました。
動かしてみましょう。

import文を追加

```go
    import (
        "fmt"
        "bytes"
        "encoding/base64"
        "crypto/aes"
    )
```

main関数 に処理を追加

```go
func main(){
    auth := NewAuth();
    // auth の中身
    fmt.Printf("%#v\n", auth)
    // 暗号化を行う文字列
    target := "test"

    fmt.Println([]byte(target))

    // 暗号化を行う
    encrypted := auth.encrypt(target);

    // 暗号化された文字列
    fmt.Println(string(encrypted))

    // 暗号化された []byte型を base64 エンコードして文字列に変換して返す
    encryptedBase64Encoded := base64.StdEncoding.EncodeToString(encrypted)

    // Base64 エンコード された文字列
    fmt.Println(encryptedBase64Encoded)

    // Base64 デコードして []byte型 にする
    encryptedBase64Decoded, _ := base64.StdEncoding.DecodeString(encryptedBase64Encoded)

    // Base64 デコード された文字列
    fmt.Println(encryptedBase64Decoded)

    // 復号化を行う
    decrypted := auth.decrypt(encryptedBase64Decoded)

    // 復号化された文字列
    fmt.Println(decrypted)
}
```

## Golang で暗号化処理を書く まとめ
暗号化・復号化出来たでしょうか？

Golang は処理速度重視で設計された言語であり、
byte の扱いが比較的得意な言語です。
速度を重視する場面においてはより最適な処理を書く事が可能です。

最後に暗号化・復号化処理の全文を記述します。

```go
package main

import (
    "fmt"
    "bytes"
    "encoding/base64"
    "crypto/aes"
)

// 暗号用構造体を定義
type Auth struct {
    SecretKey []byte
}

// 暗号化処理用のコンストラクタを定義
func NewAuth() *Auth {
    SECRET_KEY := []byte("secretGosecretGo")

    return &Auth{SecretKey: SECRET_KEY}
}

// PKCS#5 方式のパディング処理 
// 暗号化を行う為に BlockSize の倍数の byte にする必要があるため、byte数を調整する
func (a *Auth) pKCS5Padding(ciphertext []byte, blockSize int) []byte {
    // パディングが必要な byte数 を算出する
    padding := blockSize - len(ciphertext)%blockSize

    // 必要な byte数分の スライス(配列)を生成
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)

    // パディングした物を返却する
    return append(ciphertext, padtext...)
}

// PKCS#5 方式のトリミング処理
// パディングしたままでは、復号に誤りが生じるため、追加した分の文字列を取り除く
func (a *Auth) pKCS5Trimming(encrypt []byte) []byte {
    // パディングした byte数 を取得する (空文字を除く: -1 がポイントでハマりどころ)
    padding := encrypt[len(encrypt)-1]

    // トリミングした物を返却する
    return encrypt[:len(encrypt)-int(padding)]
}

func (a *Auth) encrypt(targetString string) []byte {
    // AES暗号のオブジェクトを取得する
    cipher, _ := aes.NewCipher(a.SecretKey)

    // 暗号化アルゴリズム毎に定義される BlockSize を取得する
    blockSize := cipher.BlockSize()

    // BlockSize が合うようにパディングを行う
    content := a.pKCS5Padding([]byte(targetString), blockSize)

    // 暗号化された文字列格納用の []byte型変数を用意する    
    result := make([]byte, len(content))

    // ECBモードで暗号化処理を行う
    for i := 0; i < len(content)/blockSize; i++ {
        currentIndex := blockSize * i
        // 文字列を暗号化していく
        cipher.Encrypt(result[currentIndex:], content[currentIndex:])
    }

    return result
}

func (a *Auth) decrypt(cryptedPassword []byte) string {
    // AES暗号のオブジェクトを取得する
    cipher, _ := aes.NewCipher(a.SecretKey)

    // 暗号化アルゴリズム毎に定義される BlockSize を取得する
    blockSize := cipher.BlockSize()

    // 復号化された文字列格納用の []byte型変数を用意する    
    result := make([]byte, len(cryptedPassword))

    // ECBモード で復号化処理を行う
    for i := 0; i < len(cryptedPassword)/blockSize; i++ {
        currentIndex := blockSize * i
        cipher.Decrypt(result[currentIndex:], cryptedPassword[currentIndex:])
    }

    // パディングされた byte を取り除いた物を返す
    return string(a.pKCS5Trimming(result))
}

func main(){
    auth := NewAuth();
    // auth の中身
    fmt.Printf("%#v\n", auth)
    // 暗号化を行う文字列
    target := "test"

    fmt.Println([]byte(target))

    // 暗号化を行う
    encrypted := auth.encrypt(target);

    // 暗号化された文字列
    fmt.Println(string(encrypted))

    // 暗号化された []byte型を base64 エンコードして文字列に変換して返す
    encryptedBase64Encoded := base64.StdEncoding.EncodeToString(encrypted)

    // Base64 エンコード された文字列
    fmt.Println(encryptedBase64Encoded)

    // Base64 デコードして []byte型 にする
    encryptedBase64Decoded, _ := base64.StdEncoding.DecodeString(encryptedBase64Encoded)

    // Base64 デコード された文字列
    fmt.Println(encryptedBase64Decoded)

    // 復号化を行う
    decrypted := auth.decrypt(encryptedBase64Decoded)

    // 復号化された文字列
    fmt.Println(decrypted)
}
```

