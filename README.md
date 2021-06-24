# pwalk
パイプで渡した標準入力をawkより手軽にゴニョゴニョ出来る

## HowToUse
```shell
$ cat example.txt | pwalk "echo %1"
a,b,c
b,c,a
c,a,b
$ cat example.txt | pwalk -S "," "echo %1"
a
b
c
$ cat example.txt | pwalk -S "," "echo %3 %1"
c a
a b
b c
```
