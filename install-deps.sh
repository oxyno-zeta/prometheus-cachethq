#!/bin/bash

pip3 install pre-commit==2.7.1

pre-commit install
pre-commit install --hook-type commit-msg
