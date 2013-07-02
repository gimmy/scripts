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

# Faccio le pulizie e il caffè
pulisco = "rm -f "+location+".sql"
print pulisco
os.system(pulisco)
