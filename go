#!/usr/bin/python

# ToDo: aggiungere controllo parametro -X, compattare in lista
# notifica vocale?

import sys, os
import re, subprocess

def colora(i, string):
	return "\033[%d;1m%s\033[0m" % (i, string)
try:
	dove = sys.argv[1] # salvo la meta dove andare 
except:
	print(colora(31, "\t...dove volevi andare?"))
	sys.exit()

# Controllo dove sono connesso
p = subprocess.Popen(["ip", "route"],stdout=subprocess.PIPE)
output = p.communicate()[0]
netmask = re.findall(r"(\d+.\d+.\d+.\d+)/(\d+)", output)[0] # controllo la maschera di rete   

q = subprocess.Popen(["ip", "addr", "show", "eth1"],stdout=subprocess.PIPE)
output_ip = q.communicate()[0]
ip = re.findall(r"(\d+.\d+.\d+.\d+)/(\d+)", output_ip)[0] # controllo l'ip asseggnato   

if netmask[0] == "131.114.10.0": # controllo con ip
#if netmask[1] == "128": # controllo con maschera di rete
        connesso_da = "PHC"
else:
	if ip[0] == "192.168.1.136" and ip[1] == "24": # controllo sull'ip di casa
		connesso_da = "casa"
		if dove == "nash":
		        ping = subprocess.Popen(["ping", "-c 1", "Nash"],stdout=subprocess.PIPE) # faccio un ping a Nash
		        output_ping = ping.communicate()[0]
		        check = re.findall(r"\d+.\d+.\d+.\d+", output_ping)[0] 
		        if check == "192.168.1.95": # e controllo l'ip
	                	connesso_da = "casa" 
				print "Nash collegato"
			else:
				connesso_da = "casa - Nash non collegato"
				sys.exit() 
		# gestire gli errori: check out of range se nash non connesso
	else:
		connesso_da = "fuori"

print "Parto...connessi da %s" % colora(4, connesso_da)

# Impostazioni connessione: Pronti al lancio

if dove == "nash":
	dove = colora(34, dove)
	user = "gimmy"
	# Cambio a seconda di dove sono connesso
	if connesso_da != "casa":
		host = "gimmy.homelinux.net"
	else:
		host = "nash"

if dove == "poisson":
	dove = colora(4, dove)
	user = "brocchi"
	if connesso_da != "PHC":
		host = "poisson.phc.unipi.it"
	else:
		host = "poisson"
# PHC

phc = ["fourier","daphne"]
for i in phc:
	if dove == i:
		if connesso_da != "PHC":
			print (colora(31, "\t ...ehm, non sei in PHC"))
			sys.exit()
		else:
			user = "brocchi"
			#print "Collegamento a %s --> ssh %s@%s" % (colora(4,dove), colora(31,user), colora(32,dove))
			host = dove

if dove == "dm":
	dove = (4, "ssh.dm.unipi.it")
	user = "brocchi"
	host= "ssh.dm.unipi.it"

#if dove !="nash" and dove != "poisson" and dove != "fourier" and dove != "dm":
#if dove !=["nash","poisson","fourier","dm","daphne"]:
#	print "\tNon ho ben capito dove sia %s" % colora(4, dove)
#	sys.exit()

# for i in phc:
# 	if dove != i:
# 		print "\tNon ho ben capito dove sia %s" % colora(4, dove)
# 		sys.exit()

# Lancio
print "Collegamento a %s --> ssh %s@%s" % (dove, colora(31,user), colora(32,host))
os.system("ssh %s@%s" % (user, host))
