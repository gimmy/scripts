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
p = subprocess.Popen(["ip", "address", "show", "eth1"],stdout=subprocess.PIPE)
output = p.communicate()[0]
netmask = re.findall(r"(\d+.\d+.\d+.\d+)(/\d+)", output)[0] # controllo la maschera di rete   

if netmask[1] == "/128":
        connesso_da = "PHC"
else:
	if netmask[0] == "129.168.1.136" and netmask[1] == "/24": # controllo sull'ip di casa
		connesso_da = "casa"
	# gestire gli errori: check out of range se nash non connesso
#        p = subprocess.Popen(["ping", "-c 1", "Nash"],stdout=subprocess.PIPE) # faccio un ping a Nash
#        output = p.communicate()[0]
#        check = re.findall(r"\d+.\d+.\d+.\d+", output)[0] 
#        if check == "192.168.1.95": # e controllo l'ip
#                connesso_da = "casa"
	else:
		connesso_da = "fuori"
print "Parto.. ..connessi da %s" % connesso_da

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

if dove == "dm":
	print ("\t Collegamento a ssh.dm.unipi.it in corso...")
	os.system("ssh brocchi@ssh.dm.unipi.it")
