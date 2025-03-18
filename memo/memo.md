## メソッド

- レシーバと紐づけられた関数
  - データとそれに対する操作を紐づける

```go
type Hex int // ユーザ定義型
func (h Hex) String() string {
    return fmt.Sprintf("%x", int(h))
}

// 100をHex型として代入
var hex Hex = 100
// Stringメソッドを呼び出す
fmt.Println(hex.String())
```

## レシーバ

- メソッドに関連付けられた変数
  - メソッド呼び出し時には通常の引数と同じような扱いになる
  - ポインタを用いることでレシーバの変更を呼び出し元に伝えることができる

```go
type T int
func (t *T) f() {
    println("hi")
}
func main () {
    var v T
    (&v).f() // ★
    v.f() // ★この2つは同じ意味（？）
}
```

### ポインタレシーバ

メソッドが呼び出される際、構造体のアドレスが渡される。これにより、メソッド内で構造体のデータを直接変更できる
呼び出しに伴って構造体がコピーされない

★ 基本的にはポインタレシーバを使えばよい
参考：
https://qiita.com/atsutama/items/32a3961e1e74e20bcb14
https://qiita.com/fujita-goq/items/ed8e8730b0976c3ff3a6

ポインタレシーバと値レシーバの挙動の違い
ポインタ：変数そのものを変更する
値：変数のコピーを変更する
