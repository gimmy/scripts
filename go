#!/usr/bin/python
# baco da risolvere su controllo nash con ping
# ToDo: aggiungere controllo parametro -X, notifica vocale?

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
	else:
		connesso_da = "fuori"

print "Parto...connessi da %s" % colora(4, connesso_da)

# Impostazioni connessione: Pronti al lancio

if dove == "nash":
	if connesso_da == "casa":
		ping = subprocess.Popen(["ping", "-c 1", "Nash"],stdout=subprocess.PIPE) # faccio un ping a Nash
		output_ping = ping.communicate()[0]
		check = re.findall(r"\d+.\d+.\d+.\d+", output_ping)[0] 
		
		if check == "192.168.1.95": # e controllo l'ip
			connesso_da = "casa" 
			print "Nash collegato"
		else:
			connesso_da = "casa - Nash non collegato"
			sys.exit() 
	else:
		ping = subprocess.Popen(["ping", "-c 1", "gimmy.homelinux.net"],stdout=subprocess.PIPE) # faccio un ping a Nash
		output_ping = ping.communicate()[0]
		check = int(re.findall(r"(\d+)% packet loss", output_ping)[0]) 
		print "Nash sembra essere collegato ( %d %% )" % (100-check)
	# gestire gli errori: check out of range se nash non connesso

# Collegamento efefttivo a Nash
	dove_color = colora(34, dove)
	user = "gimmy"
	# Cambio a seconda di dove sono connesso
	if connesso_da != "casa":
		host = "gimmy.homelinux.net"
	else:
		host = "nash"

if dove == "dm":
	dove_color = colora(35, "ssh.dm.unipi.it")
	user = "brocchi"
	host = "ssh.dm.unipi.it"

# PHC
phc = ["poisson","fourier","daphne"]
for i in phc:
	if dove == i:
		user = "brocchi"
		if connesso_da == "PHC":
			host = dove
			dove_color = colora(4,dove)
		else:
			host = dove+".phc.unipi.it"
			dove_color = colora(33, dove)

mete = ["nash","dm"]+phc
if not dove in mete:
	print "\tNon ho ben capito dove sia %s" % colora(4, dove)
	sys.exit()

# Lancio
print "Collegamento a %s --> ssh %s@%s" % (dove_color, colora(31,user), colora(32,host))
os.system("ssh %s@%s" % (user, host))
