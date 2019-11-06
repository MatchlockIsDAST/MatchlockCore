# MatchlockCore
```
   __  ___     __      __   __         __     ====
  /  |/  /__ _/ /_____/ /  / /__  ____/ /__  ====   Matchlock
 / /|_/ / _ '/ __/ __/ _ \/ / _ \/ __/  '_/ ====	〔 https://github.com/a-zara-n/Matchlock 〕
/_/  /_/\_,_/\__/\__/_//_/_/\___/\__/_/\_\ ====
=============================================
```
MatchlockのCore code
## About
Matchlockとは**開発者に寄り添ったセキュアな開発サイクルを実現するサポートツール**です。

## Description
このツールを使う際は、お好きなChromeかFirefox(こっちがオススメ)とSQLiteをインストールしておいてください。
## Install & Start UP
**Install**
```sh
# When putting under GOPATH
mkdir $GOPATH/src/github.com/a-zara-n/
cd $GOPATH/src/github.com/a-zara-n/
#  If you put in that directory from here
git clone https://github.com/a-zara-n/Matchlock.git
go mod download
```
**Start UP by Server Run**
```sh
go run .
# or
go build .
./Matchlock
   __  ___     __      __   __         __     ====
  /  |/  /__ _/ /_____/ /  / /__  ____/ /__  ====   Matchlock
 / /|_/ / _ '/ __/ __/ _ \/ / _ \/ __/  '_/ ====	〔 https://github.com/a-zara-n/Matchlock 〕
/_/  /_/\_,_/\__/\__/_//_/_/\___/\__/_/\_\ ====
=============================================

⇨ http server started on [::]:8888
```
**Start UP by Web UI**

次のようなURLを利用してサイトにアクセスしてください
```
http(s)?://<hostname>:8888
example : http://localhost:8888
```

## Special thanks
　このMatchlockを作成するにあたり、国立研究開発法人情報通信研究機構(以後 NICT とする)が実施する，セキュリティイノベーター育成プログラム[SecHack365](https://sechack365.nict.go.jp/)での成果物であります。
 
  また、このツールを作成するにあたり、先人のエンジニアから意見や知識をいただき、機能面や設計面での飛躍ができたと考えています。
  
  この場を借りてNICTを始めこのSecHack365に関わる全ての方々、そして協力いただいた方々に謝辞を述べたいと思います。
