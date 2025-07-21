from behave import *
from selenium import webdriver
from selenium.webdriver.common.by import By


@given('opening the orange hrm website')
def open_orangehrm_website(context):
    context.driver = webdriver.Chrome()
    context.driver.implicitly_wait(5)
    context.driver.get("https://opensource-demo.orangehrmlive.com/web/index.php/auth/login")


@when('give the valid username and password - "{username}" "{password}"')
def login_to_orange_hrm_website(context, username, password):
    context.driver.find_element(By.NAME, "username").send_keys(username)
    context.driver.find_element(By.NAME, "password").send_keys(password)
    context.driver.find_element(By.XPATH, "//button[@type='submit']").click()


@then('verify the login is successful')
def verify_the_login(context):
    try:
        user_name = context.driver.find_element(By.XPATH, "//p[text()='Zafar Qaimkhani']").text
    except:
        assert False, "Login failed bcz of invalid creds"
    else:
        assert "Zafar Qaimkhani" == user_name
        context.driver.quit()
