## How to Run on Docker Playground ? 

step 1 : Go to https://labs.play-with-docker.com/

step 2 : Add a new Instance

step 3 : Run Following Commands sequentially :
```
	1) docker pull ubuntu:latest

 	2) docker run -it ubuntu bash

	3) apt-get update

	4) apt install golang-go -y

	5) apt install git -y

	6) git clone https://github.com/Walchand-Linux-Users-Group/Quizzer.git

	7) cd Quizzer

	8) go run main.go dataPuller.go
```
