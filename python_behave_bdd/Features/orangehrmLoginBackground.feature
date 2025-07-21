# background : setup -- before executing scenario's if u want to execute something then we can
# use background.

# Allure report
  # pip install allure-behave
  # cmd - behave -f allure_behave.formatter:AllureFormatter -o "report_folder_name" featureFolder/
  # --> after running above command, A folder[report_folder_name] will create which contains the
  # report of all the tc's in json format.
  # cmd to generate allure report [ui] --> allure serve folder[report_folder_name]


Feature: Orange HRM login testcases
  Background:
    Given opening the orange hrm website

  @sanity
  Scenario:validate login with valid credential

    When give the valid username and password - "Admin" "admin123"
    Then verify the login is successful

  # data driven framework
  @suhas
  Scenario Outline: validate login with valid and invalid creds
    When give the valid username and password - "<username>" "<password>"
    Then verify the login is successful
    Examples:
      | username | password |
      | Admin    | admin123 |
      | Admin1   | admin123 |
      | Admin    | admin1234|