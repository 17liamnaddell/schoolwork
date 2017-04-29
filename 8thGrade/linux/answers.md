#Test 1
1. Which of the following is not a function of the linux kernel?
  1. Allocation memory for use by programs
  2. Allocation CPU time for use by programs
  3. creating menus in GUI programs
  4. Controlling access to the hard disk
  5. enabling programs to use a network

>C Thats just plain stupid

2. Which of the following is an example of an embedded Linux OS?
  1. Android
  2. SUSE
  3. CentOS
  4. Debian
  5. Fedora
  
>A It relies on linux for the OS, then builds a GUI on top of it

----

##Beyond the Essentials.

3. Which of the following is a notable difference between Linux and Mac OSX?
  1. Linux can run common GNU programs, whereas OSX cannot.
  2. Linux's GUI is based on the X Window System, whereas OSX's is not
  3. Linux cannot run on Apple Macintosh hardware, whereas OSX can run only on Apple hardware
  4. Linux relies heavily on BSD software, whereas OSX uses no BSD software
  5. Linux supports text-mode commands, but OSX is a GUI-only os
  
>B, never knew that before :)

4. True or false: The linux kernel is derived from the BSD kernel

>false

5. True or false: If you log into a Linux system in graphical mode, you cannot use text-mode commands in that session

>false

6. True or false: CentOS is a Linux distro with a long release cycle.

>true

7. A linux text-mode login prompt reads         (one word)

>login:

8. A common security problem with Windows that's essentially nonexistent on a Linux is            .

>Viruses

9. Pre-release software that's likely to contain bugs is known as             and       .

>beta, alpha


##Score
100%



#Test 2

1. What type of multitasking does linux use?
  1. Preemptive
  1. Multi-user
  1. Co-operative
  1. Single-tasking
  1. Single-user

>A

2. Which of the following is a characteristic of all open source software?
  1. The software cannot be sold for a profit; it must be distributed free of charge
  1. It must be distributed with both source code and binaries
  1. Users are permitted to redistribute altered versions of the original software
  1. The software was originally written at a college or university
  1. The software must be written in an interpreted language that requires no compilation

>C

3. Which of the following programs is most likely to be installed and regularly used on a desktop computer that runs linux
  1. Apache
  1. Postfix
  1. Android
  1. Evolution
  1. BIND

>D

4. T/F: VMS was a common OS on x86 PCs at the time Linux was created

>false

5. T/F: Some DVRs run Linux

>T

6. T/F: A linux computer being used as a server generally does not require X

>T

7. a Linux uses a          kernel design, as contrasted with a microkernel design

>Monolithic

8. A type of software that's destributed for free but that requres payment on the "honor system" if a person uses it is called          

>Shareware

9. A           computer is likely to run a word processor and a Web browser

>Desktop

##Score
100%



#Chapter 3

1. Which of the following is not required in order for software to be certified as open source?
  1. The license must not discriminate against people or groups of people
  1. The license must not require that the software be distributed as part of a specific product
  1. The license must require that changes be distributed under the same license
  1. The program must come with source code, or the author must make it readily available on the internet
  1. The license must automatically apply to anybody who acquires the software

>c

1. Which is true of Linux distributions as a whole?
  1. They're covered by the GPL or BSD license, depending on the distribution
  1. Sometimes, they may not be copied because of non-open source software
  1. They may be copied only after software using the MIT license is removed
  1. They all completely conform to the principals of the open source movement.
  1. They all qualify as free software, as the FSF uses the term

>a      WRONG

1. Which of the following is a key part of the FSF's philosophy?
  1. Developers should use the latest version of the FSF's GPL
  1. Users should have the right to modify free software and distribute it under a commercial license
  1. Developers should write software only for free operating systems such as GNU/linux
  1. Users should engage in civil disobedience by copying proprietary software
  1. Users must have the right to use software as they see fit

>e

1. T/F: Copyright law governs the distribution of software in most countries

>T

