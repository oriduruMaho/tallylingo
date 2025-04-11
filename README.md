# tallylingo

文字数・単語数・行数などのテキスト統計を簡単に取得できるCLIツールです．

## Description
指定されたテキストファイルと指定されたディレクトリの中の文字数・行数・単語数・バイト数をカウントします．
いくつかの出力形式をサポートします．

## Usage
```
tallylingo version 
tallylingo [CLI_MODE_OPTIONS] [FILEs...|DIRs...]
CLI_MODE_OPTIONS
  -w, --word        文字数のみカウントして標準出力します．
  -l, --line        行数のみカウントして標準出力します．
  -c, --character   単語数のみカウントして標準出力します．
  -b, --byte        バイト数のみカウントして標準出力します，
  -h, --help        ヘルプメッセージを表示します．
  -o, --output      テキストファイルに結果を出力します．
```
