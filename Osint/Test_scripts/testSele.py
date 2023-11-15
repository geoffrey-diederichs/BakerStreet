from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


brow = webdriver.Firefox()
brow.get("https://www.google.com")

button = brow.find_element("id", "L2AGLb")
brow.execute_script("arguments[0].click();", button)

searchBar = brow.find_element("id", "APjFqb")
searchBar.clear()
searchBar.send_keys("Erwan Sinck") # site:facebook.com")
searchBar.send_keys(Keys.ENTER)

try:
    element = WebDriverWait(brow, 10).until(EC.presence_of_element_located(("id", "res")))
finally:
    print("found")

for i in brow.find_elements(By.CLASS_NAME, "byrV5b"):
    iString = i.text
    if (len(iString) > 0) and (iString.count("›") == 1):
        print("/".join(iString.split(" › ")))

"""
results = brow.find_element("id", "res")
print(results.text)
"""
