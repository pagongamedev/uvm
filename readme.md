reference by : https://github.com/coreybutler/nvm-windows

# UVM : Unversal Version Manager

Manage multiple installations of sdk on a Windows computer.

# Supported
  ```
  uvm -d  : Dart
  uvm -f  : Flutter
  uvm -n  : NodeJS
  uvm -j :  Java           [Manual Install]
  uvm -g  : Golang
  uvm -oj : OpenJava       [Use Key]
  uvm -p  : Python         [Manual Install]
  uvm -r  : Ruby           [Manual Install]
  ```
# Usage
  ```
	uvm [-SDK] install <version> <tag> : Install SDK Version.
	uvm [-SDK] uninstall <version>     : The version must be a specific version.
	uvm [-SDK] list                    : List Version Installed and Show Current Use
	uvm [-SDK] use <version> <tag>     : Switch to use the specified version.
	uvm [-SDK] unuse                   : Disable uvm.
	uvm [-SDK] root            	       : Show Root Path
	uvm [-SDK] version                 : Displays the current running version of uvm
  ```
# Installing

## Windows
1. Extract Installer Zip to SDK Folder somewhere in your pc
   
  - Recommend Add Drive "D:" for reuse when formatted Windows like this

    ```
    D:\SDK\uvm
    ```

  - ! Not Recommend Installation In "Program File" , "Program Data" , "AppData" Because Losting When Formatted Windows and Consumed Space In your SSD Drive

2. Append "D:\SDK\uvm" and "%UVM_LINK%" to ENV:"Path" in System Variables
- ENV:"Path"

  ```
  D:\SDK\uvm
  ```
  ```
  %UVM_LINK%
  ```
    


2. open shell like CMD or Powershell with Administrator Mode (Use for Creak SymLink)
   and run this command for using
   
   ```
   $ uvm list
   $ uvm 
   ```


<b> Optional </b>

  For Quick Step for not more time for restart shell when ENV Update Please Create This

1. Create UVM_LINK in System Variables
  

- ENV : "UVM_LINK"
 
  ```
  C:\Program Files\uvm_nodejs;C:\Program Files\uvm_flutter\bin;C:\Program Files\uvm_golang\bin;C:\Program Files\uvm_dart\bin;C:\Program Files\uvm_java\bin;C:\Program Files\uvm_python;C:\Program Files\uvm_ruby\bin;
  ```

2. Create UVM_JAVA_CHANNEL in System Variables

- ENV : "UVM_JAVA_CHANNEL"
 
  ```
  Java
  ```

3. Delete or Create JAVA_HOME 

- ENV : "JAVA_HOME"
 
   ```
   C:\Program Files\uvm_java\bin
   ```


## MacOS

## Linux