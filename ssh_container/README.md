# SSH Container

## Prepare to use it
For using this container you need to generate ssh keys and keep it in a `sshContainer` folder in this project. It will be copied in the container while building. 
Use:
```
docker ps -a # find the container id
docker inspect | grep IPAddress # to find the ip address of the container
```
Connect to the container using: `ssh root@ip`

## Setup ssh authentication on a Linux Server
1. Generate your public-private keys: `ssh-keygen -t rsa`
2. Create `authorized_keys` file and copy a public key there
3. Go to the server and login as a user or root
4. If u are root:
	1. Go: `/root/.ssh/authorized_keys` and add a pubkey there
5. If u are a user:
	1. Create in `/home/{user}/` a folder: `mkdir .ssh && cd .ssh && touch authorized_keys`
	2. Copy your key there
6. Open config `vim /etc/ssh/sshd_config` and setup: `PasswordAuthentication no` 
7. Setup ssh on local (in a case if u have different keys): `ssh-agent /bin/bash && ssh-add .ssh/{your_private_key_file}`

