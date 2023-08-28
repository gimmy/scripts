# /bin/bash
# Rsync with Poisson
echo 'rsync -auv4 "Desktop/PlanTomorrow.md" "brocchi@poisson:~ and reverse" ' &&  rsync -auv4 --progress "Desktop/PlanTomorrow.md" "brocchi@poisson:~" &&  rsync -auv4 --progress "brocchi@poisson:~/PlanTomorrow.md" "Desktop/" 

