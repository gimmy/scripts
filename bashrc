#
# ~/.bashrc
#

# If not running interactively, don't do anything
[[ $- != *i* ]] && return

PS1='[\u@\h \W]\$ '
# alias sudo='sudo '
alias pacman='pacman-color'

alias p='pacman'
alias y='yaourt'

export EDITOR="nano"

# To auto complete after man and sudo
complete -cf sudo
complete -cf s # for shortcut s
complete -cf man
# autocomplete git command
GIT_PS1_SHOWDIRTYSTATE=true
# GIT_PS1_SHOWUNTRACKEDFILES=true
source ~/.git-completion.sh

# My alias
alias grep='grep --color=auto'
alias fgrep='fgrep --color=auto'
alias egrep='egrep --color=auto'

alias ls='ls --color=auto'
alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'

alias s='sudo '
alias ss='s -s'

alias o='xdg-open'
alias diff='diff -y --suppress-common-lines'
alias archlinux='lynx www.archlinux.org'


# Bash History
export HISTSIZE=5000
export HISTFILESIZE=5000
export HISTFILE="$HOME/.bash_history_${HOSTNAME}"
if [ "$UID" != 0 ]; then
    export HISTCONTROL="ignoreboth"   # ignores duplicate lines next to each other and lines with a leading space
    export HISTIGNORE="[bf]g:exit:logout"
fi

# Prompts
#-----------------------------------------------------------------------------
set_prompts() {

# Default
#    PS1='\[\033[34m\]\u\[\033[00m\]\[\033[01;31m\]@\[\033[00m\]\[\033[01;32m\]\h\[\033[00m\]:\[\033[01;34m\][ \W ] \[\033[00m\]\$ '
    PS1='\[\033[34m\]\u\[\033[00m\]\[\033[01;31m\]@\[\033[00m\]\[\033[01;32m\]\h\[\033[00m\]:\[\033[01;34m\][ \W\[\033[01;33m\]$(__git_ps1 " (%s)")\[\033[00m\] \[\033[01;34m\]]\[\033[00m\] \[\033[00m\]\$ '

    export PS1
}
set_prompts
unset -f set_prompts
