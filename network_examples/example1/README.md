# Set up Gateway for Docker Containers in Different Networks
This guide explains how to set up a gateway using Docker containers, allowing two containers in different networks to communicate via a third container acting as a router.

## Prerequisites
- Three Docker containers:
  - `ubuntu1` in Network A (`172.18.0.0/16`)
  - `ubuntu2` in Network B (`172.19.0.0/16`)
  - `ubuntu_out` connected to both Network A and Network B, acting as the gateway

## Step 1: Verify Container IP Addresses
First, check the IP addresses of the `ubuntu2` and `ubuntu_out` containers to ensure they are in the correct networks.
```sh
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ubuntu2
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ubuntu_out
```

## Step 2: Test Container Isolation 
After connecting by ssh to a container, let's ping `ubuntu1` from `ubuntu2` (address could be changed):
```sh
ping 172.18.0.2
```
Since they are in different networks, the ping should fail.

## Step 3: Add a Route via the Gateway

To enable communication, add a route to ubuntu2 so that it can reach ubuntu1 via the ubuntu_out gateway.
```sh
ip route add 172.18.0.0/16 via 172.19.0.2  # Replace with the IP address of `ubuntu_out` on Network B
```

## Step 4: Test Connectivity
After that we can ping our `ubuntu1` container:
```sh
root@d57b5d76f891:~# ping 172.18.0.3
PING 172.18.0.3 (172.18.0.3) 56(84) bytes of data.
64 bytes from 172.18.0.3: icmp_seq=1 ttl=64 time=0.056 ms
64 bytes from 172.18.0.3: icmp_seq=2 ttl=64 time=0.050 ms
^C
--- 172.18.0.3 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1026ms
rtt min/avg/max/mdev = 0.050/0.053/0.056/0.003 ms
```
If the route is set up correctly, you should see successful ping results.

## DONE

