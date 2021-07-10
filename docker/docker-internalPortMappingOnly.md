# Prevent Docker from punching holes in a server's firewall

Docker will directly modify iptables depending on how you map ports to containers.

For example, say you run `docker run -p 7000:80 website` - this will map port `0.0.0.0:7000` on your host machine to port `80` inside the container, and in doing so will modify iptables to make `:7000` available to the outside world.

If this is not desirable, and you'd like your service to only be available locally inside your host machine, you should run a command like `docker run -p 127.0.0.1:7000:80 website`.

This is why you should carefully read the docs before you use a piece of software :P
