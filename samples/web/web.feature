Feature: The web page displays a greeting

  Scenario: I request the page with the name bob
    Given the web server is running
    When I request the main page with name 'bob'
    Then the text that comes back contains 'Hello there, I love bob!'

  Scenario: I request the page with no name
    Given the web server is running
    When I request the main page with name ''
    Then the text that comes back contains 'Hello there, I love !'

