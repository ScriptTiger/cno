[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://docs.google.com/forms/d/e/1FAIpQLSfBEe5B_zo69OBk19l3hzvBmz3cOV6ol1ufjh0ER1q3-xd2Rg/viewform)

**DISCLAIMER!!!: THIS MODULE IS STILL IN ITS EARLY DEVELOPMENT, SO USE AT YOUR OWN RISK!**

# cno
cno is a VERY minimal syscall module (no CGO, no C, CNO!) with a heavy focus on GUI development for VERY minimal GUIs, generally small utilities which have historically been CLI and where having a GUI would be nice, but not necessary, and it can just be bolted on without much additional bloat. So, it's somewhat special-purpose and not intended for more advanced general-purpose GUI development.

The base `cno` package itself contains cross-platform utility functions, while the `cno/win` package contains Windows-specific utility functions. And then the `cno/win/gui` package contains constants, structures, and functions specific to Windows GUI development.

To import `cno`, `cno/win`, and/or `cno/win/gui` into your project:  
`go get github.com/ScriptTiger/cno`, `go get github.com/ScriptTiger/cno/win`, and/or `go get github.com/ScriptTiger/cno/win/gui`  
Then just `import "github.com/ScriptTiger/cno"`, `import "github.com/ScriptTiger/cno/win"`, and/or `import "github.com/ScriptTiger/cno/win/gui"` and get started with using its functions.

In the case of importing `cno/win/gui`, which contains many constants and structures in addition to the Windows API foreign functions and other functions, it's often more convenient to use a dot import in order to make all of those constants, structures, and functions available without having to use the `gui.` prefix, especially if you are more familiar with Windows GUI development in the context of other languages:  
'import . "github.com/ScriptTiger/cno/win/gui"'

Please refer to the dev package docs and reference implementation for more details and ideas on how to integrate cno into your project.  

Dev package docs:  
https://pkg.go.dev/github.com/ScriptTiger/cno  
https://pkg.go.dev/github.com/ScriptTiger/cno/win  
https://pkg.go.dev/github.com/ScriptTiger/cno/win/gui

Reference implementation:  
https://github.com/ScriptTiger/cno/blob/main/ref

# Other projects using cno

kanziSFX:  
https://github.com/ScriptTiger/kanziSFX

# More About ScriptTiger

For more ScriptTiger scripts and goodies, check out ScriptTiger's GitHub Pages website:  
https://scripttiger.github.io/
