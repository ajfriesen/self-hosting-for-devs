# Code for the Self-Hosting for busy Web Developers Book

> [!NOTE]
> This is a work in progress

The accommodating book will come out soon.
Give this repo a star and you will be the first to know!

## cloud-init

Take this [cloud-init](cloud-init/cloud-init.j2) template and put it into you VPS provider of choice who supports `cloud-init`.
Adjust the first couple of values and you will get a server with these things taken care of:

- [x]Create your user
  - [x] with your username
  - [x] add the user to groups users, admin, docker, systemd-journal
  - [x] set the shell to bash
  - [x] allow sudo without password, can run commands as ALL
  - [x] add your public ssh key <PUBLIC_SSH_KEY>
- [x] Set timezone
- [x] Update package repositories
- [x] Upgrade packages
- [x] Install and configure automatic update and upgrades for every package
- [x] Configure automatic reboots for Kernel updates
- [x] Install and configure sending notification mails via Gmail SMTP when upgrade happened
- [x] Install and configure automatic updates
- [x] Add Docker repository
- [x] Install `docker` and `docker-compose`
- [x] Create the `docker-proxy` network used by your proxy and applications
- [x] Configure Docker to use `journald` as logging driver
- [x] Install and enable `fail2ban`
- [x] Install and configure `ufw` as your firewall
- [x] Disable root login
- [x] Set maximum `ssh` auth attempts to 2
- [x] Allow only specific users for login
- [ ] Create folders and files for Caddy (your reverse proxy) and watchtower (container update service)
- [x] At last reboot to take care of changes an possible kernel updates
