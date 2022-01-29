
<#PSScriptInfo

.VERSION 1.0

.GUID 627c132b-52e6-4d53-a75a-cc9ecc43d091

.AUTHOR
Ononiwu Maureen Chiamaka 

.COPYRIGHT
Check me out on linkedin and tell me, you are using it: https://www.linkedin.com/in/maureen-ononiwu-49b3b212a/ 

.PROJECTURI 
https://github.com/Chinwendu20/Scripts.git


.RELEASENOTES
Recieve with love



#>

<# 

.DESCRIPTION 
 This is a script created to set up the django environment for application development 

#> 
Param()


$Folder = Read-Host -Prompt 'Input folder name:'
$Project = Read-Host -Prompt 'Input django project name:'
$App = Read-Host -Prompt 'Input django app name:'
echo "Making new directory"
mkdir $Folder
echo "Changing directory to newly created folder"
cd .\$Folder
echo "Creating virtual environment"
python -m  venv venv
echo "Activating virtual environment"
.\venv\Scripts\activate
echo "Installing django and rest framework"
pip install django
pip install djangorestframework
pip install markdown
pip install django-filter
echo "Creating django project"
django-admin startproject $Project
cd $Project
echo "Creating django application"
python manage.py startapp $App
python manage.py migrate
