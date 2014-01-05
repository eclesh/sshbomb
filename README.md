sshbomb
=======

sshbomb pretends to be sshd and sends random garbage to anyone who
dares connect. It's a dumb way to mess with script kiddies.

Usage
=====

If you don't want to run sshbomb as root, set up an iptables redirect
to forward connections from port 22 to port 2222:

	# iptables -A INPUT -i eth0 -p tcp --dport 22 -j ACCEPT
	# iptables -A INPUT -i eth0 -p tcp --dport 2222 -j ACCEPT
	# iptables -A PREROUTING -t nat -i eth0 -p tcp --dport 22 -j REDIRECT --to-port 2222

Now launch the program with your favorite service manager (nohup,
runit, whatever). With the default options sshbomb will pretend to be
SSH-2.0-OpenSSH_5.3 when accepting connections. It will spew 1MB of
junk.

	# nohup sshbomb -address="example.com:2222" 1>sshbomb.log 2>&1 &
	# telnet example.com 2222
	Trying 127.0.0.1...
        Connected to example.com.
        Escape character is '^]'.
        SSH-2.0-OpenSSH_5.37
	?????[random gibberish forever]???????????????????????
	??????????????????????????????????????????????????????

Options
=======

-advertise=true: advertise sshd version by displaying banner

-banner="SSH-2.0-OpenSSH_5.3": sshd banner to present if advertise flag is set

-listen=":2222": address to listen on

-size=1048576: size in bytes of data to send
