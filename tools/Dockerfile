FROM archlinux:latest

RUN pacman -Syu --noconfirm &&\
    pacman -S --noconfirm s3fs-fuse &&\
    yes | pacman -Scc
