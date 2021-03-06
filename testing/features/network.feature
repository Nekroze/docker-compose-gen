Feature: Network subcommand

    Scenario: By default defines nothing
        When I run `docker-compose-gen network --stdout`

        Then it should pass with exactly:
        """
        version: "2.1"
        """

    Scenario: Can specify docker compose file version
        When I run `docker-compose-gen network --stdout --compose-version 3`

        Then it should pass with exactly:
        """
        version: "3"
        """

    Scenario: Can generate an internal network
        When I run `docker-compose-gen network --stdout --name foo`

        Then it should pass with exactly:
        """
        version: "2.1"
        networks:
          foo: {}
        """

    Scenario: Can generate an external network
        When I run `docker-compose-gen network --stdout --name foo --external bar`

        Then it should pass with exactly:
        """
        version: "2.1"
        networks:
          foo:
            external:
              name: bar
        """

    Scenario: Can generate an external network that replaces the default network
        When I run `docker-compose-gen network --stdout --external bar`

        Then it should pass with exactly:
        """
        version: "2.1"
        networks:
          default:
            external:
              name: bar
        """
