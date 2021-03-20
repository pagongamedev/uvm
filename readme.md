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

2. Append "%UVM_LINK%" and "D:\SDK\uvm" to ENV:"Path" in System Variables

- ENV:"UVM_LINK"
 
  ```
  C:\Program Files\uvm_nodejs;C:\Program Files\uvm_flutter\bin;C:\Program Files\uvm_golang\bin;C:\Program Files\uvm_dart\bin;C:\Program Files\uvm_java\bin;C:\Program Files\uvm_python;C:\Program Files\uvm_ruby\bin;
  ```

- ENV:"Path"

  ```
  {{path}};%UVM_LINK%;D:\SDK\uvm
  ```

1. open shell like CMD or Powershell with Administrator Mode (Use for Creak SymLink)
   and run this command for using
   
   ```
   $ uvm list
   $ uvm 
   ```


<b> Optional </b>

1. Delete or Create JAVA_HOME 

- ENV : "JAVA_HOME"
 
   ```
   C:\Program Files\uvm_java\bin
   ```


## MacOS

## Linux
1. Extract Installer Zip to folder

  ```
  /usr/local/uvm/
  ```

2. Create uvm.sh in /etc/profile.d
  
- "/etc/profile.d/uvm.sh"
 
  ```
  export UVM_LINK=/usr/local/uvm_nodejs:/usr/local/uvm_flutter/bin:/usr/local/uvm_golang/bin:/usr/local/uvm_dart/bin:/usr/local/uvm_java/bin:/usr/local/uvm_python:/usr/local/uvm_ruby/bin
  export PATH=$PATH:$UVM_LINK:/usr/local/uvm
  ```

3. open shell like CMD or Powershell with Administrator Mode (Use for Creak SymLink)
   and run this command for using
   
   ```
   $ uvm list
   $ uvm 
   ```