1. T/F: The FSF's free software definition and the OSI's ten principals of open source software both require that users have the ability to examine a program's workings--that is, its source code

>T

1. T/F: Because their hardware designs are proprietary, hardware vendors cannot release open source drivers for their products

>F

1. A license created by the FSF and often used for libraries is the         .

>LGPL

1. An organization devoted to promoting open source-like principles in fields such as video and audio recordings is the      .

>Creative commons

1. The FSF's general principles are summarized by the term          , which refers to using copyright laws for purposes that are in some ways contrary to the copyright's original intent. 

>copyleft


##Score 
85%

#Chapter 4 

1. Which of the following are Linux desktop environments?
  1. GTX+
  2. GNOME
  3. KDE
  4. Evolution
  5. Xfce

>2, 3, 5

2. If you want to enable one Linux computer to access files stored on another Linux computer's hard disk, which of the following network protocols is the best choice?
  1. SMTP
  2. NFS
  3. PHP
  4. DNS
  5. DHCP

>NFS

3. Which of the following languages was most of the linux kernel written?
  1. Bash shell script
  2. Java
  3. C
  4. C++
  5. Perl

>C
  
4. T/F OpenOffice.org and LibreOffice are very similar office suites

>T, they are similar

5. T/F Servers can be disrupted by malicious outsiders even if the computer that runs them is never broken into

>Yes, through DDOS or DOS attacks

6. T/F Python is generally implemented as an interpreted language

>T, it is mostly interpreted

7. Thunderbird is a         program. (Specify the general category of the software)

>Mail

8. A Linux server that handles the SMB/CIFS protocol normally runs the           software

>Samba

9. A program written in a          programming language is completely converted to binary form before being run

>Compiled


##Score
100%



#Chapter 5

1. Which of the following commands provides the most information about your motherboard's features
  1. lscpu
  2. Xorg -configure
  3. fdisk -l /dev/sda
  4. lspci
  5. http://localhost:631

> lspci

2. Why might you want to partition a hard disk(select all that apply)
  1. To install more than one OS on the disk
  2. To use the ext4fs rather than ReiserFS
  3. To turn a PATA disk into an SATA disk
  4. To seperate filesystem data from swap space
  5. To seperate the disk's cache from it's main data

>1, 4

3. Which of the following devices is not commonly attached via USB?
  1. Video monitors
  2. Keyboards
  3. External hard disks
  4. Printers
  5. Scanners

>External hard disks, WRONG

4. T/F: An EM64T CPU is capable of running a Linux distribution identified as being AMD64 CPU 

>T, they are just different names for the same CPU

5. T/F: UDF is a good filesystem to use for a Linux installation on a hard disk 

>F, thats for CD's, stupid ^-^

6. T/F: The Linux kernel includes drivers for various disk controllers, network adapters, and USB interfaces, among other things

>T

7. The x86 CPU uses a     -bit architecture

>32

8. A computer power supply converts electricity from alternating current to         (two words)

>DC, or Direct Current

9. The         standard is a modern video interface that's commonly used on computer monitors

>HDMI, the book is dated


##Score
 One was wrong

#Chapter 6

1. What keystroke moves the cursor to the start of the lime when typing a command in bash
  1. Ctrl+A
  1. Left Arrow
  1. Ctrl+T
  1. Up arrow
  1. Ctrl+E

>Ctrl+A

2. How can you run a program in the background when launching it from a shell? (Select all that apply)
  1. Launch a program by typing `start command`, where `command` is the command you want to run
  1. Launch the program by typing `bg command`, where command is the command you want to run
  1. Append an ampersand (&) to the end of the command line
  1. Launch the program normally, type Ctrl+Z in the shell, and then type `bg` in the shell
  1. Launch the program normally, type Ctrl+Z in the shell, and then type `fg` in the shell

>3 and 4

3. Which of the following commands, typed at a Bash prompt, returns you to your home directory?
  1. home
  1. cd /home
  1. cd homedir
  1. homedir
  1. cd ~

