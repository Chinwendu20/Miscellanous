# Scripts
 This is a repository of scripts that I have created to help make my life easier.

 ## Description

 - **djangosetup.ps1**

 This is a powershell script that initializes the django environment for development.

### Functionalities

- Makes new project directory
- Creates python virtual environment
- Activates the virtual environment
- Installs django and rest framework
- Creates django project
- Creates django application
- Carries out initial migration for django application

### Pending updates

- Adds django newly created app to installed app list in settings.py file

### Requirement

I assume that if you have the intention of working on a django project, you have python already installed on your computer. If not, here is a  [link](https://www.python.org/downloads/) to guide in installing python.



### Installation guide

This script is published on powershell gallery. Install by running the command below:

```powershell

Install-Script -Name djangosetup

```
