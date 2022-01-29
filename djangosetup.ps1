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
