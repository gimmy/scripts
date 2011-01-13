#!/bin/bash
# ToDo: aggiungere controllo parametro -X
# notifica vocale?

if [ -z $1 ]
then
	echo -e "\033[01;31m \t ...dove volevi andare?\033[00;38m"
#	echo  "\033[00;34mprova testo a seguire"
#	echo -e "\033[37m Dove volevi andare?"
else
	TEST=$(route -n | awk '{ print $3 | "grep 255.255.255."}'  | sed s/255.255.255.//)
	if [  $TEST == '128' ] # controlla la rete
	then
		FROM="PHC"
	else
		TEST=$(avahi-browse -at | grep v4 | grep Nash | awk '{ print $5}')
		#TEST=$() # con un ping?
		if [  $TEST == "[00:60:97:8e:f4:d0]" ]
		then
			FROM="casa"
		fi
	fi

	META=$1 #salvo la meta dove andare

if [ $META == "nash" ]
then
	echo -e "\033[34m \t Collegamento a Nash in corso...\033[00;38m"
	if [ $FROM != "casa" ]
	then
		HOST="gimmy.homelinux.net" # imposto host alla meta
	else
		HOST="192.168.1.95" # (collegamento sulla rete locale)
	fi
fi

if [ $META == "poisson" ]
then
	echo -e "\033[01;33m \t Collegamento a Poisson in corso...\033[00;38m"
	if [ $FROM != "PHC" ]
	then
		HOST="brocchi@poisson.phc.unipi.it"
	else
		HOST="brocchi@poisson" # (collegamento in loale)
	fi
fi

if [ $META == "dm" ]
then
	echo "Collegamento a ssh.dm.unipi.it in corso..."
	HOST="brocchi@ssh.dm.unipi.it"
fi

ssh $HOST
fi
