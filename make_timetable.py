#!/usr/env/python

print ('Ciao, sono qui per creare il tuo orario delle lezioni')

# Variabili
lunedi=['','','','']
martedi=['','','','']
mercoledi=['','','','']
giovedi=['','','','']
venerdi=['','','','']

# Testo la scrivere
layout = \
'''
\documentclass[a4paper,landscape]{article}
\usepackage[landscape,a4paper,left=2cm]{geometry}
\usepackage[utf8x]{inputenc}
\usepackage{multirow} 
\usepackage{tabularx} 
\begin{document}
 \begin{center} \huge \textsc{Orario I semestre 2011-2012} \end{center}
 \thispagestyle{empty}
\vspace{1cm}
	\begin{tabularx}{\columnwidth}{l|>{\textbf}X|>{\textbf}X|>{\textbf}X|>{\textbf}X|>{\textbf}X}
   & \textsc{Lunedi} & \textsc{Martedi} & \textsc{Mercoledi} & \textsc{Giovedi} & \textsc{Venerdi} \\hline \hline
		9 - 11 & %s & %s & %s & %s & %s  \\ \hline
		11 - 13 & %s & %s & %s & %s & %s \\ \hline
		14 - 16 & %s & %s & %s & %s & %s \\ \hline
		16 - 18 & %s & %s & %s & %s & %s \\ \hline
	\end{tabularx}\\ 
\end{document}
''' % (lunedi[0], martedi[0], mercoledi[0], giovedi[0], venerdi[0], \
lunedi[1], martedi[1], mercoledi[1], giovedi[1], venerdi[1], \
lunedi[2], martedi[2], mercoledi[2], giovedi[2], venerdi[2], \
lunedi[3], martedi[3], mercoledi[3], giovedi[3], venerdi[3])

orario = open('orario.tex','w')
orario.write ("%s" % layout)
orario.close ()
