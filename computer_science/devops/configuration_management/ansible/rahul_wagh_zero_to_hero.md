Okay, here is the requested content formatted into a single markdown block, suitable for uploading to GitHub. It includes the notes you previously requested, integrated citations from the transcript excerpts, and the transcript excerpts themselves.

```markdown
# Ansible Full Course | Zero to Hero - Notes & Transcript

**Source:** YouTube channel "Rahul Wagh"

*(Note: The specific video URL is not provided in the transcript excerpts. Based on the title, a search suggests the video is likely at `https://www.youtube.com/watch?v=h9l7W9K-M4Y`. Please verify and update this link as needed as this information is outside the provided sources.)*

---

## Notes from Transcript

This course covers various topics related to **Ansible automation**, starting from basic concepts and moving towards practical applications and advanced features like variables, conditionals, roles, templates, error handling, and Vault. The course concludes with a project demonstrating how to use an Ansible Playbook to deploy a blog. If you are new to Ansible, it is recommended to go through all chapters. For those already familiar, timestamps are available to find specific chapters.

**What is Ansible used for?**

*   Automating the installation of software packages.
*   Automating the configuration of servers.
*   Deploying web applications or any kind of application on a server.
*   Automating tasks on remote servers.

**Ansible vs. Manual Tasks**

*   Without Ansible, installing software packages on a server involves manual steps, potentially using scripts. For example, running `apt-get update` and `apt-get install python` separately.
*   With Ansible, these tasks can be automated into a single Playbook.
*   Ansible Playbooks offer **self-explanatory task names** using the `name` attribute, unlike simple one-liner scripts. This makes them more intuitive for other developers.

**Key Concepts**

*   **Agentless**: Ansible is called agentless because you install Ansible on a developer machine (your local machine), and it connects to remote servers using **SSH** to execute commands, rather than requiring an agent installed on each server.
*   **Playbook**: An Ansible Playbook is where you automate tasks. It contains instructions for Ansible to follow. Playbooks are saved with the `.yml` extension.

**Ansible Playbook Structure and Syntax (YAML)**

*   Ansible Playbooks follow **YAML conventions**.
*   YAML requires careful attention to **indentation** to define the relationship between parent and child items. Different horizontal positions define the relationship.
*   A YAML file typically starts with **three dashes (`---`)**.
*   Tasks are defined with a leading dash (`-`) followed by attributes like `name` and the module to use.
*   The **`name`** attribute should provide a meaningful description of the task.
*   **Module names** (e.g., `file` for creating/deleting directories) are used to perform specific actions.
*   Module attributes (like `path` and `state` for the `file` module) are indented further than the module name to indicate they are child attributes. An indentation of one tab for the module name and two tabs for attributes like `path` and `state` is shown as an example.
*   A playbook can contain multiple tasks, following the same indentation rules.

**Ansible Installation**

*   Ansible is installed on the developer machine, not the servers it manages.
*   **Windows**: Use **WSL (Windows Subsystem for Linux)** to install a Linux distribution (like Ubuntu or CentOS). Documentation from Canonical (Ubuntu) is recommended for setting up WSL. Once WSL/Ubuntu is set up, you can follow the Linux instructions.
*   **Ubuntu (Debian-based)**: Commands include `sudo apt update` to update repositories and `sudo apt install ansible` to install. Additional commands like `sudo apt-get install software-properties-common` and adding the Ansible repository might be needed.
*   **CentOS**: Use the `yum` package manager. Commands include `sudo yum update`, `sudo yum install epel-release`, and `sudo yum install ansible`.
*   **macOS**: Installation is simple using the **`brew`** package manager with the command `brew install ansible`.
*   Verify installation using the `ansible --version` command.

**Ansible Project Structure**

*   A typical Ansible project structure includes directories for `inventory`, `roles`, and the main Playbook file(s).
*   **Inventory File (Hosts)**: Defines the remote machines (hosts) where Playbooks will be executed.
    *   It is usually located within the `inventory` directory.
    *   It can be named `hosts` or `hosts.ini`.
    *   It can be written with or without YAML syntax.
    *   Hosts can be defined by IP address or hostname.
    *   You can **group hosts** under meaningful names (e.g., `[gcp_host]`, `[blue]`, `[green]`).
    *   Multiple host files are possible.
    *   Groups can be used to categorize servers for deployment strategies, such as `[blue]` and `[green]`. Green servers might be actively serving production traffic, while blue servers could be a target for deployment.
    *   **Private SSH key paths** can be specified in the host file using the keyword `ansible_ssh_private_key_file`.
*   **Roles**: A way to organize tasks and other Ansible content. A role can contain multiple tasks. You can import tasks from other files within a role's `main.yml`. Roles are typically stored in a standard `roles` directory.
*   **Tasks**: Specific actions that Ansible performs on the remote host. Tasks are defined within roles or the main Playbook.

**Running a Playbook**

*   The basic command is `ansible-playbook` followed by the playbook filename.
*   You specify the inventory file using the `-i` or `--inventory` flag.
*   You can specify the remote user to connect as in the Playbook using `remote_user`.
*   Ansible uses **SSH** to connect to remote servers. Authentication can be done with username/password or SSH keys.
*   To use **SSH keys**, you need to generate a private and public key pair (`ssh-keygen`). The public key needs to be copied to the remote server. The private key path can be configured for Ansible, for example, in the host file using `ansible_ssh_private_key_file`.
*   You can limit playbook execution to a specific host or group of hosts using the **`--limit` flag**.
*   Playbook output indicates the status of each task (e.g., changed, skipped, failed). A **`changed`** status means the task modified something on the remote host. If a playbook fails, errors will be visible, often in red text. A successful run often shows a "SUCCEEDED" message with counts of changed tasks.

**Ansible Variables**

*   **Variables** allow you to store values and use them within Playbooks, making them reusable and flexible.
*   Variables can be defined using basic data types like **string**, **boolean**, **list**, and **dictionary**.
*   List elements can be accessed by their index (starting from 0).
*   Dictionary values are accessed by their key.
*   You can access elements in **nested dictionaries** using dot notation or square brackets.
*   The **`register` keyword** is used to capture the output of a task into a variable.
*   Variables can be printed for debugging using modules like `debug`.
*   Variables can be defined directly within a Playbook or in separate variable files (`vars` files).
*   Variables can be passed at **runtime**:
    *   Using the **`--extra-vars` flag** on the `ansible-playbook` command line, defining variables directly (e.g., `--extra-vars "version=1.0 who=user"`).
    *   Using the `--extra-vars` flag to specify a **YAML file** containing variables (`--extra-vars @vars_file.yml`). You can also pass JSON files this way.
*   **Variable Precedence**: Variables passed at runtime via `--extra-vars` take precedence and override variables defined elsewhere (like in Playbooks or vars files).
*   **Environment Variables**: Can be accessed within Playbooks. There's a distinction between environment variables defined at the Playbook level (accessible globally) and those defined only for a specific task (only accessible within that task's scope).

**Ansible Conditionals**

*   Ansible uses the **`when`** condition to execute a task only if a certain condition is met. It replaces traditional `if` statements found in other programming languages.
*   Conditions can be based on the value of a variable.
*   Multiple conditions can be combined using operators like `and`.
*   Conditions can evaluate the output stored in **registered variables**, for example, checking the `stdout` (standard output) or `rc` (result code) of a command. An `rc` of `0` typically indicates success.
*   Conditions can be used within loops to process only certain items.
*   Conditions can be based on **Ansible Facts**, which are automatically gathered information about the remote system (OS family, distribution version, kernel version, etc.). Ansible Facts hold metadata like IP address, operating system type (CentOS, Debian), version, CPU type, etc.. Ansible Facts are accessed using the `ansible_facts` keyword or directly (e.g., `ansible_distribution_file_variety`, `ansible_distribution`).

**Ginger2 Templates**

*   **Ginger2 templates** (`.j2` files) are used to generate dynamic configuration files on the remote server based on variables.
*   The **`template` module** is used in a task to process a Ginger2 template file.
*   You specify the source template file (`src`) and the destination path on the remote server (`dest`).
*   Variables are embedded in the template using **interpolation syntax** (`{{ variable_name }}`).
*   **Ginger2 filters** can be used within templates or with variables to manipulate data (e.g., `to_list`, `upper`). The pipe symbol (`|`) is used to apply a filter.
*   Template files are typically stored within a `templates` directory inside a role, and the `template` module looks for the source file (`src`) relative to this directory.

**Deployment Strategies**

*   Deployment strategies can be broadly categorized into **single server** and **multi-server** deployments.
*   You can manage deployments on **all servers** defined in the inventory by running the playbook without the `--limit` flag.
*   To manage deployments on a **subset or selected servers**, you use the `--limit` flag followed by the name of the host group.
*   The concept of **Blue/Green deployment** is mentioned, where servers are categorized into 'blue' and 'green' groups in the inventory, allowing deployments to be targeted to one group at a time.

**Error Handling**

*   Ansible provides features for **error handling** to manage unexpected issues during playbook execution.
*   The **`ignore_errors: true`** attribute on a task tells Ansible to continue executing the playbook even if that specific task fails. The output will show that the task was ignored.
*   You can handle errors based on the outcome of a registered variable, checking specific fields like `failed`, the error message content, or the **`rc` (result code)**. `rc` stands for result code.
*   You can use conditions (`when`) based on the content or status of a registered variable to trigger subsequent actions or fail the playbook.
*   The **`any_errors_fatal: true`** attribute can be used to make any error in a playbook immediately stop the execution and mark the entire playbook run as failed.

**Ansible Vault**

