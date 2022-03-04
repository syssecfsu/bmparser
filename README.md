# bmparser

This repo contains addition files for the BMP parser project for CIS5370. 
It provides a non-compliant implementation of bmparser and a viewer. 

> This implementation of bmparser is non-compliant because it uses the 
> go/image package to parse BMP files while the project forbids the use
> of such packages/libraries. 
> 
> In addition, bmparser only supports bmp files with 8, 16, and 24 bpp 
> BMP files. 

To build the programs on Linux, make sure you have a recent version of 
go and c++ compiler installed. You also need to install the following dependencies.

```# apt install libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libglx-dev libgl1-mesa-dev libxxf86vm-dev```
