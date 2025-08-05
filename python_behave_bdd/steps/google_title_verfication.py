from behave import *
from selenium import webdriver

@given('Opening browser')
def open_browser(context):
    context.driver = webdriver.Chrome()


@when('Providing google url in the browser')
def provide_google_url(context):
    context.driver.get("https://www.google.com/")


@then('verify title of the google page')
def verify_google_title(context):
    title = context.driver.title
    context.driver.close()
    assert title == "Google"
