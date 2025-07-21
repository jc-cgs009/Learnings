from behave import given, when, then
from selenium import webdriver
from selenium.webdriver.common.service import Service
from webdriver_manager.chrome import ChromeDriverManager


@given('Opening browser')
def open_browser(context):
    # service = Service(ChromeDriverManager().install())
    context.driver = webdriver.Chrome()


@when('Providing google url in the browser')
def providing_google_url(context):
    context.driver.get("https://www.google.com")


@then('verify title of the google page')
def verify_title(context):
    assert "Google" == context.driver.title
