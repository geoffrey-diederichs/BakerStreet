from selenium import webdriver
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


TARGET = "Erwan Sinck"

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

def fullLookup(brow:webdriver, target:str) -> None:
    facebook = fcbkPublic(brow, target)
    insta = duckResearch(brow, "site:instagram.com \"@\" "+target)
    tiktok = duckResearch(brow, "site:tiktok.com \"@\" "+target)
    twitter = duckResearch(brow, "site:twitter.com \"@\" "+target)
    github = duckResearch(brow, "site:github.com \"@\" "+target)

    print(facebook, insta, tiktok, twitter, github, sep="\n", end="\n")

def main() -> None:
    brow = setUpBrowser(True)
    fullLookup(brow, TARGET)
    
if __name__ == "__main__":
    main()  
