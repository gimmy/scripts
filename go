#!/bin/bash
# ToDo: aggiungere controllo parametro -X
# controllo rete

meta=$1 #salvo la meta dove andare

if meta!= "casa" || "poisson" || "dm":
	echo "Dove vuoi andare?"

if meta=="casa":
	host="gimmy.homelinux.net" # imposto host alla meta
#	host="192.168.1.95" # (collegamento in locale: controllare impostazioni di rete)

if meta=="poisson":
	host="poisson.phc.unipi.it"
#	host="poisson" # (collegamento in loale)

if meta=="dm":
	host="ssh.dm.unipi.it"
else:
	echo "Nulla"

ssh $host

fi
