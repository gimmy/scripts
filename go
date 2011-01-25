#!/usr/bin/python

# ToDo: aggiungere controllo parametro -X
# notifica vocale?
# from termcolor import colored # per colori?

import sys, os

try:
	dove = sys.argv[1] # salvo la meta dove andare 
except:
	print ("\t ..dove volevi andare?")
	sys.exit()

# Test della rete 	
#	TEST=$(route -n | awk '{ print $3 | "grep 255.255.255."}'  | sed s/255.255.255.//)
#	if [  "$TEST" == '128' ] # controlla la rete
#	then
#		FROM="PHC"
#	else
#		TEST=$(avahi-browse -at | grep v4 | grep Nash | awk '{ print $5}')
#		#TEST=$() # con un ping?
#		if [  "$TEST" == "[00:60:97:8e:f4:d0]" ]
#		then
#			FROM="casa"
#		fi
#	fi

if dove == "nash":
	print ("Collegamento a Nash in corso...")
	os.system("ssh gimmy@gimmy.homelinux.net")

# Cambio a seconda della rete
#	if [ "$FROM" != "casa" ]
#	then
#		HOST="gimmy.homelinux.net" # imposto host alla meta
#	else
#		HOST="192.168.1.95" # (collegamento sulla rete locale)

if dove == "poisson":
	print ("Collegamento a Poisson...")
	os.system("ssh brocchi@poisson")

#	if [ "$FROM" != "PHC" ]
#	then
#		HOST="brocchi@poisson.phc.unipi.it"
#	else
#		HOST="brocchi@poisson" # (collegamento in loale)
#	fi

if dove == "dm":
	print ("Collegamento a ssh.dm.unipi.it in corso...")
	os.system("ssh brocchi@ssh.dm.unipi.it")
