import os
import sys
import requests
import os
import json

def run(st,**kargs):
    b = json.loads(st)
    d = {"FireBrand":{}}
    d["FireBrand"]['Fire'] = "dog"
    d["FireBrand"]["ip"] = "172.0.0.1"
    d["FireBrand"]["string"] = "NewString"
    b.append(d)
    return json.dumps(b)
