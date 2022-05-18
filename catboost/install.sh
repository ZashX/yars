#!/bin/bash
git clone https://github.com/catboost/catboost.git
cd catboost/catboost/libs/model_interface && ../../../ya make -r .
export CATBOOST_DIR=$(pwd)
export C_INCLUDE_PATH=$CATBOOST_DIR:$C_INCLUDE_PATH
export LIBRARY_PATH=$CATBOOST_DIR:$LIBRARY_PATH
export LD_LIBRARY_PATH=$CATBOOST_DIR:$LD_LIBRARY_PATH