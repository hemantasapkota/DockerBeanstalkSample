# DockerBeanstalkSample
Hello World in Golang served with Docker. Deployable to AWS Beanstalk.

### Dependencies ###

* boot2docker: ```brew install boot2docker```
* Docker: ```brew install docker```
* VirtualBox: Download 4.3.26 for OSX from [here](https://www.virtualbox.org/wiki/Download_Old_Builds_4_3)
* Golang: ```brew install go```. Ensure that **$GOPATH** has been set. Usually to something like: **~/go/**
* AWS Elastic Beanstalk CLI: ```pip install awsebcli```

### Setting up Boot2Docker ###

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
  * Get boot2docker's ip: ```boot2docker ip```. Example output: ```192.168.59.103```
  * open ```http://192.168.59.103:8080```. You should have a running app now.

### Deploy to AWS Beanstalk ###

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
