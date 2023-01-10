FROM archlinux:latest

RUN pacman -Syu --noconfirm &&\
    pacman -S --noconfirm cpupower s3fs-fuse monit &&\
    yes | pacman -Scc
