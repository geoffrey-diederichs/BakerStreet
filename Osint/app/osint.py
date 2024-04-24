from selenium import webdriver
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import json
import time
import os

def get_file_path(filename):
    return os.path.abspath(filename)

API = get_file_path("api/api.txt")

def setUpBrowser(options:bool) -> webdriver:
    if (options == True):
        options = Options()
        options.add_argument("--headless")
        brow = webdriver.Firefox(options=options)
    else:
        brow = webdriver.Firefox()

    return brow

def duckResearch(brow:webdriver, request:str) -> [str]:
    brow.get("https://www.duckduckgo.com")
    searchBar = brow.find_element("id", "searchbox_input")
    searchBar.clear()
    searchBar.send_keys(request)
    searchBar.send_keys(Keys.ENTER)
   
    result = []
    WebDriverWait(brow, 10).until(EC.presence_of_element_located((By.CLASS_NAME, "Rn_JXVtoPVAFyGkcaXyK")))
    for i in brow.find_elements(By.CLASS_NAME, "Rn_JXVtoPVAFyGkcaXyK"):
        result.append(str(i.get_attribute("href")))
    
    return result

def fcbkPublic(brow:webdriver, request:str) -> [str]:
    brow.get("https://www.facebook.com/public/"+"-".join(TARGET.split(" ")))
    
    WebDriverWait(brow, 10).until(EC.presence_of_element_located((By.XPATH, "//button[text()='Decline optional cookies']")))
    brow.find_element(By.XPATH, "//button[text()='Decline optional cookies']").click()

    result = []
    WebDriverWait(brow, 10).until(EC.presence_of_element_located((By.CLASS_NAME, "_32mo")))
    for i in brow.find_elements(By.CLASS_NAME, "_32mo"):
        result.append(str(i.get_attribute("href")))
    
    return result

def fullLookup(brow:webdriver, target:str) -> json:
    facebook = [] 
    insta = duckResearch(brow, "site:instagram.com \"@\" "+target)
    tiktok = duckResearch(brow, "site:tiktok.com \"@\" "+target)
    twitter = duckResearch(brow, "site:twitter.com \"@\" "+target)
    github = duckResearch(brow, "site:github.com \"@\" "+target)

    result = {
            "facebook": facebook,
            "instagram": insta, 
            "tiktok": tiktok, 
            "twitter": twitter, 
            "github": github
    }
    return json.dumps(result)

def Lookup() -> None:
    target = "Erwan Sink"
    brow = setUpBrowser(True)
    response = fullLookup(brow, target)


def getTarget() -> (str, int):
    
    # print(API)
    if not os.path.exists(API):
        raise FileNotFoundError(f"The file {API} does not exist.")
    with open(API, "r") as f:
        data = f.read()
        print("data trouvé first step",data)
        data = data.split("\n")
        for i in range(len(data)-1):
            lines = data[i].split(";")
            if len(lines[1]) == 0:
                print("target first stepfound",data[i].split(";")[0], i)
                return data[i].split(";")[0], i
    return "", -1

def writeResult(target: str, line: int, result: str) -> None:
    data = ""
    with open(API, "r") as f:
        data = f.read()
        print("data trouvé second step",data)
    data = data.split("\n")
    data[line] = target+";"+result
    data = "\n".join(data)
    with open(API, "w") as f:
        f.write(data)

def loop() -> None:
    brow = setUpBrowser(True)
    while True:
        time.sleep(1)
        target, line = getTarget()
        print("targe found : ",target)
        if not ((target == "") and (line == -1)):
            print("trying to look",target)
            result = fullLookup(brow, target)
            print("results founds",result)
            writeResult(target, line, str(result))
            print("writing done")

if __name__ == "__main__":
    loop()