>cd ~

4. T/F: The Alt+F2 keystroke, typed in X, brings up a text-mode display you can use to log into Linux

>T

5. T/F: The filename ..\upone.txt refers to the file upone.txt in the parent of the current directory

>Wrong slash, but would if normal it would so T

6. T/F: The -r option to `ls` creates a recursive directory listing

>T

7. T/F: The         command displays the path to the current working directory

>pwd

8. To view all files, including hidden files and directories, in the current directory, you would type ls `options`

>ls -a

9. The          command displays text files or can concatenate multiple files together

>cat


##Score 
1.5 wrong


#Chapter 7

1. Which of the following commands would you type to rename newfile.txt to afile.txt
  1. mv newfile.txt afile.txt
  1. cp newfile.txt afile.txt
  1. ln newfile.txt afile.txt
  1. touch newfile.txt afile.txt

>1

2. You want to copy a directory, MyFiles, to a USB flashdrive that uses the FAT filesystem. The contents of MyFiles are as follows `contract.pdf` `outline.pdf`and `Outline.PDF` The USB flash drive is mounted at /media/usb, and so you type cp -a MyFiles/ /media/usb. What problem will occur when you attempt to copy these files?
  1. The command will fail because it tries to create links
  1. The MyFiles directory will be copied, but none of its files will be copied
  1. One of the files will be missing on the USB flash drive
  1. One file's name will be changed during the copy
  1. Everything will be fine; the command will work properly

>C

3. You type mkdir one/two/three and receive an error message that reads in part: `No such file or directory` What can you do to overcome this problem? (select all that apply)
  1. Add the --parents flag 
  1. Issue three seperate mkdir commands: mkdir one, then mkdir one/two, then mkdir one/two/three
  1. Type touch /bin/mkdir to be sure the mkdir program file exists
  1. Type rmdir one to clear away the interfering base of the desired new directory tree
  1. Type rm -r one to clear away the entire interfearing directory tree 

>2.

4. T/F: You can create a symbolic link from one low-level FS to another

>T

5. T/F: You can easily damage your Linux installation by mistyping an rm command as a regular account

>F

6. T/F: You can set a directory's time stamp with the touch command

>T

7. You want to copy a file (origfile.txt) to the backups directory, but if a file called origfile.txt exists in the backups directory, you want to go ahead with the copy only if the file in the source location is newer than the one in the backups folder. The command to do this is cp     origfile.txt backups/

>-u

8. You have typed rmdir junk to delete the junk directory, but this command has failed because junk contains word proccessing files. What command might you type to do this job?

>rm -rf or rm -r

9. Which wildcard character matches any one symbol in a filename?

>?

##Score
100%



#Chapter 8

1. Which of the following commands is an improved version of more?
  1. grep
  1. html
  1. cat
  1. less
  1. man

>4, less is more


1. Which of the following statements is a fair comparison of man pages to HOWTO documents?
  1. man pages require internet access to read; HOWTOs do not.
  1. man pages are a type of printed cocumentation; HOWTOs are electronic
  1. man pages describe software from a user's point of view; HOWTOs are programmer's document
  1. man pages are brief reference documents; HOWTOs are more tutorial in nature
  1. man pages use a hyperlinked format, whereas HOWTOs do not. 

>4

1. A user types `whatis less`. 
  1. A short one-paragraph description of the purpose of the `less` command
  1. The complete path to the less command in the Linux filesystem
  1. Summary information from `man` pages whose Name sectioons mention `less`
  1. The complete man page for less, which you would then scroll through with your terminal
  1. The URLs for Web sites with information on the less command

>3

1. T/F: You can force man to display a man page in a specific section of the manual by preceding the search name with the section number, as in `man 5 password`

>T

1. T/F: info pages are a Web-based documentation format 

>F

1. T/F: Linux documentation in the /usr/share/doc directory tree is almost always in OpenDocument Text format

>F

1. File formats are described in man section         .

