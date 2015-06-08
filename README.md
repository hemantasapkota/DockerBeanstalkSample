# DockerBeanstalkSample
Hello World in Golang served with Docker. Deployable to AWS Beanstalk.

### Dependencies ###

* boot2docker: ```brew install boot2docker```
* Docker: ```brew install docker```
* VirtualBox: Download 4.3.26 for OSX from [here](https://www.virtualbox.org/wiki/Download_Old_Builds_4_3)
* Golang: ```brew install go```. Ensure that **$GOPATH** has been set. Usually to something like: **~/go/**

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
