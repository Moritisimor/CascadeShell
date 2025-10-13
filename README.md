# CλSH
The Portable and User-Friendly Shell written in Go.

## What is this project about?
This project is a shell entirely written in go. It is designed to be portable, small and run anywhere, only including what it has to and leaving the rest to the OS.

## Known issues
Like with virtually any other project out there, CλSH has its issues too. These include, but are not limited to:

- Uparrow for pasting the last command into the prompt does not work. (Or rather is not implemented yet.)

## Planned features
While this project does not have a strict roadmap or something along those lines, features may always be planned for the future! Current plans include:

- Add a .cashrc 
- Add persistent memory in form of a history file


## Supported platforms
- Linux ✅ Works flawlessly
- FreeBSD ✅ Works flawlessly also
- Windows ✅ Works just fine
- Other BSDs ❓ Not tested yet, but compilable for NetBSD and OpenBSD. Feel free to test.
- MacOS/Darwin ❓ Not tested yet, but also compilable. Feel free to test.
- Android 🟨 The shell itself works fine, but Epsilon Fetch does not work. Probably because Android is too restrictive.

## Commands
Unlike with many other shells, CλSH talks a lot and will inform the user of success or failure. This is an intentional design choice and serves to help beginners or those who simply want a talkative shell. 

CλSH has a few built-in commands, these include: 

### Chdir
```Console
chdir <Destination>
```
This command will change the current directory to the target. Note that using just ```cd``` to go to the home directory will not work. Use ```gohome``` instead.

The only valid variation of this command is ```cd```

### Lookaround
```Console
lookaround
```
This command will list the entries of a specified Directory. By default, it does not show hidden directories (those starting with '.') but it can be made to do so with the ```--showhidden``` flag. Specifying no arguments will make it default to the current directory.

The only valid variation of this command is ```ls```.

### Whereami
```Console
whereami
```
This command will print the current directory as an absolute path.

The only valid variation of this command is ```pwd```

### Whoami
```Console
whoami
```
This command will print the current user to the console.

### Let
```Console
let greeting hello
```
This command will create a new variable in-memory with the specified name. To access it you put an '@' symbol before the actual name. Base variables like ```@shell```, ```@user```, ```@userID``` and ```@host``` are initialized along with the shell.

### Unlet
```Console
unlet greeting
```
This command will uninitialize a variable, deleting it from the internal map which stores variables.

The only valid variation of this command is ```free```

### Clear
```console
clear
```
This command will clear the screen and scrollbuffer.

The only valid variation of this command is ```clearscreen```

### Gohome
```Console
gohome
```
This command will take the user to their home directory.

### Print
```Console
print <Content>
```
This command will print the specified Content to the console. Can also be used to print variables by putting an '@' symbol before the variable name.

The only valid variation of this command is ```say```.

### EpsilonFetch
```Console
epsilonfetch
```
This command is a built-in Fetch program which tells you basic info about your CPU, Memory and Runtime. It does not work on Android.

The only valid variation of this command is ```epsilon```

### Exit
```Console
exit
```

This command will break the Read-Eval-Print-Loop, effectively terminating the running CλSH Process and exiting the program.

The only valid variation of this command is ```bye```.

### Running Commands in Sequence
```Console
let x = hello ; let y = world ; say @x @y
```

Semicolons ';' are used for running commands in sequence. What's important to note is that semicolons must stand completely alone, not appended to a character.

###### Note: The list of Commands is always subject to change as this project is WIP.
