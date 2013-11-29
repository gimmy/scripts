autoload -U compinit promptinit
compinit
promptinit

autoload -U colors && colors

# Git
source ~/.zsh/git-prompt/zshrc.sh	# status in prompt
source ~/.git-completion.sh		# autocomplete git command

# PROMPT='%B%m%~%b$(git_super_status) %# '
PROMPT='%{$fg[green]%}%n%{$reset_color%}%{$fg[red]%}@%{$reset_color%}%{$fg[blue]%}%m %{$fg[yellow]%}%1~%{$reset_color%}$(git_super_status) '
# RPROMPT="[%{$fg[yellow]%}%?%{$reset_color%}]"

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

alias archlinux='lynx www.archlinux.org'
alias pacman='pacman-color'
alias p='pacman'
alias y='yaourt'
