#!/usr/bin/python

import sys
import paramiko

try:
	dove=sys.argv[1] # prende in argomento dove si vuole andare
except:
	print "Dove vuoi andare?" # se non ci sono argomenti

ssh = paramiko.SSHClient()

if dove=="casa":
	ssh.load_system_host_keys() # importo le host key dal sistema
	ssh.connect('192.168.1.95', username='gimmy') # mi collego in locale 
	ssh.invoke_shell # apre una shell remota (?)
	stdin, stdout, stderr = ssh.exec_command("ls") # lanciamo un comando di prova
else:
	print "Mmm..."
