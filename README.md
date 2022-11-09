# Overview

VBox VM setup script - Create VM and install OS with single command

## Quick help

VBox does not support silent install for every OS. If you receive some errors or the result will be kind of unexpected it would be good to check if for your ISO file is supported by VBox unattended install... It can be checked in VirtualBox app like blow:

![image](https://user-images.githubusercontent.com/62462151/200083625-b081d7e7-4ec6-4ab8-a11c-a013b07f5fd7.png)


# Build 
## Requirements

* GoLang installed in version 1.19 or higher
* Maven 3

### GoLang

* Follow installation instructions from: https://golang.org/doc/install
* Installation verification: executing `go` should bring GoLang

### Maven 3
* Follow instructions from https://maven.apache.org/install.html

## Building

Run command `mvn install` in root directory

### Output

Output binaries will be created inside `./bin` directory

# Run

To run script correctly user has to provide configuration file. By default `config.yml` from the script root directory will be used. It can be also provided via `-config-path <config_path>` parameter.
Use `-debug` flag if you want to see detailed info which commands are executed under the hood.

## Configuration

Configuration is split into two main categories `predefined` and `raw`. 

`predefined` is designed to create basic VM with hard drive and `nat` network settings including port forwarding. I designed it this way to have simple VM ready for further automation scripts.

Example - Simple LUbuntu 10GB VM with port 22 forwarded and openssh server installed:

```yaml
vbox-manage:
general:
  name: "TestVM"
  iso-file-path: <iso_root_dir>/lubuntu-18.04.5-desktop-amd64.iso
create:
  predefined:
    vms-home: <VMs_path>\VirtualMachines
    nat-rule: "guestssh,tcp,127.0.0.1,22,10.0.2.15,22"
    hdd-size: 10240
install:
  predefined:
    full-user-name: "test"
    credentials:
      user: "test"
      password: "test123"
    post-install-command: "apt-get install openssh-server"
```

`raw` section - this is for flexibility. Based on VirtualBox documentation: https://docs.oracle.com/en/virtualization/virtualbox/7.0/user/vboxmanage.html#vboxmanage-cmd-overview

User can provide his own commands to be applied **in addition** or **in replace of** `predefined` configuration. `raw` commands are simply executed after `predefined` configuration. Eventually, when `predefined` section is empty/not provided, it will not be executed and then, only `raw` commands will run.

Example - Raw commands:

```yaml
vbox-manage:
general:
  name: "TestVM"
  iso-file-path: <iso_root_dir>/lubuntu-18.04.5-desktop-amd64.iso
create:
  predefined:
    vms-home: <VMs_path>\VirtualMachines
    nat-rule: "guestssh,tcp,127.0.0.1,22,10.0.2.15,22"
    hdd-size: 10240
  raw:
    commands:
      name: modifyvm
      args: 
        - "TestVM"
        - "--memory=4096"
install:
  predefined:
    full-user-name: "test"
    credentials:
      user: "test"
      password: "test123"
    post-install-command: "apt-get install openssh-server"
```

## configuration - all available options
```yaml
# Path to vboxmanage binary, if not provided script will assume that it is available in PATH env variable
vbox-manage:
general:
  # VM name that will be visible on VirtualBox list
  name: 
  # Full path to the OS Iso file
  iso-file-path: 
create:
  predefined:
    # Path where to create root directory for your VM files
    vms-home:
    # OS Type (by VBox convention, See: vboxmanage list ostypes
    os-type:
    # Amount of CPUs - default 2
    cpus:
    # Memory size - default 2048 [MB]
    memory:
    # Virtual memory size - default 128 [MB]
    virtual-memory: 
    # Graphics controller name - default: vmsvga
    graphics-controller:
    # default: on
    rtc-use-utc:
    # clipboard mode - default: bidirectional
    clipboard-mode:
    # drag and drop mode - default: bidirectional
    drag-and-drop:
    # HDD size in Megabytes - default: 10240 [MB]
    hdd-size:
    # When provided it will configure port forwarding as specified - default: emty
    # example - Forward port 22 to Host: guestssh,tcp,127.0.0.1,22,10.0.2.15,22
    nat-rule:
    # VBox shared folder definition
    shared-folder:
      # VBox wide folder name
      name: 
      # Path on host file system
      host-path: 
      # Path on guest file system - it has to be absolute path
      auto-mount-point: 
  raw:
    # Commands to execute after predefined image is created
    commands:
      # command name, eg: modifyvm
      - name: 
        # Command arguments - syntax <parameter-name>=<parameter-value
        # example: --memory=4096
        args: 
          - 
install:
  predefined:
    # Full name of the user
    full-user-name: 
    credentials:
      # Username to be created in the installed system
      user: 
      # Password of newly created user
      password: 
    # Command to be executed after installation is done - it is OS dependent value
    post-install-command: "sudo apt-get update || sudo apt-get install openssh-server"
    raw:
    # Commands to execute after predefined OS is installed
    commands:
      # command name, eg: modifyvm
      - name:
        # Command arguments - syntax <parameter-name>=<parameter-value
        # example: --memory=4096
        args:
          - 
```

# Tips

## Unattended Templates
It turned out that VBoxManage is not the greatest tool... and it has some drawbacks

It's often better to modify "UnattendedTemplates" @see: `C:\Program Files\Oracle\VirtualBox\UnattendedTemplates`
instead of using `post-install-command`

eg. to install ssh server it may be useful to pust there something like 
```
log_command_in_target apt-get -y install openssh-server
```

## Sudo privileges - LUbuntu

to have sudo privileges on Ubuntu OS we can use `post-install-command` like below:
`chroot /target usermod -a -G sudo <user_name>`

## Automount folder permissions

VBox automount will create shared folder on your guest OS with permissions 770. If you would like to grant read permission for other user, you can do it eg. by manual mount:
`sudo mount -t vboxsf <shared_folder_name> /<guest_path>`
