## 使い方

```bash
$ dlv debug hello.go
"" dlv debug [filename]
"" dlv test -- -test.run [TestFuncName]
```

```dlv
"" (dlv) b [filename]:[line_number] ブレークポイントセット
(dlv) b hello.go:6
"" (dlv) n ステップオーバー
"" (dlv) s ステップイン
"" (dlv) restart ブレークポイントを残したままリスタート
"" (dlv) locals 変数の確認
"" (dlv) ls [func] もしくは ls [pkg].[func] ソースの確認

```
