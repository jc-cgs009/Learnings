from behave import *
from selenium import webdriver
from selenium.webdriver.common.by import By


@given('open orange hrm website')
def open_orangehrm_website(context):
    context.driver = webdriver.Chrome()
    context.driver.implicitly_wait(10)
    context.driver.get("https://opensource-demo.orangehrmlive.com/web/index.php/auth/login")


@when('Login with using creds in query params')
def login_with_creds_in_table_data(context):
    for r in context.table:
        print(r["username"])
        context.driver.find_element(By.NAME, "username").send_keys(r["username"])
        context.driver.find_element(By.NAME, "password").send_keys(r["password"])
        context.driver.find_element(By.XPATH, "//button[@type='submit']").click()

@then('validate successful login')
def validate_successful_login(context):
    try:
        user_name = context.driver.find_element(By.XPATH, "//p[text()='Zafar Qaimkhani']").text
        print(user_name)
    except:
        assert False, "Login failed bcz of invalid creds"
    else:
        assert "Zafar Qaimkhani" == user_name
        context.driver.quit()
