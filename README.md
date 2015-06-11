# DockerBeanstalkSample
Hello World in Golang served with Docker. Deployable to AWS Beanstalk. For OSX.

### Dependencies ###

* Docker Machine: Download latest release from https://github.com/docker/machine/releases
* boot2docker: ```brew install boot2docker``` ( Only required if docker machine does not work )
* Docker: ```brew install docker```
* VirtualBox: Download 4.3.26 for OSX from [here](https://www.virtualbox.org/wiki/Download_Old_Builds_4_3)
* Golang: ```brew install go```. Ensure that **$GOPATH** has been set. Usually to something like: **~/go/**
* AWS Elastic Beanstalk CLI: ```pip install awsebcli```

### Setting up Docker Machine ###

* Download the latest release from https://github.com/docker/machine/releases
* Rename executable to **docker-machine**, copy it to ```/usr/local/bin/docker-machine``` and ```chmod +x /usr/local/bin/docker-machine```
* Create a VM using: ```docker-machine create -d virtualbox dev```
        
        Creating VirtualBox VM...
        Creating SSH key...
        Starting VirtualBox VM...
        Starting VM...
        To see how to connect Docker to this machine, run: docker-machine env dev

* Execute: ```docker-machine env dev```

        export DOCKER_TLS_VERIFY="1"
        export DOCKER_HOST="tcp://192.168.99.100:2376"
        export DOCKER_CERT_PATH="/Users/hemantasapkota/.docker/machine/machines/dev"
        export DOCKER_MACHINE_NAME="dev"
        # Run this command to configure your shell:
        # eval "$(docker-machine env dev)"
        # Save the exports to ~/.bashrc or ~/.zshrc file
        
* Check the status of the VM: ```docker-machine ls```

        NAME   ACTIVE   DRIVER       STATE     URL                         SWARM
        dev    *        virtualbox   Running   tcp://192.168.99.100:2376
        
        # Take a note of the IP address of the machine

### Setting up Boot2Docker ( Only set this up if, docker-machine fails for some reason ) ###

* ```boot2docker init```
```
Latest release for github.com/boot2docker/boot2docker is v1.6.2
Downloading boot2docker ISO image...
Success: downloaded https://github.com/boot2docker/boot2docker/releases/download/v1.6.2/boot2docker.iso
        to /Users/hemantasapkota/.boot2docker/boot2docker.iso
```

* ```boot2docker up```

```
Waiting for VM and Docker daemon to start...
....................ooo
Started.
Writing /Users/hemantasapkota/.boot2docker/certs/boot2docker-vm/ca.pem
Writing /Users/hemantasapkota/.boot2docker/certs/boot2docker-vm/cert.pem
Writing /Users/hemantasapkota/.boot2docker/certs/boot2docker-vm/key.pem

# Run the following exports. Note the IP address for DOCKER_HOST in your machine may be different.
export DOCKER_HOST=tcp://192.168.59.103:2376▫
export DOCKER_CERT_PATH=/Users/hemantasapkota/.boot2docker/certs/boot2docker-vm▫
export DOCKER_TLS_VERIFY=1
```

Add those variables to your **~/.bashrc** or **~/.zshrc** ( if you're using Z Shell or oh-my-zsh )

### Docker ###

* Clone this repo: ```git clone https://github.com/OrganicCoffeeNepal/DockerBeanstalkSample```

* Build docker image: ```docker build -t docker-beanstalk-sample-app . ``` 
    * You should get a **Successfuly built 9ac180ca5317** message in th end.

* Run the app: ```docker run -it --rm -p 8080:8080 docker-beanstalk-sample-app```
    * Log: ```[negroni] listening on :8080``` 

* Test the app
  * List the IP using ```docker-machine ls```
  * If you're using boot2docker, get boot2docker's ip: ```boot2docker ip```. Example output: ```192.168.59.100```
  * open ```http://192.168.59.100:8080```. You should have a running app now.

### Deploy to AWS Beanstalk ###

You'll need Amazon AWS account. If you don't have one, you can register for Amazon's free tier service.

#### Elastic Beanstalk Init - eb init ####
```
cd DockerBeanstalkSample

eb init -i

#Select your region

# AWS Credentials - http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSGettingStartedGuide/AWSCredentials.html
# Provide your aws-access-id
# Provide your aws-secret-key
# Note if you're creating new user through AWS IAM, you'll need to attach AmazonEC2FullAccess and AWSBeanstalkFullAccess policies to the user.

# Application Name: DockerBeanstalkSample

# It appears you are using Docker. Is this correct ?
# Answer: Y

# Do you want to set up SSH for your instances ?
# Answer: Y

# Select a keypair:
# Answer: 2

# Type a keypair name:
# (Default is aws-eb): aws-eb
# Generating public/private rsa key pair.
# Enter passphrase ( emptry for no passphrase): sample
```

#### Elastic Beanstalk Create - eb create ####

```
eb create

# Enter environment name (default is DockerBeanstalkSample-0): 
# Enter DNS CNAME prefix (default is DockerBeanstalkSample-0):

```

About 5 - 10 minutes later, you should get:
```INFO: Successfully launched environment: DockerBeanstalkSample-0```

http://dockerbeanstalksample-0.elasticbeanstalk.com/

### Elastic Beanstalk Update - eb deploy ###

Just execute ```eb deploy``` to push out new versions.

### Author ###
* [Hemanta Sapkota](https://twitter.com/ozhemanta) / [Blog](http://hemantasapkota.github.io/) / [Linked In](https://au.linkedin.com/in/hemantasapkota)

### LICENSE ###

The MIT License (MIT)

Copyright (c) 2015 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
