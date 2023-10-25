from googleapiclient.discovery import build
import json

apiKey = ""
cseKey = ""

def googleSearch(search, **kwargs):
    service = build("customsearch", "v1", developerKey=apiKey)
    res = service.cse().list(q=search, cx=cseKey, **kwargs).execute()
    return res

result = googleSearch("site:facebook.com \"geoffrey diederichs\"")

print(result)

nbrResult = result["searchInformation"]["totalResults"]
print(nbrResult)

if int(nbrResult) > 0:
    for k in result["items"]:
        print(k["title"], k["link"], end="\n\n")
