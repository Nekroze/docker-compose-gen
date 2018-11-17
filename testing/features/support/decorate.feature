Feature: Decorate subcommand

    Scenario: By default defines nothing
        When I run `docker-compose-gen decorate --stdout`

        Then it should pass with exactly:
        """
        version: "2.1"
        """

    Scenario: Can specify docker compose file version
        When I run `docker-compose-gen decorate --stdout --compose-version 3`

        Then it should pass with exactly:
        """
        version: "3"
        """

    Scenario: By default adds nothing to given services
        When I run `docker-compose-gen decorate --stdout foo`

        Then it should pass with exactly:
        """
        version: "2.1"
        services:
          foo: {}
        """

    Scenario: Can set the dns for a given service
        When I run `docker-compose-gen decorate --stdout --dns 1.1.1.1 foo`

        Then it should pass with exactly:
        """
        version: "2.1"
        services:
          foo:
            dns: 1.1.1.1
        """

    Scenario: Can set the dns for multiple given services
        When I run `docker-compose-gen decorate --stdout --dns 1.1.1.1 foo bar`

        Then it should pass with exactly:
        """
        version: "2.1"
        services:
          bar:
            dns: 1.1.1.1
          foo:
            dns: 1.1.1.1
        """
