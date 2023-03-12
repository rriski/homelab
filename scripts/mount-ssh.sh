#!/usr/bin/env bash

# Using `-v $HOME/.ssh:/root/.ssh:ro` produce permissions error while in the container
# when working from Linux and maybe from Windows.
# To prevent that we offer the strategy to mount the `.ssh` folder with
# `-v $HOME/.ssh:/.ssh:ro` thus this entrypoint will automatically handle problem.

if [[ -d /.ssh ]]; then

  cp -R /.ssh /root/.ssh
  chmod 700 /root/.ssh &>/dev/null
  chmod 600 /root/.ssh/* &>/dev/null
  if compgen -G "/.ssh/*.pub" > /dev/null; then
    chmod 644 /root/.ssh/*.pub &>/dev/null
  fi
  chmod 644 /root/.ssh/known_hosts &>/dev/null

fi

exec "$@"
