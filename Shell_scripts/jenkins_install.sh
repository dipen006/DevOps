#check if you have java

jenkinsinstall()
{
echo "Checking if Java is installed"

java --version

echo "If java is not installed please install java"

#installing wget

sudo apt -y install wget

#getting gpg keys with wget

sudo wget -O /usr/share/keyrings/jenkins-keyring.asc \
  https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key

# Adding Jenkins software repo to the source with auth keys


echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null


# updating the system

echo "Updating the system"

sudo apt-get update


# Installing jenkins

echo "Installing jenkins"

sudo apt-get install jenkins

}


jenkinsinstall


