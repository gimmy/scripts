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
print "Scarico db" #print dump
os.system(dump)

# Comprimo
comprimo = "tar -czvf "+location+" "+location+".sql"
print "Comprimo il db" #print comprimo
os.system(comprimo)

# Quando salvare anche altrove il backup
giorno = time.strftime("%d")
mese = time.strftime("%m")
anno = time.strftime("%y")
if giorno == "01":
    print "Creo backup mensile in public_ftp/dbackup"
    monthly_location = location.replace("dbackup", "public_ftp/dbackup")
    monthly_backup = "zip -P "+pw+" "+monthly_location+".zip "+location
    os.system(monthly_backup)
    ## Take Away ##
    print "preparo il Take Away..."
    take_away_location = "/home/wpagvdfy/public_ftp/take_away/"
    take_away = "cp "+monthly_location+".zip "+take_away_location
    # # Cancello backup più vecchi
    # clean_take_away_location = take_away_location.replace(giorno, "*")
    # clean_take_away_location = take_away_location.replace(mese, "*")
    # clean_take_away_location = take_away_location.replace(anno, "*")
    clean_take_away= "find "+take_away_location+" -mtime +2 -exec rm {} \;"
    print " pulisco.."
    os.system(clean_take_away)
    print " Creo nuovo take away.."
    os.system(take_away)

# Faccio le pulizie e il caffè
pulisco = "rm -f "+location+".sql"
print pulisco
os.system(pulisco)

# Pulizie di Primavera
def spring_clean( n, string ):
    print "Cancello backup più vecchi di n giorni"
    clean_old = "find "+string+" -mtime +%d -exec rm {} \;" % (n, string)
    os.system(clean_old)

location = location.replace(giorno, "*")
location = location.replace(mese, "*")
location = location.replace(anno, "*")
spring_clean( 30, location )
