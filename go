#!/bin/bash
# ToDo: aggiungere controllo parametro -X
# controllo rete

if [ -z $1 ]
then
	echo "Dove vuoi andare?"

meta=$1 #salvo la meta dove andare

if [meta=="casa"]
then
	host="gimmy.homelinux.net" # imposto host alla meta
#	host="192.168.1.95" # (collegamento in locale: controllare impostazioni di rete)

fi

if [meta=="poisson"]
then
	host="poisson.phc.unipi.it"
#	host="poisson" # (collegamento in loale)

fi

if [meta=="dm"]
then
	host="ssh.dm.unipi.it"
fi

else
	echo "Nulla"

#ssh $host

fi
