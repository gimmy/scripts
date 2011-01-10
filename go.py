#!/usr/bin/python

import sys
import paramiko

#from termcolor import colored # per colori?

try:
	dove=sys.argv[1] # prende in argomento dove si vuole andare
except:
	print ("\t Dove vuoi andare?") # se non ci sono argomenti
	sys.exit()

ssh = paramiko.SSHClient()

if dove=="casa":
	ssh.load_system_host_keys() # importo le host key dal sistema
	ssh.connect('192.168.1.95', username='gimmy') # mi collego in locale 
	ssh.invoke_shell # apre una shell remota (?)
	stdin, stdout, stderr = ssh.exec_command("riaggiorna") # lanciamo un comando di prova

if dove=="mamma":
	ssh.load_system_host_keys() # importo le host key dal sistema
	ssh.connect('192.168.1.47', username='gimmy')

else:
	print "Mmm..."
