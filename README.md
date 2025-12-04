[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://docs.google.com/forms/d/e/1FAIpQLSfBEe5B_zo69OBk19l3hzvBmz3cOV6ol1ufjh0ER1q3-xd2Rg/viewform)

**DISCLAIMER!!!: THIS MODULE IS STILL IN ITS EARLY DEVELOPMENT, SO USE AT YOUR OWN RISK!**

# cno
cno is a VERY minimal syscall module (no CGO, no C, CNO!) with a heavy focus on GUI development for VERY minimal GUIs, generally small utilities which have historically been CLI and where having a GUI would be nice, but not necessary, and it can just be bolted on without much additional bloat. So, it's somewhat special-purpose and not intended for more advanced general-purpose GUI development.

While GUI applications can currently be developed in cno targeting both Linux and Windows, and can be cross-compiled from any modern version of Go without the need for a C compiler, the code you develop for a Windows GUI will not be portable to a Linux GUI, nor vice versa. This is an intentional design choice. Rather than make a framework to unify the procedural nature of the Windows API with the object-oriented nature of the Linux GUI toolkits, primarily GTK, cno strives to simply bring the different paradigms under one project for the sake of simplicity while retaining their uniqueness without any additional layers of abstraction or overhead. If you're looking for a project which supports truly portable code, I'd highly recommend Fyne. On the contrary, cno strives to keep things as minimal as possible by reducing the amount of abstractions and overhead in order to produce a simple, small, and functional GUI for mostly the purposes of small utilities without introducing a new abstract paradigm as those other modules do at the cost of bloating the executables.

Please refer to the dev package docs and reference implementation for more details and ideas on how to integrate cno into your project.

Dev package docs:  
https://pkg.go.dev/github.com/ScriptTiger/cno  
https://pkg.go.dev/github.com/ScriptTiger/cno/win  
https://pkg.go.dev/github.com/ScriptTiger/cno/win/gui
https://pkg.go.dev/github.com/ScriptTiger/cno/lnx  
https://pkg.go.dev/github.com/ScriptTiger/cno/lnx/gui

Reference implementations:  
https://github.com/ScriptTiger/cno/blob/main/_ref

# Other projects using cno

kanziSFX:  
https://github.com/ScriptTiger/kanziSFX

# More About ScriptTiger

For more ScriptTiger scripts and goodies, check out ScriptTiger's GitHub Pages website:  
https://scripttiger.github.io/
