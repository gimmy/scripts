#!/usr/bin/env python2
# -*- coding: utf-8 -*-
import os, sys, time

# Prendo i dati
database = sys.argv[1]
location = sys.argv[2]
user = sys.argv[3]
pw = sys.argv[4]

# Converto nel tempo effettivo la location
location = location.replace("DAY", time.strftime("%d"))
location = location.replace("MONTH", time.strftime("%m"))
location = location.replace("YEAR", time.strftime("%y"))
location = location.replace("HOUR", time.strftime("%H"))
location = location.replace("MIN", time.strftime("%M"))
location = location.replace("SEC", time.strftime("%S"))

# Scarico
dump = "mysqldump --opt --single-transaction -u "+user+" -p"+pw+" "+database+" > "+location+".sql"
print dump
os.system(dump)

# Comprimo
comprimo = "tar -czvf "+location+" "+location+".sql"
print comprimo
os.system(comprimo)

# Quando salvare anche altrove il backup
giorno = time.strftime("%d")
mese = time.strftime("%m")
anno = time.strftime("%y")
if giorno == "07":
    print "Creo backup mensile in public_ftp/dbackup"
    take_away_location = location.replace("dbackup", "public_ftp/dbackup")
    take_away = "zip -P "+pw+" "+take_away_location+".zip "+location
    print take_away
    os.system(take_away)

# Faccio le pulizie e il caffè
pulisco = "rm -f "+location+".sql"
print pulisco
os.system(pulisco)

# Pulizie di Primavera
print "Cancello backup più vecchi di 20 giorni"
location = location.replace(giorno, "*")
location = location.replace(mese, "*")
location = location.replace(anno, "*")
clean_old = "find "+location+" -mtime +20 -exec rm {} \;"
os.system(clean_old)