>5

1. Each document in an info page is known as a           .

>node

1. The              command searches a database of filenames, enabling you to quickly identify files whose names match a term you specify

>appropos


##Score 8/9

#Chapter 9

1. Which of the following tools is best suited to installing a software package and all its dependencies on a Debian computer
  1. yum
  2. zypper
  3. dmesg
  4. rpm
  5. (apt-get)
2. What is the usual name of the first process that the Linux kernel runs, aside from itself?
  1. (init)
  2. bash
  3. cron
  4. login
  5. grub
3. Where do most log files reside on a Linux computer?
  1. (/var/log)
  2. /etc/logging
  3. /usr/log
  4. /home/logging
  5. /log/usr
4. (T): When using suitable commands, you can normally install a program and be sure that all the software on which it depends will also be installed, provided you have an Internet connection
5. (T): By default, the first process listed in top is currently consuming the most CPU time
6. (T): The dmesg command may produce different output after a computer has been running for weeks than when it first started.
7. Most linux distros maintain info on what packages are installed in the (package database). (Two words)
8. You're using bash, then you type emacs to launch the emacs editor. In this case, emacs is bash's (child).
9. General system messages are likely to be found in /var/log/messages or /var/log/syslog, depending on your distro

##Score %100

#Chapter 10

1. Which of the following commands will print lines from the file world.txt that contain matches to changes and changed?
  1. `grep change[ds] world.txt`    <-correct
  1. `tar change[d-s] world.txt`
  1. `find "change'd|s'" world.txt`
  1. `cat world.txt changes changed`
  1. `find change[^ds] world.txt`
1. Which of the following redirection operators appends a program's standard output to an existing file, without rewriting that file's original contents?
  1. |
  2. 2>
  3. &>
  4. >
  5. >>    <-correct
1. You've received a tarball called data79.tar from a colleague, but you want to check the names of the files it contains before extracting them. Which of the following commands would you use to do this?
  1. tar uvf data79.tar
  1. tar cvf data79.tar
  1. tar xvf data79.tar
  1. tar tvf data79.tar    <-correct
  1. tar Avf data79.tar
1. T: The regular expression `Linu[^x].*lds` mathes the string "Linus Torvalds".
1. T: the find command enables you to locate files based on their sizes
1. F: To compress files archived with zip, you must use an external compression program such as gzip or bzip2 in a pipeline with zip
1. The character that represents the start of a line in a regex is `^`  .
1. Complete the following command to redirect both stdout and stdin from the bigprog program to the file out.txt
  1. `$ bigprog &> out.txt`
1. The gzip, bzip2, and xz programs all perform lossless compression, in which the decompressed data exactly matches the origional pre-compession data


##Score 100%



#Chapter 11

1. Which type of file is nano least likely to be useful for examining or editing?
  1. /var/log/messages
  2. a .html file
  3.  an e-mail message saved from an e-mail client
  4. a libreOffice word processing document
  5. /etc/X11/xorg.conf

>its an useless editor, but the correct answer is #4

1. Which keystrokes invoke the pico or nano search function? (select all that apply.)
  1. F3
  2. F6
  3. Esc-S
  4. Ctrl+f
  5. Ctrl+W

> F6, Ctrl+w
1. How would you remove two lines of text from a file using Vi?
  1. In command mode, position the cursor on the first line and type 2dd
  1. In command mode, position the cursor on the last line and type 2yy
  1. In insert mode, position the cursor at the start of the first line and press Ctrl+k twice
  1. Select the text with the mouse and then select File>delete from the menu

>2dd

1. T/F: Unicode is useful for encoding most European languages but not Asian languages

>F

5. T/F GUI text editors for ASCII are superior to text-mode ASCII text editors because the GUI editors support underlining italics and multiple fonts.

>Opinion piece, but probably T(depending if the TEXT editor supports underlining) bad question

1. T/F: Many but not all config files use a hash mark to identify comment lines.

>T

