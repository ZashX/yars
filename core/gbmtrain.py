import numpy as np
import pandas as pd
from catboost import CatBoostRanker, Pool
from copy import deepcopy

np.random.seed(42)

train = Pool("gbm/train.csv", column_description="gbm/train.cd", delimiter=',')
test = Pool("gbm/test.csv", column_description="gbm/test.cd", delimiter=',')

default_parameters = {
    'iterations': 500,
    'verbose': True,
    'random_seed': 0,
}

parameters = {}

def fit_model(loss_function, additional_params=None, train_pool=train, test_pool=test):
    parameters = deepcopy(default_parameters)
    parameters['loss_function'] = loss_function
    
    if additional_params is not None:
        parameters.update(additional_params)
        
    model = CatBoostRanker(**parameters)
    model.fit(train_pool, eval_set=test_pool)
    
    return model

model = fit_model('YetiRank')
model.save_model('gbm/model.cbm')