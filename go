#!/usr/bin/python

# ToDo: aggiungere controllo parametro -X
# notifica vocale?
# from termcolor import colored # per colori?

import sys, os
import re, subprocess

try:
	dove = sys.argv[1] # salvo la meta dove andare 
except:
	print ("\t ...dove volevi andare?")
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
	                	connesso_da = "casa - Nash collegato"
			else:
				connesso_da = "casa - Nash non collegato"
				sys.exit() 
		# gestire gli errori: check out of range se nash non connesso
	else:
		connesso_da = "fuori"
print "Parto...connessi da %s" % connesso_da

if dove == "nash":
	print ("\t Collegamento a Nash in corso...")
	# Cambio a seconda di dove sono connesso
	if connesso_da != "casa":
		os.system("ssh gimmy@gimmy.homelinux.net")
	else:
		os.system("ssh gimmy@nash")

if dove == "poisson":
	print ("\t Collegamento a Poisson...")
	if connesso_da != "PHC":
		os.system("ssh brocchi@poisson.phc.unipi.it")
	else:
		os.system("ssh brocchi@poisson")

if dove == "fourier":
	print ("\t Collegamento a Fourier...")
	if connesso_da == "PHC":
		os.system("ssh brocchi@fourier")
	else:
		print ("\t ...ehm, non sei in PHC")

if dove == "dm":
	print ("\t Collegamento a ssh.dm.unipi.it in corso...")
	os.system("ssh brocchi@ssh.dm.unipi.it")
