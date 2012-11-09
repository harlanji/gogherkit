Registration and Login

Scenario: Keith buys his first case and registers it.

Given the reference data
When Keith attempts to login with email 'keith@emota.net' and password '1234' and name 'Keith'
Then the login result should be positive
And Keith should exist in the database
And Keith should have password '1234' and name 'Keith'


Scenario: Stephen gets a new phone and must re-login.

Given the reference data
When Stephen attempts to login with their reference credentials
Then the login result should be positive


Scenario: Alexa is Jealous of Tracy and tries to guess her password.

Given the reference data
When Alexa attempts to login with email 'tracy@emota.net' and password 'wrong' and name 'Tracy'
Then the login result should be negative


Scenario: Stephen and Tracy just got married, so Tracy logs in with her new last name

Given the reference data
When Tracy attempts to login with email 'tracy@emota.net' and password '1234' and name 'Tracy GotMarried'
Then the login result should be positive
And tracy@emota.net should have password '1234' and name 'Tracy GotMarried'   


Scenario: Stephen got a new email address, so he changes it

Given the reference data
When Stephen attempts to login with their reference credentials
And Stephen changes their email to 'stephen2@example.com'
Then stephen@emota.net shouldn't exist in the database
And stephen2@example.com should exist in the database
And stephen2@example.com should have password '1234' and name 'Stephen'


Scenario: Stephen would like to change his password

Given the reference data
When Stephen attempts to login with their reference credentials
And Stephen changes their password to '5678'
Then the login result should be positive
And Stephen should have password '5678' and name 'Stephen'

