# CλSH
The Portable and User-Friendly Shell written in Go.

## Ever heard of POSIX-Compliancy?
We don't care about that around here.

Ever feel like you're reading utter nonsense when looking at classical
Shell-Scripts?

Then CλSH may be the Shell for you! 

Designed to be highly user-friendly and portable, this Shell is really made for everyone. The commands are intuitive and read like plain english.

## Known issues
Like with virtually any other project out there, CλSH has its issues too. These include, but not are limited to:

- Using lookaround on directories will not work
- Readfile is not always very reliable and sometimes does not read a file.
- The ```--showhidden``` flag for lookaround does not work.

## Commands
The commands themselves are not a reinvention of the wheel, but rather an abstraction of what may otherwise appear like unreadable text to the average user.

Unlike with many other shells, CλSH Talks a lot and will inform the user of success or failure. This is an intentional design choice and serves to help the average user. 

CλSH also uses mnemonics, built-in aliases if you will, making commands have several correct versions. Most commands are also case-insensitive, but arguments may be.

Commands have no other valid mnemonics unless explicitly specified.

### Chdir
```Console
Chdir <Destination>
```
This command will change the current directory to the target, telling the user if all went right, or if something failed.

Valid Mnemonics include Changedirectory, Chdir and Changedir.

### Gohome
```Console
Gohome
```
This command will take the user to their home directory. Functionally, it is identical to Commands like
```Console
Chdir
```
or
```Console
Chdir /home/exampleuser
```

### Lookaround
```Console
Lookaround <Target>
```
This command will show the subdirectories and files of a given Target directory. If no target is given, it will use the current directory.

### Whereami
```Console
Whereami
```
This command will print where you currently are in the Filesystem as the absolute path.

### Makefile
```Console
Makefile <Target>
```
This command will create a new, empty File at the specified Target.

### Readfile
```Console
Readfile <Target>
```
This command will read the contents of the specified file and print it to the console.

### Writefile
```Console
Writefile <Target> <Content>
```
This command will write the specified content to the specified file.

### Removefile
```Console
Removefile <Target>
```
This command will remove the specified file.

### Print
```Console
Print <Content>
```
This command will print the specified Content to the console.

Valid Mnemonics include Echo, Say and Print.

### Exit
```Console
Exit
```
This command will break the Read-Eval-Print-Loop, effectively terminating the running CλSH Process and exiting the program.

###### Note: The list of Commands is always subject to change as this project is WIP.
