# TinyTeX on the fly

用途是自动化 TinyTeX 的装包过程（虽然现在主要在用 tectonic）。

实现非常脏，所以并不是所有的包都会被识别到，可能还需要自己去看报错找没安装的包。

## 数据来源

texlive files data 来自 [clearlinux-pkgs/texlive](https://github.com/clearlinux-pkgs/texlive)

具体文件为 [texlive.spec](https://raw.githubusercontent.com/clearlinux-pkgs/texlive/master/texlive.spec)

## 灵感来自
- [tectonic#717](https://github.com/tectonic-typesetting/tectonic/issues/717#issuecomment-757340814)
- [jpfairbanks/tlmgrlookup](https://github.com/jpfairbanks/tlmgrlookup)

## LICENSE

AGPLv3