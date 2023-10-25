from selenium import webdriver

brow = webdriver.Firefox()
brow.get("https://www.google.com")

button = brow.find_element("id", "L2AGLb")
brow.execute_script("arguments[0].click();", button)
