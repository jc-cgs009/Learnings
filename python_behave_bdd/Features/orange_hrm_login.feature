Feature: Orange HRM login testcases

  Scenario:validate login with valid credential
    Given opening the orange hrm website
    When give the valid username and password - "Admin" "admin123"
    Then verify the login is successful

  Scenario Outline: validate login with valid and invalid creds
    Given opening the orange hrm website
    When give the valid username and password - "<username>" "<password>"
    Then verify the login is successful
    Examples:
      | username | password |
      | Admin    | admin123 |
      | Admin1   | admin123 |
      | Admin    | admin1234|