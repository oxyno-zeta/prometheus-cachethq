#!/bin/bash

pip3 install pre-commit==2.15.0

pre-commit install
pre-commit install --hook-type commit-msg
