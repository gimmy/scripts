#!/bin/bash
# ToDo: aggiungere controllo parametro -X
# controllo rete - notifica vocale?

if [ -z $1 ]
then
	echo -e "\033[01;31m \t ...dove volevi andare?\033[00;38m"
#	echo  "\033[00;34mprova testo a seguire"
#	echo -e "\033[37m Dove volevi andare?"
else

	META=$1 #salvo la meta dove andare

if [ $META == "casa" ]
then
	echo -e "\033[34m \t Collegamento a Nash in corso...\033[00;38m"
#	HOST="gimmy.homelinux.net" # imposto host alla meta
	HOST="192.168.1.95" # (collegamento in locale: controllare impostazioni di rete)

fi

if [ $META == "poisson" ]
then
	echo -e "\033[01;33m \t Collegamento a Poisson in corso...\033[00;38m"
	HOST="brocchi@poisson.phc.unipi.it"
#	HOST="brocchi@poisson" # (collegamento in loale)

fi

if [ $META == "dm" ]
then
	echo "Collegamento a dm.unipi.it in corso..."
	HOST="brocchi@ssh.dm.unipi.it"
fi

ssh $HOST
fi
