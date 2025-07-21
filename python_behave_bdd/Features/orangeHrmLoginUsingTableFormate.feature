# table formate

Feature: Orange HRM login page testcase

  Scenario: Validate login with table data
    Given open orange hrm website
    When Login with using creds in query params
      | username | password |
      | Admin    | admin123 |
#      | Admin123 | addsfj   | # StaleElementReferenceError
    Then validate successful login