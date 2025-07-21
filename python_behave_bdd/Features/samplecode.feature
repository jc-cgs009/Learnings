Feature: Sample code for testing

  Scenario: Opening google page and verify title
    Given Opening browser
    When Providing google url in the browser
    Then verify title of the google page