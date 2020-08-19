# gbrn

go batch rename

## build

```
$ make
```

## e.g.

```
$ls demo
test01.txt
test02.txt
test03.txt

$ make
$ cd demo
$ ../gbrn

# open vim
- test01.txt
- test02.txt
- test03.txt
+ demo01.txt
+ demo02.txt
+ demo03.txt

$ls
demo01.txt
demo02.txt
demo03.txt
```

## demo

[![asciicast](https://asciinema.org/a/ADlvzpLmNzrjXvvt8xuHCrFsi.svg)](https://asciinema.org/a/ADlvzpLmNzrjXvvt8xuHCrFsi)
