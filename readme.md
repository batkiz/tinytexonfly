# TinyTeX on the fly

[![test](https://github.com/batkiz/tinytexonfly/actions/workflows/test.yml/badge.svg)](https://github.com/batkiz/tinytexonfly/actions/workflows/test.yml)
[![goreleaser](https://github.com/batkiz/tinytexonfly/actions/workflows/release.yml/badge.svg)](https://github.com/batkiz/tinytexonfly/actions/workflows/release.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/batkiz/tinytexonfly.svg)](https://pkg.go.dev/github.com/batkiz/tinytexonfly)
[![GitHub license](https://img.shields.io/github/license/batkiz/tinytexonfly)](https://github.com/batkiz/tinytexonfly/blob/main/LICENSE)

用途是自动化 TinyTeX 的装包过程（虽然现在主要在用 tectonic）。

实现非常脏，所以并不是所有的包都会被识别到，还是需要看错误日志。

## 使用

通常情况下，你可以使用下述命令
```shell
tinytexonfly dir # 默认为 "."
```

`tinytexonfly` 会默认递归搜索指定的文件夹（无输入时为当前文件夹）下所有的 `tex, dtx, sty, cls` 文件，并输出需要执行的命令。

也可以指定特定文件，如
```shell
tinytexonfly <file>
```

如此 `tinytexonfly` 便只会处理特定文件。


当然，正如前文所述，`tinytexonfly` 的实现很脏，所以会有些包处理不到，此时你可以根据错误日志的输出，通过
```shell
tinytexonfly search foo.sty
tinytexonfly s foo.sty # alias
```
查询，然后按需安装。

比如下面这个例子：

错误日志：
```text
Package fontspec Error: The font "XITSMath-Regular" cannot be found.
```

搜索
```shell
tinytexonfly s XITSMath-Regular
```

输出
```text
fonts/opentype/public/xits/XITSMath-Regular.otf
```

此时只需运行 `tlmgr install xits` 就行了。

## 数据来源

texlive files data 来自 [clearlinux-pkgs/texlive](https://github.com/clearlinux-pkgs/texlive)

具体文件为 [texlive.spec](https://raw.githubusercontent.com/clearlinux-pkgs/texlive/master/texlive.spec)

## 灵感来自
- [tectonic#717](https://github.com/tectonic-typesetting/tectonic/issues/717#issuecomment-757340814)
- [jpfairbanks/tlmgrlookup](https://github.com/jpfairbanks/tlmgrlookup)

## LICENSE

AGPLv3