*   **Ansible Vault** is used to **encrypt sensitive data** like passwords or keys, preventing them from being stored in plain text in your Playbooks or variable files.
*   Basic commands for managing vaults include:
    *   `ansible-vault create <vault_file>`: Creates a new encrypted file and prompts you to set a password. Secrets are stored as standard YAML variables within the encrypted file.
    *   `ansible-vault edit <vault_file>`: Edits an existing encrypted file, requiring the password to decrypt and re-encrypt upon saving.
    *   `ansible-vault view <vault_file>`: Views the decrypted content of a vault file in read-only mode, requiring the password. (Mentioned conceptually, but a specific command demo wasn't in these excerpts).
    *   `ansible-vault rekey <vault_file>`: Changes the password of an existing vault file, requiring the old password to set a new one.
*   When running a Playbook that uses Vault-encrypted variables, Ansible needs the vault password. You can provide the password in several ways:
    *   Using the **`--ask-vault-pass` flag**, which prompts you to enter the password manually at runtime.
    *   Using the **`--vault-password-file <path>` flag**, which reads the password from a specified plain text file. This file can contain a randomly generated password.
    *   Exporting the password file path to the **`ANSIBLE_VAULT_PASSWORD_FILE` environment variable**. Ansible will then automatically use the file at that path. This is considered a neater approach.

**Project Example: Deploying a Blog**

*   The course demonstrates a real-world project: using Ansible to deploy a static blog hosted on SiteGround.
*   The blog uses **Hugo** static site generator, which processes markdown files (`.md`) to generate static HTML pages.
*   The deployment process involves running Hugo locally to generate HTML, compressing the HTML output (the `public` directory) into a **zip file**, copying the zip file to the hosting server using Ansible (via **SSH/SCP**), extracting the zip file on the server, and finally removing the zip file.
*   This automation saves manual effort when updating blog posts and associated files like images, which would otherwise require manual copying.

---

## Transcript Excerpts


hello there and welcome to this anible course this is Rahul and before we jump into our anible course let's take a pause over here and take a look onto the topics which we are going to cover into this particular Series so here I have listed all the topic and I will also leave the time stamp of each chapter into the description section so if you're just uh getting started with the anible then I would highly recommend to go through all this chapter and for those who are already familiar with these kind of a concept then just check the Tim stamp and look for the spefic


specific chapter uh which you want to practice or which you want to learn and the last chapter of this inable series I'm just going to cover one project where I have written my own anible Playbook to deploy my own blog so this Playbook will help you to understand how you can use your concept to create your own anible automation so don't miss that chapter if you want to really learn like how the anible work in real time so without further Ado let's start with our chapter one now back to the important questions


ation of a server the third thing and which is the most important thing is you can deploy your web application or any kind of application on a server and you can even automate that task using the anible All right so I don't want to stuff you with a lot of theory about anible and why don't we switch to our desktop and see a very small Playbook where we will try to see like how we can automate the installation of a software packages using the anible and also we will see if we don't have an anible then how we are


this is the way I'm just going to do if I don't have the anible the next thing which I'm just going to show a glimpse of a anible Playbook where we are just going to automate all this task into a single Playbook which is we call it as a anible Playbook all right moving to the anible part so this is the smallest sensible Playbook which I have for this uh particular very short demo so I'll start from the top over here so here you can see uh this is the first task which I have written over here for installing the python so here uh


one disclaimer I would like to give you here this is a introductory part here I'm just trying to explain the difference between if you don't use anible and when you use anible so don't try to uh like remember everything over here because we will take a deep dive look into the each file inventory task and the playbooks into the letter session and also we will take a look on how to install the anible also so we will start from that particular uh stage where we don't have the anible installed onto over


ation deck where I have shown you so here you will see the oneliner commands so these oneliner commands are not self-explanatory unless and until you know or you have used those command in the past so here you will see although it is really short and although you can uh combine all these commands into a single shell script file but it doesn't look so intutive if any other developer is working on to this bash script then he will find a little bit of difficulty understanding uh those command because those are


not so much they don't have the description behind them but if you take a look over here into the anible Playbook so here this is the first task here you have the name this name is really self-explanatory like what it is trying to do here is the second task where I'm also writing my own custom name like create directory I can explain further more like what kind of a directory I want to create and here is the third task where I'm just saying uh explicitly hey I'm trying to install the python so this is the


by one command onto my terminal so the first thing which I'm just going to do I'm just going to do the appt gate update which is uh for updating the package manager of my Linux machine since I'm using devian base UB 2 that's why I'm using the app uh package manager so I'm just going to paste the command over here and this is the way we are just going to update the package manager I'm just going to speed up some of the commands so First Command has been executed the next command which I'm just
is going to run is the install python so I'm just going to open the terminal once again I'm just going to paste it over here so here I just wanted to show you uh like if you don't use the anable then this is how you are just going to set up your server where you're just trying to install one package uh one after another so here you can see the python version is already by default installed so it is not going to install the python moving further so here I'm just going to create the directory so this is the command uh I'm just going to paste it over here


I'm just going to paste the command and hit enter and here you can see we have started our Apache 2 server the next thing which we're going to do we're just going to verify the status of a PES so I'm just going to run system CTL uh and here you can see uh the status of Apache to server service is green that means we have successfully installed the Apache to server on
to our Linux machine okay so now you have seen the manual way when we don't have the anible with me then this is the way I'm just going to install the software packages onto my remote server now I'm just going to switch over to the anible way where I'm just going to run only single Playbook and that will install all these packages using the anible All right so now we have seen the manual way of installing the software packages using the bash script but now let's uh take a look onto the Playbook once again so this is the Playbook which


have explained to you earlier and the name of the Playbook which I will be using for running this Playbook is VM setup playay book. yml this font might be small but uh when I'll switch over to terminal then I can show you the exact name okay so I'm just going to switch to my terminal and this is the command which I'm just going to run so that starts with the keyword anible Playbook after that you need to specify the remote host uh which we call it as a inventory in anible we will take a look uh uh in the letter session for the


and you just simply need to hit enter and if this is the first time you are connecting to your remote server then it is going to ask are you sure you you want to connect you simply need to type yes and hit enter and then it will take maybe couple of minute or 3 minute to complete the Playbook but uh meanwhile I can show you some of the output which it generate to tell you that what task it has finished so here you can see it is installing the python and here you can see uh
the status change that means it has installed the python the next task which it is doing it is creating a directory and the third task which it is doing is to install the Apache and as soon as all the three task has been finished and you will see a su a message over here that like three changed okay for and you will not see any error message that means your anible playbook has successfully finished and now we can SSH into server and verify that it has installed the required


version and here you can see python has been installed the third thing which I need to verify is the Apache 2 server so I'm just going to run the system CTL status and Apache 2 and here you can see it has uh installed the Apache 2 also so that means our anible Playbook has finished the task and it has successfully install all the three things which we have asked all right so now we have taken a look onto the very small demo of
anible and I really felt that demo is really essential for you to understand the anible better now in the next session we will take a look like why anible is called agentless and also we are going to take a look like how to install anible onto your local machine so that you can uh run the anible Playbook and you can just make a software installation change config and deploy your web application so those we are those are the topics which we are going to cover into the next session welcome to the second


and on the right hand side you will see the remote servers uh it can be on premise remote server it can be uh hosted on a Google Cloud platform or AWS or any other cloud provider but that REM uh server has to be remote okay so now uh there is anible and that anible we have installed onto the developer machines not onto the servers so pay attention carefully we have installed anible on a developer machine so if I'm working on my playbook then I have installed
nible on my local development machine so I will prepare my uh ncble Playbook I will write some instructions inside my ncble playbook and I I will run the ncble Playbook command and as soon as I run the ncble Playbook command it will connect to those servers using SSH it is same as when you try to log to your remote server using SSH so it will use the SSH it will connect to those server and it will execute all the required command which is necessary or which we have a speci


classified into our playbooks so this is why the anible is called agentless all right so now we know that anible is agentless now the next question comes like how can I write the anible Playbook and what's the extension uh in which I should save the anible Playbook so as you can see onto the diagram onto the left hand side so you need to save your anible playbook with the extension. yml so yml extension is used for anible Playbook and you have to follow the yml
conventions to write your anible instructions inside your anible playbook and here you can see I have taken a very basic uh code snippet screenshot where I have created uh like anible instruction for installing the Apache server so this file I'm just going to save with the yml extension and then I'm just going to execute the N Playbook command to uh run it on a different different server so this is the uh diagram which I just wanted to show you like how you're going to create and save your uh ncble Playbook with the


yml extension now the next thing which you need to do you need to install the anible based on the operating system which you are using so first I will start with the windows and as you know from the Windows 10 and windows 11 we have the capability to run the Linux uh operating system inside your Windows machine so for that we will be using WSL uh utility provided by our Windows to install the ubun 2 or a Santos based operating system inside your Windows machine because uh it is really easy to run


the anible in inside your Linux machine so for that purpose we are going to use the WSL utility for Windows here onto the screen you can see there is a really good documentation from a canonical entu community and this is the official page of the canonical entu where they have shown like how to set up your WSL on a Windows 10 as well as Windows 11 machines so here first of all you need to uh install the WSL utility so this particular section or these two links actually uh
this is for Windows 10 and this link is for Windows 10 as well as the windows 11 so I'll put both of the link into the description section of the video so if you are using the windows then I would highly recommend to follow these uh instruction which is mentioned on uban 2 canonical okay so let's go through it so this is the overview like uh what you need to do uh second if you click on the WSL then they have taken a really good screenshots like how to proceed forward so here uh you first go


installed the ubu then the First Command which you need to run is the Pudo appt update onto the terminal so that it can update its packages and at last you can install and use the GUI packages uh which is provided by a WSL for your ubu so that you can use the graphical interface uh for your UB 2 machine so just follow these instructions and uh uh set up your WSL uh based on the Windows 10 or Windows 11 operating system and once you have a WSL and ubun 2 setup on
your Windows machine then you can simply jump uh to the next instruction uh like uh the second and the third operating system like Ubuntu or Santo so here I'm assuming that on your Windows machine you are using Ubuntu so you can just follow the Pudo apt to get install anible command which I'll I'll show you onto my one to machine uh that how to install the anible so you can just simply follow those instructions to install uh anible because on Windows machine you are running to


then you just need to copy these instructions over there and you can uh just get yourself started with the so here is my terminal and this is my Ubuntu machine so first of all let's verify the version of my Ubuntu Machine by running this command so here you can see this is the Debian based uh operating system so deian based also like Ubuntu is a Debian based operating system so all the command which I'm running over here will be uh applicable for any Ubuntu machine so if you are using Ubuntu on your Windows machine then you can


simply run these commands to install the nle so the First Command which I'm just going to run I'll first of all clear the screen is Pudo app gate update so this is the first command I would recommend you to run so that it can update all the uh apt packages all right so that has been done successfully the next command which I'm just going to run is to in install the anible uh repository but before that I'm just going to run this uh install software properties common command type yes it
will install it okay that's been done I'll clear the screen now I'm just going to install the repository from where we are just going to fetch the an so this is the command for installing or installing the anible repository onto my Ubuntu machine so I'm just going to Simply hit enter and just going to press enter and here we have downloaded the repository for anible and the next command which I'm just going to run is the Pudo epate update once again because we have just recently added the uh anible repository so this is the


command to update the uh repository definitions over here and here we can see our anible repository has been updated okay I'll clear the screen and the final command which I'm just going to run is the Pudo apate update install anible and hit enter type yes all right so here you can see our anible installation has been finished so I'm just going to clear the screen and I'm just going to run the anible version command to verify
by the version uh I would say I'll C the screen once more I need to add one more Dash over here to get the correct version and here you can see this is the version uh Nal 2.10 which we have installed onto our2 machine I'll put all this command into the description section uh so that you can follow this command to install the uh anible onto your ub2 system the next uh we are going to do we are just going to install the anible onto the sent OS the Y


version so here I'm using the Santos release 8 version so that we have verified the next thing which we are going to do we are just going to run the Yum update command so that uh uh we update all the package manager so the command is sudo uh yum update okay so now we have finished the Yum update command so I'm just going to clear the screen the next command which I'm just going to run is the Pudo uh sorry yum install eepl release and I'm just going to add Pudo over here
so this will basically install the EPL release package and after that we will be able to install the anible onto our Santo machine type yes and here we have uh successfully installed the EPL release package and the next command which I'm just going to run is the I'm just going to clear the screen first of all sud sudo yum install anible all right so now the anible uh y install anible command has just finished I'm just going to clear the screen


and I'm just going to verify the anible installation by running the command anible version and here you can see uh we have installed the anible 2.3.5 onto our Santo system the next thing which we are going to try out is we are just going to install uh anible onto our Mac OS if you are using the Mac OS then installation of a anible on a Mac OS is pretty simple uh you just need to rely on the Brew package manager and the command for installing the anible is the Brew install
anel and that should take care of the anible installation onto your Mac OS so here you can see I have already installed anible that's why it is showing me that anel has already installed but in case if you don't have anel installed onto your Mac operating system then you just simply need to execute PR install anible and that will take care of the anible installation now after installing the anible the next lap session we are going to discuss more about the anible project setup so here I have taken a very brief screenshot


where I have shown a project structure of my anible project so where I have shown like how I have created the inventory file uh where is my host file what are the roles I need to create inside my anible playbook uh what does it mean by the task uh and how should I set up my anible uh Playbook yml so this is the project setup we will take a deep dive into for our next session welcome back and till now so far we have seen like what the
is and how to do the Nel installation but in this lab session we are going to take a look onto the project structure of our Nel so here we will talk about the inventory file we will talk about the task and our main Playbook and how to execute it so here you can see onto the screen this is the typical project structure of my anible so here we are talking about the inventory roles and my main Playbook file but don't worry we are going to take a look onto the each uh file which is present inside our project structure


moving to the next slide so here you will see a remote server with a dedicated IP so I'm just assuming that this is going to be my dedicated server and for This Server I'm just going to write my anible playbook so first of all in this slide you will see on the right hand side my dedicated server which is a remote server present somewhere and after that you will on the left hand side you will see this is the typical project structure of my anel so we'll start from the host file so if you'll see uh whenever you work with


mle entries inside this host file so this is how it works so here you will have a host file in your project inside the inventory directory and inside that inventory directory you can create a multiple host file also that's absolutely okay so I'm just taking one example uh for Server remote server so that's why I have only one host file and this host file contains the IP address and this IP address will point to our remote machine so that's the first uh basic file which we need
for creating our anible Playbook and that is our host file all right moving to the next file where we will be talking about the roles and the task associated with particular role so here onto the screen if you'll see this is the main. yml which is a task present inside a role and the role name is python so pay attention carefully over here for anible Playbook you can have a multiple role and each role can have a task so here I have created a role that is Python and inside that pyth


python we have a task and that task we have defined inside our main. yml and what's the purpose of that particular task so in this task we we are trying to execute uh the python installation and after that we are just trying to create a directory and then after that we just trying to install the Apache so this task basically contains the three uh instruction which we need to uh perform onto our remote machine using anel so once we execute this anible plate
book then this particular task that is python task will be executed and once it is executed then these three instructions will be pass to the server using our anel so that's the purpose of our roles and the task inside our anel all right so now the final piece inside our anible project is our anible Playbook so if you look onto the screen so this is the name of my playbook so the Playbook name is VM setup playay book. yml you can keep any name of your choice for your anible


Playbook that's absolutely okay now since we have talked about uh task as well as our host file so we need to include both the host as well as the task inside our anible Playbook so if you'll see over here so this is my uh VM setup Playbook yml that is a my playbook for anible and in this file you will see two thing so at the line number three you will find host that means what host it is going to connect so here I have specified all so that means
it is going to execute this Playbook on all the host which I have defined inside my host file so here in my host file I only defined one remote server so this Playbook will run on that particular remote machine so if you have more than one host inside my host file then this Playbook will run on all those host machine so that's why the host all we have defined over here the next thing we need to Define is the remote user so whenever you try to connect to any remote


here uh the the second uh picture if you'll see over here so where we have defined the main. yml and where we have defined the different different tasks so this is the role which we need to include and that is Python and if you that python role is available inside our project screenshot where you can see roles Python and task so here we have included the python role so this is how this complete anible works and how you put host as well as your roles and task as well as uh your
own user if you're are having your own user assigned to that remote machine so you need to put the usern name onto your anible playb all right so now you know like how the N project looks like and what is your host inventory roles and task now what's the command for executing this whole project so here onto the screen you will see the anible Playbook command which I have mentioned for this particular project and this anible uh Playbook command I have break I have broke into three parts so first part is the anible


Playbook that's the actual keyword which we need to use after installing the N second we need to specify like which host on which we want to run so here we have a specified inventory and then inside that inventory we need to specify the exact path of our host file so that host file is present under inventory VM setup Playbook and host that's the name of my host file so you need to carefully assign the host name or the host path over here the third thing is
is the actual Playbook name so that's my actual Playbook which I have created which consist of roles uh remote user and the host file on which host we want to execute actually so this is the anible Playbook command which you need to run to execute this whole project all right so now you have a little bit of fair understanding on the anible project structure now we will take a deep dive into the each individual file which I have just explained to you so here onto the screen the first file


which are which we are going to see over here is the VM setup Playbook yml so here first of all we need to specify the name of the Playbook uh that's something a description which you can put second here we need to specify the host so that I have specified as all so let's take a look onto the host file first of all so here onto this project you will find a host file which I have created and inside that host file I have created a suitable name like gcp host because that virtual machine is
on to Google Cloud but that's something I will explain it later like how to set up the Google cloud or AWS but that's something for later sessions but here uh I have specified the IP address and uh this is the uh private key which I will be using to connect to that particular machine but I have a next chapter for this SSH private key and public key setup with the anible but the key thing over here is this particular IP address so that's the remote machine IP address


which I will be needing or this anible Playbook will be needing so that's why uh this uh uh this Playbook contains the host all so that means all the host inside this particular uh this particular host file so here I have only one if in case I have more than uh one or two then I can execute the Playbook on all those host that's the purpose of that all inside that particular and civil Playbook okay moving back to the Playbook once again so now we have seen the name we have seen the host
now we need to look onto the remote user so as you remember whenever you connect to the remote machine then you need to have a remote user so here the remote user is rul p so that's the user which exist onto that remote machine I'll I'll show you all right so this is my terminal and I'm just going to SSH into this machine using the username Rahul V so I'm just going to paste the command over here so if you'll see this is the SSH and this is my host uh this is my priv


private SSH key and this is the username and the IP address of my machine and in case if you're are not using the private key then you can simply use the username and the password assigned to you so here I'm just using private uh and the public key that's why I don't need to assign my username and password but this just an example uh if you have a username and password then it's absolutely okay to SSH using the username and password and you can just simply enter type yes to accept it
and here you can see uh I just logged in into my machine and the username is Rahul V because that's the user I have created on that Linux machine so I just wanted to show you because uh if you go to this ncble Playbook once again so that's why I have used the remote user over here as Rahul V so you need to know like what kind of a user you want to become over there and if you want to execute with the root then you can mention the root also over here that's absolutely fine with the anible now the last


I'm just creating One Directory over here and then I'm finally I'm installing the apach so that's the task definition or the test task configuration of our anible looks like so these instructions going to execute onto our anible Playbook now the next thing which I'm just going to do I'm just to Simply run the Playbook and see if that works or not all right so now I have opened the terminal once again and I'm just to execute the anible Playbook command so this is my anible playbook command and I'm just going to Simply hit enter uh type yes


for accepting it and since I'm using the private key so I don't need to input my username and password and in case if you're not using the uh public and private key then you need to enter the username and password so that anible Playbook can run so that's the difference between this setup and the username password setup and this ncble Playbook might take a couple of minute to execute it so I'll just fast forward this uh session and I'll come back when this Playbook has executed successfully all right so here


here you can see our ncble Playbook had just finished and you can see the status over here three Chang that means it has installed python it has created a directory and it has also installed the aach to that's why the count as three and just in case if Playbook fails and if there is there are any error then those errors will be visible over here in the red color font but here our ncble Playbook has successfully finished so I'm just to clear the screen over here and we are just going to verify uh by logging into the server or just as


you and the public key will always go onto the remote server so first of all we are just going to create those keys and then we are going to use the private key for our host inside our host file and the public key uh which we are going to copy it to our remote server let's switch over to our desktop and see the demo so here onto the screen you can see uh on the left hand side this is my anible playbook where my private key is and then we are going to execute the anible Playbook and then it is going to communicate with our remote


on to the same location and the name of the key is anible uncore demo but you can specify your own path where you want to create those SSH keys I'm just creating it over here that that just for an example and simply hit enter I don't want to associate any pass phrase so I'm just going to keep it empty again empty and that's how we are just going to create the SSH keys so here you will see this is our first key uh that is ncore demo that is
is our private key and this is going to be the second key which is our public key uh that is also present under SSH and it is ending with the extension. pub that means that is our public key now let's take a look onto the previous slide which I have shown so here uh we have created the anible demo that is our private key as well as the public key the next thing which we need to do is we need to copy the public key to our remote server that means we need to copy _ demo. pup key to our remote server


remote server so once it is done then you just simply need to hit enter and it will ask for the password of our remote server so again I'm just to copy the password from my notepad and paste it over here and here you can see the one key has been added so that means we have successfully added our public key to the remote server now coming back to the presentation once again so our the part of copying the public key to the remote server is over now we need
to associate our private key with our anible Playbook okay so here's my anible playbook uh this is my anible playbook and where we need to associate our private key part so you need to go to the host file this is the IP address of my uh remote machine and this is the path actually this is the path of my private key so here this is the keyword anible SSH private key file that's the keyword from anible and here you need to specify the path of your


uh from my local machine to connect to our remote server so here you can see our uh command has successfully completed and this time I have reduced the Playbook so over here if you'll see so I'm just creating a very basic directory just to verify that I'm able to connect to remote server and I'm able to create a uh directory into my remote machine Let's SSH into the OR login into the server and verify that our directory has been created or not so I'm just going to use this command to SSH into my system


and here we are able to log in into our remote server and I'm just to run the ls command and here you can see our directory has been created with the name basic HTTP server all right so now you have seen like how to create the public key and the private key and how to use those keys to SSH into your remote server while you're running the anible Playbook but now into the next session we'll take a look on one of the most basic concept that is yml which is used for writing the anible Playbook so we'll take a


look onto the example yml which I have used in previous session so that is going to be one of the basic yml and we will focus on the concept like what are the indentation rule what are the syntax you need to follow while writing your own anible Playbook so that is going to be the next session for our anible series all right so let's talk about the yml and why should we use yml for writing your own anible Playbook obviously the first question comes in my mind that what is yml and what
it stands for so the full form of yml is yml and a markup language as the name suggests yml isn't a markup language so that means you don't need to use those traditional Syntax for writing your ymls so if there are no any markup syntax which you need to use for writing your own yml then what kind of a rules you should use for writing your yml so whenever you try to write a yml whether it's for your anible Playbook or if you're trying to Define any configuration so whenever


you write yml then you need to carefully look on the indentations now you might be wondering like what is indentation so you might be familiar with the concept of indentation in some other programming languages but when we talk about the ymls then you need to follow that indentation rule very carefully so here onto the screen you can see the example of indentation which I took so here you will see item one and item two that that are starting on a certain horizontal position but sub item one and sub item two is


is starting on a different horizontal position from item one and item two so these different horizontal position for each line item is called as indentation but you need to be really careful on defining the parent item and the child item so here the parent item I'm talking about is the item two is the parent and the sub item one and sub item two is the child so parent item will have a different horizontal position but sub item will have a different horizontal position so these horizontal
Position will Define like which is parent and which is child and using this concept we are going to write our anible Playbook where we will be writing the parent item with the task and the description and then also we are going to write the child item inside our Playbook let's take an example where I have written a very basic yml for my anible so whenever you write a yml then you should start your first line with the three dashes so here you can see onto the screen my first line start with the three dashes moving to the


next line and there where it comes the name so the name attribute defines like what your task does so here you can put a meaningful sentence like which describe what kind of a task you're trying to do so here I'm just trying to create a directory that's why I have put a description over here that is Meaningful that particular task but the second line always starts with a single Dash so it is always like single Dash and then name and then your description of that particular task
now moving to the third line and the third line is depend upon like what kind of an anible Playbook you are writing so here I'm just trying to create a directory so here I need to use the module name so the module which uh I need to use for this Playbook is the file so that's why my third line I'm using the keyword that is file and that file module is responsible for creating or deleting the directory whenever I'm working with the anle so that's why the file name comes


from and if you look carefully the line number third that starting at a certain horizontal position and that horizontal position is defined by a single tab all right so now we have defined the parent attribute as file the next to line where we are going to define the child attribute so whenever you're working with the file module of a anible then you need to define the two more attribute that is path and the state and if you look carefully over here the indentation of a path and the state
attribute is little bit more than the parent attribute of file so here you will see I have put two tabs for defining path as well as two tabs for defining the state and with these two additional tabs or the indentation of a horizontal space you tell yml that these are going to be my child attribute so here path and the state are going to be the child for a file module and what's the purpose of a path and a state attribute so whenever you work with the file module and if


can have multiple tasks so this is just a one task where I'm creating a directory but you can have a n number of a task inside your playbook but the rules are same you just need to repeat this process and follow the same indentation rules for writing your next task for your anible playbook I hope the yml part is clear to you the next file which I'm talking about is the host file so in the host file where you trying to define the host where you want to execute your anible play
the host file can exist with the name hosts or it can exist with the name hosts. inii and you can write the host file uh with yml and without yml also so the first example which you can see over here where I have written the host file without any yml rules so here I'm just trying to define the name of my host so here I have just taken an example gcp host and after that in the next line I have written the IP address of my remote machine so that's


the simplest way of writing the host file without using yml syntax but the next example which I'm showing to you is little bit more complex and here I have taken the same example which I have written without yml as well as with yml so in the host file you can follow both the uh rules like with the yml and without yml and you can construct the parent and the child relationship between the different hosts so this is the way you can create your own host file so it just depends upon
you like which kind of a snex you want to prefer but always remember try to put a meaningful name on the each host so that you can identify and you can pass the correct host while running your anible playbook so this is the host file which you can write either in vml or without using yml syntax I hope this session will help you to understand anible in a more better way and you can write your own anible playbooks with multiple task and the host file in the next session of this anible series we


updated content can reflect onto your website so how would you automate the whole process so the first task in this process is copy the updated HTML content onto your application server after that you once you have copied all the content you need to restart the application server so these two task you need to handle and to handle these kind of a task with anible you should use the anible handlers all right I will not make the things complicated so here onto the screen you can see the anible playbook for my anible


Handler so the first task over here is to set up the enginex application server this is just an example that's why I'm keeping the things simple over here so my first task is to install the engine server and after that I'm just to copy the content or the HTML file which I just want to update onto my application server so this is the second task which I will use to copy or to update the content on my website now you might be wondering like how should we incorporate the anible handlers over here so here I will


over here is the Handler and inside that Handler we are just to write the task to restart the engine server all right so now you have a little bit of understanding on what the anible handlers are let's switch over to the desktop and see the anible playbook in demo let's start with the demo and onto the screen you can see the ncble Playbook but let's first start with the project structure so on the left hand side you can see this is the ncble Handler project which I have created and inside that anible handlers project you will find inventory and the roles trory so let's take


next task is the roles so here in the roles I have created a role named Handler and whenever you create a role then you need to define the main. yml in any anible paybook so here I have created the main. yml so here I have written the task for installing my application server which is enginex which you can see over here and here I have just written the second task to copy the index.html the content of my website and the third task which I have written is to update the content of my website so
here I have here I'm just trying to copy the updated HTML to the destination directory so this is the uh simplest anible Handler Playbook which I have written it over here and in the last task I'm just trying to call the anible Handler so that it can restart the engine server so this is an anible Handler Playbook which I'll be using for this particular demo now one thing I forgot to show you is where I have written the Handler so here you just need to open the main Playbook that is anible Handler Playbook


engine service so the name uh the attribute which I'm just going to use is the engine service name over here and the state is restarted which means anel is going to instruct that please restart this particular service so this is how uh I'm just going to define or structure my anible handlers Playbook and also I'll put the GitHub repository link of this particular repo into the description section so feel free to copy that particular raer to use this anible Handler example now we have seen the anible handlers uh


updated the HTML content and that's for the reason we need a restart to our engine X server okay so this is my terminal and I'm just to run the ncble Playbook command so here this is my ncble playbook command so that's the keyword that is ncble Playbook after that you need to specify the name of the inventory uh that's the inventory flag which you first need to use and after that you need to specify the host file where your host file decides and if you go back over here so here is the inventory and this is the host file


type yes all right so as you can see my anible playbook had just finished so let's start each task one by one so here this is the first task update uh uh at cache and install engine so this is the first task which I have written so let me open the Playbook so yeah this is the first task which I just wanted to execute and that will install the enginex so if you go to the terminal so here you can see this first task has been executed second task is to copy the index.html and remember this is the first


time we are copying the index.html and we don't need to restart so that's why uh this task will not call the anible Handler so this is the second task we just wanted to execute so just check uh this is the index.html we just wanted to copy and this index.html is present into the same directory where I have defined the task or the main. yml all right let's go further after copying the index.html the next task is to copy the updated HTML to the
destination directory and just take a look so this is the updated HTML which is present into the same directory and I'm also I just wanted to copy into the same directory where my index.html is so it will just rename this direct uh particular file which is update. HTML to the index.html once it execute all right so as soon as it copies the update. HTML it should call the restart engine X Handler so let's take a look onto the terminal so here it is copying


sensible variable session so we will first start with the very basic data types that is string Boolean and then gradually we will take a look onto the list and the dictionaries also and furthermore into this session we are also going to take a look how to pass the variables at a runtime using the command line and as well as how to access the variables from your vars file and at last we are going to take a look onto the variables precedence so this is going to be a really comprehensive guide onto the anible variables all right now switching over to the desktop and onto this script


screen you can see this is the project which I have created where I have created a lots of variables and the sample files over here so I'm just to share this link of this GitHub repository into the description section so feel free to use this project and you can reference this example from this GitHub repository let's start with a very simple example so here onto the screen you can see this is the Playbook which I have written and the first variable which I'm just to create is the string variable and here onto the line number 11 you will see I have created a variable and the


to access them the next thing which I'm just to show you is the Ginger 2 filters so here uh this is the task which I have created for Ginger 2 filters and these Ginger to filters are really helpful when you when you want to do some kind of a string manipulation like converting lower case to upper case so these Ginger to filters are really handy but that's converting to lower case and uppercase is just one example or one use case but there are and number of application of Ginger 2 filter so the idea is just to learn the concept of a ginger


two filters over here so here what I'm going to do I'm just to create I'm just to access the keys of my fruit prices so if you go over here so this is our fruit prices dictionary and the keys are apple banana and orange so I'm just to access only the keys after that I'm just to convert that particular keys to the list so here if you see the pipe then this pipe is a Ginger 2 filter which is converting this particular keys to list after that I'm just to use a
again the same Ginger to filter and convert it into uppercase so this is how I'm just to do the uppercase conversion of my keys and finally I'm just to print it in I'm just to convert it into again a list and this particular task is going to print the value of that particular list now let's run this particular Playbook once again and I'm just to oh the screen is already clear so I'm just to run the Playbook command and here you can see using the ginger to filters on variables so this is the name of the keys like apple banana orange


and which got converted into the uppercase and these are Now list these are not dictionaries anymore so this is how you're just to use the ginger to filters inside your anible playbook moving further so till now you have seen like how to create the variables inside your ncble playbook and how to access those particular variable now if I talk about the scenario if you want to create a variable inside a different Wars file and you don't want to create the variable inside the same Playbook so then how you're going to do that


so if you take the next example which I'm just to show you over here so this is the sixth example which I'm just taking where I'm just to declare the variable or I'm just to create the variable in a different Wars file so the name of the file is my wars. yml and if I open that file so inside this anible variable project that there is a file which exists like my wars. yml and here I have created that particular variable so this is the name of my variable that is vs from my my Wars yml so this is just a very


and here I have used the variable that is vs from my Wars file so if you copy this and if you go over here then you should be able to find it so this is the variable which I'm trying to access it all right so now you have seen this particular example let's go to terminal and execute it so I'll clear it and here you can see this is the name of the variable and this is the value of that particular variable so idea was over here is like how to include or how to use the variable file
which is external not external but at least present into the different directory of that particular project and references into your anle playable so this is how your just going to do the inclusion of that particular variable file all right so now we have seen like how to create a variable inside the Playbook how to create a variable inside a variable file but next thing which I'm just to show you like how to pass the variable at run time so here I have just created a one task the task is pretty simple again so this is the name of


here but I'm just to pass that particular value of that variable at run time and how to do that so if you open this particular readme file so till now we were just simply executing this anible Playbook but here at the line number 10 this is the Playbook command which I'm just to use to pass the variable at run time and to do that you what you need to do you just need to add extra flags that is extra vars for extra variable and here you need to specify the name of the
variable that is version and its value and similarly you can pass one more variable other variable and the value that is who so here you can pass the runtime variable to your anible playbook and you cannot just only pass one variable but you can pass multiple variables let's copy this command from this particular R me and go to your terminal over here and clear the screen and paste the command and here you can see get the value from a runtime and if you go back over here


here so get the value from the runtime and the variable is version and you can see the value that is 1.0 and if I show you the again the readme file then you can see 1.0 so this is how you're just to pass the Val uh variable at a run time using your ncble playbook command and you can access the same variable inside your playbook all right so now we have seen like how to pass the variable at a runtime from the command line the next concept I'm talking about is can you pass the variable file


pass the yml file so instead of passing an runtime variable individually I'm just to pass the complete yml file name at runtime so this is the way you should do it and you can even pass the Json file also so what you need to do you just need to put at the red sign and after dat the name of that particular file but always remember try to put the correct path so if this file is present inside some of the directory then just use that particular path relative to that your anible label but anyway this


Playbook is available over here and uh this yml is also available onto the same level so I'm just to go to terminal I'm just to clear the screen and I'm just to run the command and here you can see the value has been printed the other variable F and the print the value of a variable when v file is passed at a run time so that's the name of that particular task so this is how you're just to uh create uh not create but uh you're just to pass the variable uh from the command line


so this is what I just wanted to show you uh this is one more example where you can use the runtime variables now the last thing which I just wanted to cover into this variable session is the variable precedence so how the variable precedence happen when you have multiple variables with the same name so let's take an example to understand this so in the my Wars files I have created a variable that is other variable with the value F but in the readme file at a runtime I'm just passing the same variable with the value f


variables but in the next session we're going to talk about the environment variables and how to use those environment variable inside your anible clickbook so that's going to be the next session and after that we are just to talk more about the if then else statement and how to use those if then else statement inside your anable play so stay tuned and see you into the next session of anible and if you're interested in any other devop topic like like terraform kubernetes then please go out and check my channel where I have uploaded those sessions so till then take care and bye-bye


I just wanted to demonstrate like I just wanted to access I just wanted to access the both variable that is example which is this as well as this uh like a task level environment variable which is this one so ideally uh this particular task we should only get the value of this variable but not this one because this my task level variable is a task level uh environment variable and I cannot access it because it is outside of this part uh outside of the scope of that particular task
so these are the three tasks which I have created uh let's switch over to terminal and run this playable all right so this is my terminal and this is the command which I'm just to run but in case if you are wondering like where to find this command then I have created this readme file over here inside this particular project so from here you can just copy this particular command and this whole repository I'll post it onto my description section so you can clone this particular repository okay so coming back to my terminal which is over here so I'm just to add one more flag for


verbos that is V so that we can see the output of our task I mean to say uh if I open this one so we can see this output like a built-in command where I'm doing the EO so that's why I have added a verbos parameter over there you just simply need to execute it okay so let's uh Analyze This output task level so first we will start from this particular task the name of the task is environment bar at Playbook level and here we are doing the Eco of example so if you go to terminal once again so
so yeah this is the task which we have executed and this is the output which we are getting that is fuar and if you check so yeah this is the fuar and we are able to access that environment variable inside that particular task and because this is a global or ncble level uh Playbook level variable all right now coming to the next task which is environment for only task level so here at the task level what we have done we have created this variable and the value is hello world and if you go back to terminal


statement so stay tuned for the next episode and in the upcoming weeks I'm going to upload more uh similar session onto the anible Playbook and anible Series so stay tuned for that and if you are new to this channel then you can just subscribe to this channel where you will find the content on terraform Docker kubernetes and sible and there are many sessions which I have uploaded for a Helm chart also so just go and check those lab session if you're new to the DeVos so stay tuned for that and see you into the next


next step or we will be executing the next sensible instructions here onto the screen you can see on the left hand side it's a traditional if condition with any other programming languages so where you specify if and inside that you specify a condition and after that you write a code which you want to execute based on that particular condition but in anible we actually don't have a if condition instead we have a when condition so whenever we need to put any condition then we are just to use the when conditional inside your anible playbook


in the beginning it might sound complicated but in actual it is really simple to implement a conditional using a when condition so here onto the screen if I scroll it down I have taken a very basic example to understand this van conditional so here this is my playbook where I'm just to install an Apache server onto my remote server but I want to install this Apache server based on one condition so here you can see this is the condition which I have specified using the V so here inside that condition you will see a


use the when conditional inside your loops and the last topic before the demo is anible facts so we can use the condition the van condition on anible facts also so anible facts are something which is like inbuilt so if you're trying to run an anible Playbook on a remote machine then anible facts will always hold the value of that particular remote machine and what does it mean by value I'm not referring to a single value but I'm referring here to the multiple values so it might contain what
the remote IP what's the operating system whether it's a sento whether it's a daan and what's the version it is running what's the CPU type what's the chip type so all these information which is like a metadata about that particular machine will be stored automatically inside the anible facts and using those in insensible facts which you can see I have just taken one example of this anible facts from the anible documentation so these are the facts you can capture for that particular remote machine and these uh particular variables like


that's the when condition I'm just to evaluate since I'm running the uban 2 and which is a devian base so I just I'm expecting the value is to be devian so these conditionals you can put onto the anible facts also so this is just a One More Concept which is left inside the conditional which is I just wanted to show you before the demo all right so moving to my intellig where I have written my anible playbook and on the left hand side you will see this is the project which I have created onto my GitHub repository I'll put the link of this particular


GitHub repository into the description section so feel free to clone that particular repository so here this is uh the project structure and here I have created a part nine and inside that you will find a task and variables file and also the main Playbook and after that you will find a read me file so you can just simply copy this anible Playbook command from here to run this particular anal Playbook but in the host file you need to uh update the remote host IP so here I'm just using my AWS ec2 instance


which is running onto the uh AWS Cloud so that's why I'm using it over here but instead you can use your own remote machine all right so that's all about the practicalities of this particular project but let's go back to the uh actual code where I have specified these conditionals so here this is the first block where I have specified install Apache web server based on the condition that is install Apache flag and that's a Boolean variable so let's check the value of that particular V
buan variable and that is true so that means if I'm going to execute this particular anible Playbook then it is going to install the Apache server onto my remote machine so that's the first use case what I I'm going to do I'm just to remove rest of the details so that we only run the required one I'm just to save this thing go to the readme file over here copy this anible Playbook command open the terminal and here into the terminal uh you can just simply paste that particular command for


a version which is greater than or equal to 20 but rest of the conditions are same so I just wanted to demonstrate you that's why I have put this particular task over here so onto the terminal if you go back I'll clear the screen once again and you can just simply run the N playable and here you can see the setting complex condition that was my task name and we can verify it over here so setting complex condition and here we are having two condition this is the first one and this is the second one and onto the terminal you will see the output the anible kernel version which


item so that's the simplest uh like a for Loop or The Loop which I have implemented over here so I'm just to head over to terminal I'm just to clear the screen and run the Playbook and here you can see the output onto the terminal but if you take a look carefully then you need to check which part has been skipped and which part has been changed so this skipped part which means we we are we are not printing 0 2 and four I mean these values are script because those were lesser than five


expecting value to be Debian if it's Debian then I'm just to print the distribution name of that particular anmol F so go to terminal clear the screen uh screen is already cleared so I'm just to execute the Playbook and here you can see uh anible distribution which is one two so that's how I'm just to evaluate uh my anible facts and after evaluating inside the van condition I'm just to print the web all right so that concludes our ninth lab session on anible and till now we have covered


fourth chapter which you can see over here from the content page and in the next chapter we'll be talking about the roles and the task associated with your Anvil Playbook and in that session we'll be talking about how to import the roles how to import the different task inside your anible playbook so stay tuned for that and if you interested into the similar content on devops where I talk about terraform kubernetes stalker containers and Helm chart then please consider subscribing to this channel because there are similar session I have planned in the future and upcoming months which you will see onto this


channel so till then take care and bye-bye hello there this is Rahul and welcome to the sensible series and in this 10th part we will be taking a look onto the task and the roles the idea of this session is to understand what are roles and how to create multiple task inside your anible playbook so in this session we'll first try to look and understand what is Task and the rules and after that we are just to perform a demo to understand both task and the roles all right so to
begin with it's really important for us to understand the structure of our roles and the task so here onto the screen if you'll see this is the screenshot which I have taken from my projects and into this screenshot you will see there is a roles so this is the roles which I have created this is the standard roles directory where we are going to define the roles and here I have created couple of roles so this is the first one and this is the second one so the name of the first role is custom role because I just wanted to create my own custom Ro but here in the second role if you'll see there is a unique name that is install Apache


still dependent on the subtasks so here there are a couple of tasks which I have defined that is messages. yml and that is packages. yml so these are the further subtask which I want to execute when I'm executing this particular anible Playbook so inside the main. yml you first of all you will see I have included or I have imported the package. yml that is package task and after that I have also imported the message. yml that is my message task and after that I'm just writing one my task which is to start the Eng genic Serv


so let's take a look into the packages. yml first of all so if you open this one then what it's going to do it is just to install some packages it is going to install enginex it is going to install G and it is going to install curve so that's the package. yml and it is a subtask present within the custom role and after that we need to include that into our main. yml all right so that's our packages. yml the second thing is the message. yml which I have already included or imported as a task so here what I'm doing simply I'm just copying uh the content


of our uh like a HTML file which is present for Engen server and after that I'm just including it over here into the message. yml and finally we are just in restarting or starting our engine service onto my UB Tu machine so that's how you're just to create a multiple subtask within that uh particular custom role and it can depend uh of course the name can be differ in your case but once you create those role then in that role you can create a subtask also and you can import those task within your main


yml all right let's run this Playbook to verify the task and the roles which we have created so what you need to do you just need to go to the readme file and here you will find the Run Playbook instruction so you simply need to copy this command so it's basically doing uh it's just running this Nel import Ro Playbook and from the host and the host I have already defined over here in the ncble import Ro and the host so that's where you will find the host definition okay so I'm just to open the terminal over here and I'm just to P the command and hit enter so it should execute and it should include


into the similar content on a devops then please considering subscribing to this channel so see you into the next session of anible and devops till then take care and bye-bye the objective of this session is to understand what is Ginger 2 template in anible and how to use the Ginger 2 template for your anible playbooks before we jump to the demo let's try to understand like what is Ginger 2 template so as you can see onto the screen I have listed two benefits of ginger 2 template the first one is the dynamic configuration and the second one is the


ible project that should end with the extension file name that is J2 that is Ginger 2 template and you can Define any name for your Ginger 2 template so here I have kept the name as a my config but in your case you can keep any name of your choice now once you have defined your Ginger 2 templates then I'm assuming your Ginger 2 templates is empty but here I have just taken a one screenshot and don't worry we are just to take a demo and where I'm going to show you the complete configuration also but here onto the


that's the directory and here I have created a roles directory and don't worry this whole project is available onto my GitHub repository so you can find the link of my GitHub repository into the description section and here into the roles directory I have created one rule that is install Light httpd that's the one more server it's simple HTTP server and here inside the task I have created a main.yml but along with that I have a my config do J2 that's my Ginger 2 template and
and remember I'm just to use this gingu template inside the main. yml that's why I have kept it over there but in case if you have kept that particular Ginger 2 template in some other directory then you need to put the correct path when using your Ginger 2 template all right that's that's something related to how to include your Ginger 2 template but if you take a look onto This Ginger 2 template then here at line number 14 this is the uh variable value which I'm passing dynamically so here onto the screen


you can see this is the interpolation syntax which is available here on both left hand side and right hand side and here is the value of that particular variable which we will be passing when we are running the anible label let's take a look onto the main. yml and how to include our gingu template so just double click onto this main. yml and here you can see at line number 11 that's the source directory and where I'm just to specify my ginger2 template file name so here I have a specified that is my config do J2 and remember


always keep this Source directory inside the template tag so that's the key thing when you when you are using the Ginger 2 template all right now the next thing which you need to do over here is to define the vars and in the vars you need to specify what variable you're passing inside your gingu template so here I'm specifying server undor port and the value is ET and which you can see over here into the gingu template which is server P so that's how I'm just to include my Ginger temp
template and I'm just to assign some value or passing some value to the server Cod variable all right so that's all about the ginger to template and how to include that Ginger to template inside your anible playbook I I would like to explain my anible playbook and what we are going to achieve with that anible Playbook so here I'm just to install Light httpd web server onto my Linux machine so that's the first task which I'm just to do secondly here in the uh second task I'm just creating a configuration


this is an httpd conf or like a server configuration which I'm planning to replace with my custom configuration and my custom configuration I would like to start my server on a port R so that's why I'm just to replace this particular file this light httpd con which is present on Etc and /l httpd directory so I would like to replace this particular file which is present over there and I would like to replace this file with my config dog
2 which is my ging 2 template and here I'm just to start finally my light httpd server let's run this Playbook and verify our Ginger 2 template whether it's working or not so what you need to do you just need to head over to readme file and here I have written the command to execute this particular anible Playbook so I'm just to copy this command go to terminal and paste this command so as you can see over here our task has just finished so here I have installed the light HTT


TP web server after that we have created our own custom configuration for that server and then finally we have started our server so what I'm going to do next I'm just to SSH into my server and I'm just to verify my uh server configuration whether we have properly replaced those server configuration with my Ginger 2 template or not so this is my terminal and I'm just to do I'm just to SSH into my server uh which is my ec2 instance running on AWS simply hit enter
and here I can uh see that I have logged in into my uh Linux server I'm just to clear the screen and I'm just to run the cat command to go to the ETC light HTTP and then I'm just to view the con file and if we scroll up then you can see this is the server port and the value which I have assigned over here is 80 so that's how we we are just able to replace our custom configuration for our server so that's a one place I have modified but


lebook more powerful all right so now you have seen the gingu template and how to use the gingu template to put your custom server configuration but I just want to take the same example and I just want to take it uh to the next level where I'm just to change the port from 80 to 9090 so why I'm doing this because you can use the same n Playbook and instead of running on a port 80 you can change the port to run it on a different port but that's not an ideal scenario but it can happen
happen in future that you want to change some of your configuration and instead of running a server on a particular Port you want to change to run that Port that server onto different port so that's a hypothetical scenario but that's not in actual real use cases but something similar can happen into your own configuration also so what I'll do over here so instead of running it on a 80 I'm just to change it to run it on a 9090 so with the help of this anible Playbook and Ginger 2 template you can simply update your anible play


and pass a new value to your server configuration so here I'm just passing the value 9090 I'll just save this changes and I'll head over to my terminal uh go to my ncble playbook I'm just to clear the screen and rerun the playable so this time what it is going to do it is just to update the server configuration and it is going to start my same server on a different port which is 9090 so here you can see uh it has changed the configuration
and you can see the color has changed over here which means our server configuration has been changed and what we can do we can again go back to the another uh terminal over here where we have already logged in I'm just to clear the screen over here and we are just to verify whether our changes has been reflected there or not so I'm just to rerun the cat command to view the changes and if we scroll up and here you can see so we are able to update our server configuration just by running the Playbook so that's the power


of your Ginger to template which you can extend and customize your anible playbook so that was all about the ginger to template in anible Playbook and so far we have covered sixth chapter and in the next chapter we will be talking more about the deployment strategy and where we will talk about like how you can use the anible Playbook on a single server and when you can use the same anible Playbook to run on a multiple server so that's going to be the next Topic in this series and if you new to this channel then please


please considering subscribing to this channel where you will see the similar content onto the devops especially like a terraform anible AWS gcp kubernetes and Docker so please considering subscribing to this Channel and if you have any question related to the today's session then please put down into the comment section and I'll try to get back to you so see you into the next session of the sensible series till then take care and bye-bye there are various deployment strategies when you are using anible to automate your deployment processes


which I have uh taken a screenshot over here and I have categorized my host into two category one is the blue and second one is the green so in the blue uh like a category I have two uh IP addresses so these are the two remote machines which I am using and here in the green also I have specified two remote IPS so here this is the third and this is the fourth server so here I'm managing the fourth server and I'm trying to deploy or automate this deployment process using the anible for our four
remote host machine now you might be wondering like what is this blue and green category so whenever you're having more than one server then it is always difficult to remember the IP addresses of each server so what we do in anible is like we create a category and here we have created two category one is the blue and second one is the green let's talk about the green category so in the green category in our example we are having two servers and those two servers are actively running in the production environment and actively serving the reest


so this is the way you can just manage the uh deployment uh using this limit flag and you can Target the servers which are present under the blue category let's start with the demo and here onto the screen you can see this is the part 12 and this is the project which I have created for this deployment strategy and I'll put the link of this particular project into the GitHub section so feel free to clone that particular project all right so talking about this project so first of all let's take a look onto the host file and first we will target onto the single server deployment


so this is the host file and inside the host file I have created a default and default category and inside the default category I have mentioned the IP address of my remote machine so I'm just to run this Playbook and this Playbook is going to run on this default machine this this default remote server so what I'm going to do I'm just to open the readme file over here I'm just to copy this command go to terminal I'm just to clear the screen I'll just go to the directory where my


project is clear the screen I'm just to shorten this path so that it is visible to you also okay I'm just to clear it and I'm just to paste the command over here and here you can see in this Playbook command I'm not specifying any limit so I just want this Playbook to run on all the host which are present inside my host file simply type enter type yes now my NC Playbook has just finished and it has inst
install python onto my remote server with this particular IP address all right so now we have seen the single server deployment strategy using the Nel let's update the host file and put the blue and green category along with the servers so I'm just to go to my project I'm just to open my host file and I'm just to erase everything from here and I'm just to add the blue and green category along with the remote server server IP address let's open the readme file over here once again and here I have mentioned the


anible Playbook command with the limit flag so what you need to do is you need to remove the line braks over here otherwise you might not be able to run the sensible Playbook okay that's been done so I'm just to copy this command from here I'm just to head over to terminal I'm just to clear the screen and I'm just to place the command over here so here you can see I'm just running the sensible Playbook with the limit blue so I'm just to update the blue category of servers so simply hit enter type yes


all right so now my ncble playbook has just finished and as you can see there are two IP address which has been updated so one of them is like this one 52 59 205 and 47 so this bye-by error handling is really a powerful feature which helps you to address the unexpected errors which may occur when you are running your anible playbook let's talk about the couple of scenarios which might help you to understand why you should need


see is the ignore errors flag so here let me explain this particular task and how we are going to simulate this particular error and then we are just to ignore that particular error so here at line number four if you pay attention then I'm just trying to copy index.html to my remote server at this particular path and the path is home u 2 so I'm just to first of all show you my remote server so this is my remote server and if I do the PWD so this is the path where I would like to copy that index. HTML


but to simulate this ignore error what I did is I just didn't provided the index.html so this index.html is completely missing so when this task run then it will not be able to find the index. HTML and if it is not able to find the index. HTML then it will result into error and at line number six I have put the flag that ignore errors that's true so which means if there is any error in involved in this particular task then this particular flag flag is going to
to ignore that particular error let's run this Playbook and if you are looking for this particular project then I'll provide the GitHub repository link of this particular project into the description section so feel free to use that particular Rao so here I have created a readme file uh which you can copy The anible Playbook command from here and also I would like to show you uh that this index.html so ideally I'm expecting this index.html to be present along with my main. yml but this particular file I have not provided it all


over here so that's just a a small thing I would like to mention before running the anible playable so I'm just to head over to my terminal I'm just to open another one I'm just to clear the screen and I'm just to paste the command over here and simply hit enter and here you can see uh index HTML copy with ignore error so that was our task name so you which you can see over here that was our task name and if you open the terminal then here you can see
could not find or access index.html so which tells you that this index.html is completely missing and anible Playbook couldn't find that index.html but since we have used that particular flag so it is ignoring this particular error which you can see over here and at the same time if you look at uh the final result of the sensible Playbook then you will see that it has completed but it has ignored that particular task so this is how you're just to use the ignore error flag inside your anible playbook


error message so if you take a look onto the previous task then here I'm just trying to capture the result of this particular operation which is a copy operation so here it is trying to perform a copy operation and whatever happens in that operation I'm just trying to register that output into this particular variable that is copy result so what I'm expecting is that in that copy result uh variable I'm expecting a message and that message should contain could not find or access
although that message is going to be a really long one but I just want to find this particular words inside that message so if you're trying to copy a particular file and if that file doesn't exist then it will result some error and that result error message will contain could not find or access index.html and I'm just trying to find that particular keywords inside that you can see what's happening inside or what are the values present inside this particular copy result variable so I have


just enable this particular debug so when I'm just to run the sensible Playbook so it is going to spit out uh what are the content inside this particular message and then we are just to see like what are the condition which it is evaluating inside this third task okay so I'm just to open the terminal over here I'm just to clear the screen and I'm just to run the ncble Playbook and here you can see uh we'll just start from the first so this is our first use case where we have ignored that particular error now second


and string variable from men Playbook so this is where I'm just trying to print the values or whatever message which is present inside that copy uncore result variable which I have registered so which you can see over here I'm just trying to print the values onto the console so that I can see what are the values present inside this particular register variable okay so the key thing which we need to look over here is we need to find this particular text inside that message and also the uh flag which is like copy result. should be true


so I'm just to open the terminal and uh first of all this is the variable that is copy underscore result that's okay and then I'm just to find the keyword so here you can see could not find or access so this particular keyword I'm just trying to find as an string inside that particular variable so yeah that is present and also I'm just to try to find this copy result and here is the file and that is true so these are the two condition I was force forcing uh in that particular task


to be evaluated and these both conditions are met and once these two conditions are met then I'm just making this whole anible Playbook fail with the any errors fatal fly so if you move down further then here you can see here I'm just trying to copy using failed when and any error fatal flag and here you can see this whole Playbook has just failed and here you can see the final status of our ncble Playbook that is failed equals 1 and ignore equals 1 so ignore one was our first use case and this is our second


change but it's actually a result code that's the actual full form of that particular RC having said that let's switch over to on sharing the similar topics on a weekly basis so see you into the next session of ncble or devops till then take care and bye-bye anible vault is a really important concept when you want to encrypt the data which you don't want to show in a plain text format while working with your anible playbook since we are talking about the anible VA then we should know the basic command on


your ncble vault which are instored inside your encrypted yml file the second way of using the ncble vault is to store the password inside a Vault password file so that you don't need to key in or you don't need to enter the password from the command line you can simply store your password for your Vault inside a plain text file and then you can supply that file uh while you're executing your anible playbook and the third way of using the ncble vault is to export the path of your ncble Vault password file


into the environment variable and the benefit you will get from that is you don't need to enter the password of your anible Vault and also you don't need to supply the path of your ncble Vault password file while you're executing the ncble Playbook you can just simply run your ncble playbook let's switch to my code editor and the first thing we'll start is by creating our anible vault so this is my code editor and this is the part 14 and this whole project is available onto my GitHub repository so I'll post the link of my GitHub


repository into the description section so feel free to clone this particular code repository for your own use let's start with the walk through of My ncble Vault project so here this is the part 14 anible VA which I have already created and inside that project you will find a readme file which I have already opened over here so here I have written down all the commands which I'm just to use inside this particular demo so the first thing which we are going to do is we are just to create an anible Vault and for that the command is anible VA create


and then after that you need to supply the path so here I will be creating My anible Vault yml inside the groups War directory and the name of my vault yml is my vault. yml I'll copy the command from here to create my anible Vault so I'll copy the command I'll go to terminal and I'll I from the keyword so that I can go into the insert mode and the first thing which I'm just to


create is like my uh Vault or maybe I can type in like my secret and then I can enter uh this is the first secret from anible Vault so this is the uh secret text which I'm keying in Inside My anible Vault and once you have entered your secret then what you can do you can press the escape from your keyboard and then save and quit which is w and Q and then hit enter
so by this way you are just to create your first anible Vault inside your anible project and if you go back to your project structure over here then in the group ports you will find a file that is myor wall. yml which has been created and if you double click on it then you can see this whole content inside that particular file is encrypted and you cannot view the content which we have written which we have written a simple message Inside My ncble Vault since we have created our first n


do is I'm just to use that anible Vault or the password which we have or the secret text which we have created or inserted inside my anible VA into my anible playbook so here if you take a look carefully so this is the uh variable uh which is storing my sensitive information so I'm just to use this variable to print the value or to debug that particular value inside my ncble pleb well this is just a demo where that's why I'm just I'm just using the debug but in case if you want
to use this for particular let's say for example uh creating a database connection and you want to print uh or you want to use the database password then you can use a similar approach and just Supply your variable where which is storing your database password so this is just in another example in the real time scenario how you're just to use this uh secret values but for the time being I'm just to use this particular uh variable from here I'm just to copy it and into my row


to there is a task and there is a main. yml so this is a a very debug message which I just wanted to print and here I have created a variable that is war and here I'm just to paste the name of my variable that is my secret now moving forward what I'll do I'll just run this anible Playbook using my anible P and for that you just needml using the


e keyword okay so this is the path where I'm storing my anible VA yml which you can find over here also so that's the important thing and the third parameter which you need to supply over here is the ask VA password so that you can key in or you can enter the password so that that you can open that particular anible world and you can use the values or the secret value which you have store so this is the ncble Playbook command which I'm just to execute from my my terminal and then after
that we're just to see the value which is printed onto the console all right so just copy this command from here go to terminal I'm just to clear the screen over here and paste the command and hit enter and this is the first way of running your anible playbook with your ncble Vault using the ask Vault parameter so this is the parameter which I have just applied with your with my ncble playbook command okay so I'm just to key in my password and hit enter and here


you can see this is the message uh this is the value which I have printed onto my console and it is able to fetch the value this is a first uh secret sorry there is a typo so this is the first secret message from my anible mod so this is the text I have stored Inside My anible Vault yml and now I'm able to fetch that particular information onto my console using the debug provided by n so this is how I'm just to use the secret value inside my anible playbook all right
so now you have seen the first demo where I have created an anible Vault and then I am able to run my anible plea book using the ask Vault pass parameter where I have key in and entered My anible Vault password okay so before moving forward let's take a few more example of our your existing password and you can set the new password over here
and you need to reconfirm the password and once you use this command successfully then you will be able to change the password of your anible VA yml so this is an useful command when you want to change the password uh from the security reasons all right so till now we have seen like how you can key in your ncble Vault password to execute your ncble playbook and to use your ncble vault in the second example what I'm just to do I'm just to create an password file and


ible VA okay so now we have generated our password file now what I need to do is I need to create a one more Vault using that particular password file so this is the command which I'm just to use and this command is again same anible Vault create and I'm just to create a new yml for My ncble Vault using the name My Vault with b64 pass yml and I'm just to supply the password file over here because I'm not keying in the password manually or by
from my keyboard I just want to use my existing uh password file which I have just generated so here I will just Supply the path of that particular password file uh to my this anm just to go press I from here and I'm just to type in my secret so let's say my uh secret uh uh from base 64 pass
file so this is the variable and I'm just to key in my text message uh text message uh from base 64 I'm just to type in base 64 VA and this is just a randomly text message I'm just to key in and after that save and quit so now what we have done we have created a password file and using that password file we have created an n


fault so next thing what I'm just to do I'm just to use that particular debug message inside my ncble fbook to print the value of that particular variable which I have just created using the anible Vault so here's a read me file and I'm just to run my ncble playbook but this time I'm just not to enter my password manually but I'm just to use my password file to run my anible playbook so this is the Playbook command which is anible Playbook uh keyword
after that you need to supply the flag inventory and Supply the host file where you are storing the host information the name of your playbook which is ncble Vault yml and after that you need to supply uh the Vault file path which is this one which is a new one which I have just created and after that you need to supply the password file over here and you should provide the correct path of your password file okay so I'm just to copy this command from here go to my terminal and


and I'm just to paste it over here and hit enter and here you can see this is the message which I printed onto my console using the debug and this is the text message which we have just keying in into our new andil VA and remember I have not entered My you need to


keep in and Export into the environment variable and once you export into the environment variable then you don't need to even Supply the path of your password file you just simply run your anible playbook and anible Playbook will resolve that environment variable which is this one and it will find the path of the password file and it will use that password file to supply that password to our anible Playbook so which is a more neater approach uh while handling your anible playbook and anible wallet


so in this readme file you will find a command to export your anible uh password file path and here remember this is the keyword which you need to keep uh keep the same name which is anible Vault password file because that is an environment variable and after that this is the particular path from here uh till your password file uh which you can see over here so this is the complete password file path uh so this path will change based on your local workspace setup so this is my current workspace and


this is the current path for my uh for my local uh development setup but just change this path according to your password file path so I'm just to copy this command from here and head over to my terminal I'm just to paste the command over here and if you take a look onto the command so this is the command uh here I have supplied the inventory after that this is


The anible Playbook command and after that the Vault file yml which I'm just to use it and after that I have not specified the ask VA password parameter for manually entering the password and neither I have specified my password file path so I'm just to Simply Supply My Vault so that it can understand which Vault I want to use and after that I'm just to hit enter and here you can see our n Playbook has started and here you can see this is the message it has fed from
from our ncble world so this is another way uh to run your ncble playbook and to export the path of your password file into the environment variable so this was a very small demonstration on how to use your anible Vault inside your anible playbook and you can use either of the three ways which I have mentioned into this lab session uh I'll put down the GitHub repository link into the description section along with the readme file and sometimes I also create a blog post on the similar lab session so


if you don't find these read me file inside which is J hook.com here you can see this is the blog which I have been hosting for a quite some time and it is registered on a domain J hook.com so this is the blog which I have been writing on a various devops topic but this particular blog is uh using the service provided by a site ground hosting so this is the hosting provider where I'm hosting my blog so I have
already purchased this domain and then I have registered this domain with this site ground hosting provider and here you can see this is the dashboard of my uh particular blog where I manage the deployment of my whole blog before I walk you through the various options available onto this s ground hosting and how I am hosting this blog and how I am uploading the files to update my blog post it is really important for you to understand that how this blog works so first of all this blog is not a WordPress blog it


is a static blog what does it mean by a static blog so this blog is completely written in markdown language as well as the pages which you see are just the simple HTML Pages there is no database running in the behind the scene or in the background so here onto the screen you can see this is the GitHub repository where my whole blog is hosted not hosted but I' your local system and you can also start uh building your own blogs so this is the framework which I'm using uh for my own blog development and writing the various blog post okay so first thing first like how does it work so if you go back to the intellig so this is the GitHub repository where all of my code recites for my blog so what I do generally is I just start my Hugo server so here if I go to the terminal I'll just open in a window


mode uh let me view mode change it to the Window mode okay so I'll just go back to the directory and yeah this is the correct directory where I'm in right now so the command which I use it uh remember you need to install the Hugo onto your local system to make it work okay so I'm assuming that you have already installed the Hugo and after that I just simply run the command Hugo ser and here you can see it is running on a local host 13313 port and if I go back to my browser and here if I
I type Local Host 1313 and here you can see uh this is the blog uh which is running onto my local system so here I make the changes I verify the changes and once I a little brief about like how I uh structured my uh various blog post topic so here you will find a Content directory and in the content directory you will find a post directory and inside the post I have created a various directory where I keep on writing the various blog posts


so for example anible AWS devop stalker GitHub so these are the various categorization which I have done so let's take a look onto the nple because the latest post which I see is this one which I have written on September 19th last year I have not been active on this blog but let's take a look uh so here this is the uh markdown file which I generally create for my writing my blog so uh this is a markdown uh structure which you need to follow for writing this blog post I
I know it it might look complicated for you but if once you get familiar with using the markdown then it's really easy to manage your blog post and you can write it from anywhere and you don't need any Hy fancy tool to manage your blog you just need a editor I'm using intellig you can use Sublime notepad++ and you can manage your blog anyway so what I'm just to do over here is I'm just to change a few things over here so let
post and here you can see the last modified date has been updated to 6th of March okay so this is the modification we have done and we are able to verify this change okay now the next thing is I'll show you how I generate the HTML file and here you can see this whole uh block post is still into markdown file which you can see it is ending with the file. MD which is a markdown extension so using this uh markdown file we are just to generate an HTML


page and that HTML page we are just to copy it to the site ground server now how to generate the HTML page using the markdown which I have just updated so for that we are going to use the Hugo command utility because I have already shown you like I have already installed the Hugo onto my local system so here onto the terminal what I'll do I'll just go I'll clear the screen over here and first thing the uh whenever I generate the HTML file then it will generate a public folder over here so


here you can see there is no public folder which you can see and in that public folder it is going to generate all the HTML file after passing the markdown file okay so let's go back over here uh onto my this public folder you will file all find all the HTML for all the blog post so here you can see this is a really long list because I have posted quite a lot of post onto my


my blog so all the pages are all available over here and if you open any of the page let's say I open this one then you will see a index.html page and if you'll open this one then you will find all the blog post in the form of HTML so this is how I write my markdown uh blog post and then I generate the HTML page so this is the basic setup for my uh blog and this is how I write and generate the HTML in the next uh step
we are going to take a look like how I have written my anible playbook to copy this HTML this public folder to my site ground hosting server okay so let's recap so here is my laptop here I have writing the markdown file and which markdown file I'm paring using the Hugo and Hugo is generating me the HTML pages so that's a very basic setup which I have onto my local laptop now the next thing is how to copy that particular file to uh my remote server which


key with key with me that is only available for me so here in the file manager you will find uh the public HTML directory which you can see and here in the public. HTML you will find all the HTML pages so if I open this one and here you can see the index.html page so these are the files which I generally copy from my local


machine to my site ground hosting server okay but how do I do that so what I do is like this is a flow which I have implemented using the N so I hope now the SSH key part is clear to you I'm copying HTML file from my local laptop to my site ground hosting server okay so this is the HTML files out of it and this is how I do the deployment of my website onto the site ground using the N Okay so


that part is clear to you now coming back to the point I need SSH keys I need a private key and I need a public key so that our whole uh uh communication is secured between the side ground on my and also my laptop okay so I'm just to open the terminal over here so let's open the terminal over here and in the terminal you can just type the command SSH key gen and it will generate the public key and the private key for you so here as soon as you enter the SS
keys in onto your terminal if you're using the Windows system then I would highly recommend to use git bash because that's where you can run these command or you can use the wsl2 which is like a uh sub Linux subsystem which you can install onto to your laptop uh if you're using the windows so I would highly recommend to use those so that you can generate the SSH keyen I'm using the uh Mac so here I have the terminal so here I can run these kind of a command so it will generate


the key onto this particular directory so keep in mind uh just always I'm just to copy this path clear the path and I'm just to run the command LT and I'm just to paste the path so here you can see this is our private key and this is Pub is going to be our public key so this is how you're just to


generate the keys for you and these Keys we are going to use to upload our HTML Z file okay I just wanted to show you the key because the keys I have already generated and I have already set up for my site ground hosting so I'm not going to change those thing but all the steps which I your whatever hosting you're using so this is that simple okay so now uh this is our SSH key manager I'm not going to use this one and here we already have those keys with us okay now I have explained you the whole setup and how we are going to implement these things and what are the prerequisites which is needed to achieve this thing so now I will explain to you like how I have written the anible Playbook to implement the whole scenario okay so in the same uh reposit


itory where I'm having my blog uh I have created anible Playbook directory and here in the anible Playbookbook with this particular remote user and this is the task which I have created so let's take a look onto the task so this is the J hook uh go to inventory not inventory go to roles go to Jook go to task and there is a main. yml okay so this is the U main Playbook which is responsible for copying the content all right so first thing uh
this is the simple task which is just outputting the path which is PWD the current path second it is running the Hugo command it is just a simple Hugo command to generate the HTML content so if I uh if you remember then I have shown you like this is the command which I have run previously so this is the Hugo command which I have run so which has generated the basic HTML Pages for us in the public directory which is uh


this one where it is this one okay so I'm just simply running the Hugo command so that I can generate the HTML Pages for my block posts after that uh this is simple debug I have written just for debugging whether it is on a correct path or not second here I'm running the command to compress the directory so I just wanted to compress the public directory which is this directory so here this is the path of that uh sorry
this is all the content and in the destination it is again using the same project path and here it is creating a zip file name and what is the zip file name variable it is jhop do zip okay so here you can see uh we have the public directory but I have not run the Playbook so there is no zip uh J hook. up which has been created I'll show you once I run this Playbook okay so here uh
here till this part we are just creating the Z file so if you take a look onto our presentation once again so here uh HTML has been generated and now we have created a zip using the zip uh anible task okay picore HTML and if I show you over here then if I go
uh to file manager and if I go here you can see this is the path actually I'm inside if you go here then here you can see this is the path which is uh where I need to upload it job.com public HTML and which is this one exactly okay and let's go back to the task uh task yml once again and after that I'm just after copying that particular file


uh if I take a look over here so here I just copied it the zip file over there after that I have written extract task so this private key so in the same anible Playbook directory in the inventory you will find I have used this private key which is J hook this is our private key which is I'm only


keeping it over here although I have copied this public key also but I'm not I'm not using this public key anywhere I just copied it so that if I uh if I change my laptop if I format my laptop then I have a copy of that particular key I don't need to regenerate that key so that's the purpose I have kept it over there and in the host file if you take a look uh if I go and show you once again so this is the uh sorry it should be it should be J hook
uh I have just changed the host name so here is the name of my Host this is the uh path of my private file and this is the port where it will be communicating okay so this is the defination or this is the configuration for my host file for running this anible okay


now we have seen everything now we need to run this Playbook okay and if I go back to my website uh refresh this page over here and here you can see this uh uh question mark is still there and if generally run so this I'll just copy this command I'm just to explain to you also


that one and I'm just to paste it so this is the simple command which is n Playbook then there is a inventory and inventory I'm specifying the host file the host location of the file and the name of the Playbook which I'm just running which I have just explain to you okay it is that simple and after that you just simply hit enter and here you can see it started to uh create uh the public HTML that it is going to now it is compressed
in the file now it is copying uh automated my whole deployment because uh what happens is like I generally write quite a lot of blog post and I modify quite a lot of blog post so it's a quite
manual thing if I write a four or five blog post and those blog post has certain images so a single blos block post update can contains four or five images four or five files so I just need to manually copy those files if I don't use anible so instead of uh doing that task manually I just return this anible Playbook which just generate everything zip it everything and upload it to the server and that's how I just automated this task

[16 the next session till then take care and bye-bye
```
