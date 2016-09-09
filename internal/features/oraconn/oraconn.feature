Feature: Connect to oracle
  Scenario:
    Given a running oracle instance
    When provide a connection string for the running instance
    Then a connection is returned
    And I can select the system timestamp from dual

  Scenario:
    Given a connection string with no listener
    When I connect to no listener
    Then an error is returned