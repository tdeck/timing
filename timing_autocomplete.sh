#!/bin/bash
# Setup: ln -s timing_autocomplete.sh /etc/bash_completion.d/
# (or add to .bashrc)

_recent_timed_tasks()
{
  local cur opts
  cur="${COMP_WORDS[COMP_CWORD]}"
  opts=$(daylog | tail -n15 | grep $'^\t.*\t' | cut -f 3 | uniq)
  COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
}
complete -F _recent_timed_tasks n
