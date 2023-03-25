# Hugo to WordPress Migration Tool

このリポジトリは、Hugoで作成された記事をWordPressに記事を移行するためのGo言語製ツールです。

# ChatGPT(GPT-4)を使って書いています。

このリポジトリのコードはすべてChatGPTにコーディングさせるというルールを用いて開発しています。
最初に与えた条件は以下のとおりです。

```
HugoのmarkdownをパースしてWordPressのREST APIを使って記事を移行するプログラムを書いてください。

- 言語はgolangを使うこと
- 画像の移行も対応すること
- Hugoで定義されているカテゴリをWordPressに移行できること
- Hugoで定義されているTagもWordPressに移行できること
- WordPress側ではREST APIを利用すること
- Golangファイルが分割されていること
- Golangテストが書かれていること
- 分割されたファイルは、設定、Hugo、WordPressでそれぞれパッケージがわかれていること
- パッケージ名はgithub.com/polidog/hugo-to-wordpressでお願いします。
- 設定ファイルはyamlになっており、wordpress, hugoという形に構造化されていること
```