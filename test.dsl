workspace "Getting Started" "This is a model of my software system." {

    model {
        ss = softwareSystem "test" {
            app = container "app" {
                scraper_test_command_Command = component "command" "" ""
                scraper_test_Repository_Repository = component "repository" "" ""
                scraper_test_Client_Client = component "client" "" ""
                scraper_test_commandBis_command = component "command-bis" "" ""
                scraper_test_query_Query = component "query" "" ""
                scraper_test_commandBis_command -> scraper_test_Repository_Repository
                scraper_test_commandBis_command -> scraper_test_Client_Client
                scraper_test_query_Query -> scraper_test_Repository_Repository
                scraper_test_query_Query -> scraper_test_Client_Client
                scraper_test_command_Command -> scraper_test_Client_Client
                scraper_test_command_Command -> scraper_test_Repository_Repository
            }
        }
    }

    views {
        component app "Test" {
            include *
            autoLayout
        }

        styles {
            element "Software System" {
                background #1168bd
                color #ffffff
            }
            element "Person" {
                shape person
                background #08427b
                color #ffffff
            }
        }
    }
}