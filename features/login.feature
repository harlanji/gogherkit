Feature: Login System

  Scenario: I am a new user and try to login before registering

    Given the login system with no data
    When I login with username bob and password 1234
    Then the result should be false

  Scenario: I register my account and log in successfully

    Given the login system with no data
    When I register with username bob and password 1234
    Then the result should be true
    When I login with username bob and password 1234
    Then the result should be true

  Scenario: I register my account and log in with the wrong password

    Given the login system with no data
    When I register with username bob and password 1234
    Then the result should be true
    When I login with username bob and password 1234
    Then the result should be true