1. ASCII supports     unique characters

>128\*2 (depending if u use 8 bits instead of 7)

1. Three keystrokes that can initiate a search and replace operation in nano are F14,     , and     .\

>Esc+6

1. While in Vi's command mode, you can type     to undo a change


> :u in ex mode, and u in command mode


##Score
88%


#Chapter 12
1. After using a text editor to create a shell script what step should you take before trying to use the script by typing it's name?
  1. Set one or more executable bits using chmod            <-correct
  1. Copy out the script to the /usr/bin/scripts directory
  1. Compile the script by typing bash scriptname where scriptname is the script's name
  1. Run a virus checker on the script to be sure it contains no viruses
  1. Run a spell checker on the script to ensure it contains no bugs
2. Describe the effect of the following short script, cp1, if it's called as cp1 big.c big.cc
  ```
  #!/bin/bash
  cp $2 $1
  ```
  1. It has the same effect as the cp command
  1. It compiles the C program big.c and calls the result big.cc
  1. It copies the contents of big.cc to big.c eliminating the old big.c            <-correct
  1. It converts the C program big.c into a C++ program called big.cc.
  1. The script's first line is invalid so it won't work.
3. What is the purpose of conditional expressions in a shell script?
  1. They prevent scripts from executing if license conditions aren't met
  1. They display info about the script's computer environment
  1. They enable the script to take different actions in response to variable data            <-correct
  1. They enable scripts to learn in a manner reminiscent of Pavlovian conditions
  1. They cause scripts to only run at a time of day
1. F: A user types myscript laser.txt to run a script called myscript. Within myscript the $0 variable represents holds the value laser.txt
1. F: Valid looping statements in Bash include for, while, and until.
1. F: The following script launches three simultaneous instances of the terminal program.
  ```
  #!/bin/bash
  terminal
  terminal
  terminal
  ```
1. You have written a simple shell script that does nothing but launch programs In order to ensure that the script works with most user shells, what should its first line read?
/bin/sh
1. What command can you use to display prompts for a user in a shell script?
read
1. What bash scripting command can you use to control the program flow based on a variable that can take many values (such as all the letters of the alphabet)?
? bad question, case?


##score 
8.7/9


# Chapter 13

1. What is the purpose of the system account with a UID of 0?
  1. It's the system admin's account                <-----correct
  2. It's the account for the first ordinary user.
  3. Nothing; UID 0 is left intentionally undefined
  4. It varies from one distro to another
  5. It's a low privilege account thats used as a default by some servers
2. What type of info will you find in /etc/passwd for ordinary user accounts? (select all that apply)
  1. A uid number             <-----correct
  2. A complete listing of every group to which the user belongs
  3. The path to the account's home directory              <-----correct
  4. The path to the account's default GUI desktop enviornment
  5. The path to the account's default text-mode shell             <-----correct
3. You want to run the command iptables -L as root, but youre logged in as an ordinary user. Which of the following commands will do the job, assuming the system is configured to give you root access via the appropriate command?
  1. sudo iptables -L             <-----correct
  2. root iptables -L
  3. passwd iptables -L
  4. su iptables -L
  5. admin iptables -L
4. F: whoami provides more info than id.
5. T: Linux stores info on its groups in the /etc/groups file.
6. T: as a general rule, you should employ extra care when running programs as root
7. The file that associates usernames with UID numbers in Linux is /etc/passwd
8. To learn who is currently logged into the computer and what programs there currently running, you can type   w  (one word)
9. UIDs below 500 or 1000 are reserved for use by program accounts


#Chapter 15

1. What command would you type (as root) to change the ownership of somefile.txt from ralph to tony?
  1. chown ralph:tony somefile.txt
  2. chmod somefile.txt tony
  3. chown somefile.txt tony
  4. chown tony somefile.txt                  <-Correct
  5. chmod tony somefile.txt
