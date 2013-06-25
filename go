#!/usr/bin/python
# TODO: -p per ping prima di ssh (f.e. a casa), notifica vocale?

'''
go
A simple script to lookup where I am 
in order to connect to ssh host faster.

'''

import sys, os
import re, subprocess

def colora(i, string):
	return "\033[%d;1m%s\033[0m" % (i, string)
try:
	dove = sys.argv[1] # salvo la meta dove andare 
except:
	print(colora(31, "\t...dove volevi andare?"))
	sys.exit()

try:
	opz = sys.argv[2] # parametri opzionali
	if opz != 'X':
		print(colora(31, "\t...come prego?"))
		sys.exit()		
except:
	opz = ''

# Controllo dove sono connesso
sub = subprocess.Popen(["ip", "route"],stdout=subprocess.PIPE)
output = sub.communicate()[0]
# netmask = re.findall(b"(\d+.\d+.\d+.\d+)/(\d+) dev tap0", output)[0] # controllo la maschera di rete   

if output == "": # controllo connessione
	print(colora(31, "\t...sei collegato?"))
	sys.exit()		

netmask = re.findall(b"\d+.\d+.\d+.\d+", output)[0] # controllo la maschera di rete   

if netmask.decode('utf-8') == "10.167.60.0": # controllo l'ip per la vpn
	print(colora(31, "\t...hai attivato la vpn?"))
	sys.exit()		

if netmask.decode('utf-8') == "131.114.10.1":
#if netmask[1] == "128": # controllo con maschera di rete
        connesso_da = "PHC"
else:
	q = subprocess.Popen(["ip", "addr", "show", "wlan0"],stdout=subprocess.PIPE)
	output_ip = q.communicate()[0]
	ip = re.findall(b"(\d+.\d+.\d+.\d+)/(\d+)", output_ip)[0] # controllo l'ip assegnato   
	
	# decode for python 3
	if ip[0].decode('utf-8') == "192.168.1.136" and ip[1].decode('utf-8') == "24": # controllo sull'ip di casa
		connesso_da = "casa"
	else:
		connesso_da = "fuori"

print ('Parto...connessi da %s' % colora(4, connesso_da))

# Impostazioni connessione: Pronti al lancio
if dove == "nash": 
	if connesso_da == "casa":
		ping = subprocess.Popen(["ping", "-c 1", "Nash"],stdout=subprocess.PIPE) # faccio un ping a Nash
		output_ping = ping.communicate()[0]
		check = re.findall(r"\d+.\d+.\d+.\d+", output_ping)[0] 
		
		if check == "192.168.1.95": # e controllo l'ip
			connesso_da = "casa" 
			print ("Nash collegato")
		else:
			connesso_da = "casa - Nash non collegato"
			sys.exit() 
	else:
		print ("Non sei connesso a casa")
		sys.exit()

if dove in ["casa", "pi"]: 
	# if connesso_da == "casa":
	# 	ping = subprocess.Popen(["ping", "-c 1", "Russell"],stdout=subprocess.PIPE) # faccio un ping
	# 	output_ping = ping.communicate()[0]
	# 	check = re.findall(r"\d+.\d+.\d+.\d+", output_ping)[0] 
		
	# 	if check == "192.168.1.106": # e controllo l'ip
	# 		connesso_da = "casa" 
	# else:
	# 	ping = subprocess.Popen(["ping", "-c 1", "gimmy.homelinux.net"],stdout=subprocess.PIPE)
	# 	output_ping = ping.communicate()[0]
	# 	check = int(re.findall(r"(\d+)% packet loss", output_ping)[0]) 
	# 	print "Nash collegato al %d %%" % (100-check)
	# gestire gli errori: check out of range se nash non connesso

	dove_color = colora(34, dove)
	user = "pi"

	if connesso_da != "casa":  # Cambio a seconda di dove sono connesso
		host = "brocchi.dyndns.tv"
	else:
		host = "pi"

if dove == "dm":
	dove_color = colora(35, "ssh.dm.unipi.it")
	user = "brocchi"
	host = "ssh.dm.unipi.it"

if dove == "bhalo":
	dove_color = colora(35, "srv-hs11.netsons.net")
	user = "wpagvdfy"
	host = "srv-hs11.netsons.net -p 65100"

phc = {"a":"apollo",
       "b":"brucaliffo",
       "c":"cassiopea",
       "d":"daphne", 
       "e":"escher",
       "f":"fourier",
       "g":"geppetto", 
       "h":"hitchhiker",       
       "p":"poisson"} # PHC

while dove in phc:
	dove = phc[dove] # setto meta
	user = "brocchi" # setto username
	if connesso_da == "PHC":
		host = dove
		dove_color = colora(4,dove)
	else:
		host = dove+".phc.unipi.it"
		dove_color = colora(33, dove)

home = ["nabucco", "russell", "nash", "jarvis", "silvana-laptop", "fede-laptop"] # dizionario anche qui
for i in home:
	if dove == i:
		user = "gimmy"
		if connesso_da == "casa":
			host = dove
			dove_color = colora(36,dove)
		else:
			print ("\t Non sei a casa")
			sys.exit()

phc = ["apollo","brucaliffo","cassiopea","daphne","escher","fourier","geppetto","poisson"]
for i in phc:
	if dove == i:
		user = "brocchi"
		if connesso_da == "PHC":
			host = dove
			dove_color = colora(36,dove)
		else:
			host = dove+".phc.unipi.it"
			dove_color = colora(33, dove)

mete = ["casa", "pi", "dm", "bhalo"]+phc+home
if not dove in mete:
	print ("\tNon ho ben capito dove sia %s" % colora(4, dove))
	sys.exit()

# Lancio
if opz:
	print ("Collegamento a %s --> ssh %s %s@%s" % (dove_color, colora(33,'-X'), colora(31,user), colora(32,host)))
	os.system("ssh -X %s@%s" % (user, host))
else:
	print ("Collegamento a %s --> ssh %s@%s" % (dove_color, colora(31,user), colora(32,host)))
	os.system("ssh %s@%s" % (user, host))
