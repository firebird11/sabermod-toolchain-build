#!/bin/bash

# Build all with gcc 4.8
bash arm-androideabi-4.8 && bash arm-androideabi-4.9 && bash arm-eabi-4.9;
echo "done with all needed gcc";