2. Typing ls -ld wonderjaye reveals a symbolic file mode of drwxr-xr-x
  1. wonderjaye is a symbolic link.
  2. wonderjaye is an executable program.
  3. wonderjaye is a directory                  <-Correct
  4. wonderjaye may be read by all users of the system
  5. wonderjaye may be written by any member of the file's group                  <-Correct
3. Which of the following commands can you use to change a file's group?
  1. groupadd
  2. groupmod
  3. chmod
  4. ls
  5. chown                  <-Correct
4. T: A file with permissions of 755 can be read by any user on the computer, assuming all users can read the directory in which it resides
5. F: Only root may use the chmod command.
6. F: Only root may change a file's ownership with chown
7. What option causes chown to change ownership on an entire directory tree.    `-R`
8. What three character symbolic string represents read and execure permission but no write permission?  `r-x`
9. What symbolic representation can you pass to chmod to give all users execute access to a file, without affecting other permissions. `+x`

##Score 7/9



#Chapter 16

1. What types of files are you likely to find in /usr/lib, according to the FHS
  1. Liberty files
  2. Liberated files
  3. Libra files
  4. Library files     <---correct
  5. Liberal files
2. You want to discover the sizes of several dot files in a directory. Which of the following commands might you use to do this
  1. ls -la            <---correct
  2. ls -p
  3. ls -R
  4. ls -d
  5. ls -ld
3. When should programs be configured as SUID root?
  1. At all times; this permission is required for executable programs
  2. whenever a program should be able to access a device file
  3. Only when they require root privileges to do their job.          <---correct
  4. Whenever the program must be able to access an accounts user ID numb er
  5. Never, this permission is a security risk
4. T: Print spool files are stored in a subdirectory of /var.       <---probably
5. F: On a properly configured Linux system, any user can delete any file from /tmp
6. T: If you hide a file in Linux by making it a dot file, you must change any existing references to that file in config files if those references are to continue working
7. Typically optical discs and USB flash drives are mounted in subdirectories of /mnt or  /media       ?
8. Temp files that are guaranteed to not be deleted during a reboot reside in /tmp            .
9. You want to set the sticky bit on an existing directory, subdir, without otherwise altering its permissions. To do so, you should type chmod +t  subdir.

##Score 8/9


#Chapter 17

1. You want to set up a computer on a local network via a static TCP/IP configuration, but you lack a gateway address. Which of the following is true?
  1. Because the gateway address is necessary, no TCP/IP networking functions will work             <--correct
  2. TCP/IP networking will function, but youll be unable to convert hostnames to IP addresses and vice versa
  3. You'll be able to communicate with machines on your local network segment, but not with other systems
  4. The computer won't be able to tell which other computers are local and which are remote
  5. You'll be able to use the computer as a network server system, but not as a network client
2. Which of the following types of info are returned by typing ifconfig eth0? (select all that apply)
  1. The names of programs using eth0       
  2. The ip address assigned to eth0        <--correct
  3. The hardware address of eth0
  4. The hostname associated with eth0
  5. the gateway with which eth0 communicates <--correct
3. The ping utility responds normally when you use it with an ip address, but not when you use it with a hostname that you're positive corresponds to this ip address. What might cause the problem? (select all that apply)
  1. The route between your computer and its DNS server may be incorrect.
  2. The target computer may be configured to ignore packets from ping
  3. The DNS configuration on the target system may be broken <--correct
  4. Your computer's hostname may be set incorrectly
  5. Your computer's DNS configuration may be broken          <--correct
4. T: IPv4 addresses are four bytes in length
5. T: The /etc/resolv.conf file tells the computer whether to use DHCP for its network configuration 
6. T: You can check the current status of your routing table by typing route at a shell prompt
7. The iw           program serves as a multi-purpose networking tool; it can do many of the same things as ifconfig, route and several others
8. The traditional name for the first ethernet interface in linux(but not in recent versions of Fedora) is eth0      .
9. A firewall is a program or system configuration that blocks or enables network access to, from, or through a computer based on criteria you specify


##Score  5/9
