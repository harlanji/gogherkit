Feature: The Beach

  Scenario: 90 degrees
    Given today it is 100 degrees at the beach
    When the time reaches 1200 
    Then the kids are at the beach
  
  Scenario: 90 degrees again
    Given today it is 100 degrees at the beach
    And some other awesome stuff
    When the time reaches 1200 
    Then the kids are at the beach
