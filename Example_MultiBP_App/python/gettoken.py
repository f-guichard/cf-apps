#!/usr/bin/python
import requests
import os
import json

#Easy design/ KISS : get token via requests
#ConfigMap || VCAP_SERVICES || env 
UserName = os.getenv("USERNAME")
Password = os.getenv("PASSWORD")
Tenant   = os.getenv("TENANT")
AuthApi  = os.getenv("AUTOR_API_URL")

#Should control var values before that :)
AuthentBody = {"username":UserName,"password":Password,"tenant":Tenant}

#Drop TLS checks
#requests.exceptions.SSLError: [Errno 1] _ssl.c:510: error:14090086:SSL routines:SSL3_GET_SERVER_CERTIFICATE:certificate verify failed
#headers = {'Accept':'application/json'}
Token = requests.post(AuthApi, json=AuthentBody, verify=False)

#Dump query
#print vars(Token)
print '========================= Response ==============================='
print Token.text

#Several fix to do, may the courage be with you
try:
    with open('../ressources/token.json', 'w') as TokenFile:
        TokenFile.write(json.dumps(Token.json(), indent=6))
except Exception as exc:
    print(exc)
    exit(-1)

#list catalog items
#https://docs.vmware.com/en/vRealize-Automation/7.3/com.vmware.vra.programming.doc/GUID-41B9F325-31A3-4756-A9E1-F4E10C23E87F.html
VraResponse = Token.json()
headers = {'Authorization':'Bearer '+VraResponse['id']}
Catalog = requests.get('CatalogUrl', headers=headers, verify=False)

#Dump query
#print vars(Catalog)
print '========================= Response ==============================='
print Catalog.content
print Catalog.json()

CatalogResponse = Catalog.json()
print 'Contenu Catalogue : ' + CatalogResponse['content'][0]['description']
print '======================== Get Network Profile ====================='
NetworkProfiles = requests.get('vra/iaas-proxy-provider/api/network/profiles', headers=headers, verify=False)
print NetworkProfiles.content
print NetworkProfiles.json()
