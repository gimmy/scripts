#!/bin/bash
# ToDo: aggiungere controllo parametro -X
# controllo rete - notifica vocale? 

if [ -z $1 ]
then
	echo -e "\033[01;31m \t ...dove volevi andare?"
#	echo -e "\033[31m Dove volevi andare?"
else

	META=$1 #salvo la meta dove andare

if [ $META == "casa" ]
then
	echo "Collegamento a Nash in corso..."
#	HOST="gimmy.homelinux.net" # imposto host alla meta
	HOST="192.168.1.95" # (collegamento in locale: controllare impostazioni di rete)

fi

if [ $META == "poisson" ]
then
	echo "Collegamento a Poisson in corso..."
	HOST="poisson.phc.unipi.it"
#	HOST="poisson" # (collegamento in loale)

fi

if [ $META == "dm" ]
then
	echo "Collegamento a dm.unipi.it in corso..."
	HOST="ssh.dm.unipi.it"
fi

ssh $HOST

fi
