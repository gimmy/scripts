#!/usr/bin/python
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
sub = subprocess.Popen(["ip", "route"],stdout=subprocess.PIPE)
output = sub.communicate()[0]
netmask = re.findall(r"(\d+.\d+.\d+.\d+)/(\d+)", output)[0] # controllo la maschera di rete   

if netmask[0] == "131.114.10.0": # controllo l'ip
#if netmask[1] == "128": # controllo con maschera di rete
        connesso_da = "PHC"
else:
	q = subprocess.Popen(["ip", "addr", "show", "wlan0"],stdout=subprocess.PIPE)
	output_ip = q.communicate()[0]
	ip = re.findall(r"(\d+.\d+.\d+.\d+)/(\d+)", output_ip)[0] # controllo l'ip asseggnato   
	
	if ip[0] == "192.168.1.136" and ip[1] == "24": # controllo sull'ip di casa
		connesso_da = "casa"
	else:
		connesso_da = "fuori"

print "Parto...connessi da %s" % colora(4, connesso_da)

# Impostazioni connessione: Pronti al lancio
if dove == "nash": # controllo se Nash e' collegato
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
		print "Nash collegato al %d %%" % (100-check)
	# gestire gli errori: check out of range se nash non connesso

# Collegamento effettivo a Nash
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

phc = {"a":"apollo",
       "b":"brucaliffo",
       "c":"cassiopea",
       "d":"daphne", 
       "e":"escher",
       "f":"fourier",
       "g":"geppetto", 
       "h":"hitchhiker",       
       "p":"poisson"} # PHC

while phc.has_key(dove):
	dove = phc[dove] # setto meta
	user = "brocchi" # setto username
	if connesso_da == "PHC":
		host = dove
		dove_color = colora(4,dove)
	else:
		host = dove+".phc.unipi.it"
		dove_color = colora(33, dove)

home = ["jarvis", "silvana-laptop", "fede-laptop"] # dizionriare anche qui
for i in home:
	if dove == i:
		user = "gimmy"
		if connesso_da == "casa":
			host = dove
			dove_color = colora(36,dove)
		else:
			print "\t Non sei a casa"
			sys.exit()

phc = ["apollo","brucaliffo","cassiopea","daphne","escher","fourier","geppetto","poisson"]
mete = ["nash","dm"]+phc+home
if not dove in mete:
	print "\tNon ho ben capito dove sia %s" % colora(4, dove)
	sys.exit()

# Lancio
print "Collegamento a %s --> ssh %s@%s" % (dove_color, colora(31,user), colora(32,host))
os.system("ssh %s@%s" % (user, host))
