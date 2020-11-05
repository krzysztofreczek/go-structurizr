workspace "Getting Started" "This is a model of my software system." {

    model {
        ss = softwareSystem "test" {
            app = container "app" {
                1817426477 = component "query.getPrintedDocumentHandler" "application query" ""  "APP"
                1915510816 = component "fleetfareconfiguration.ClientMock" "adapter client" ""  "ADAPTER"
                1867705285 = component "generator.InvoiceSnapshotGenerators" "domain component" ""  "DOMAIN"
                285429996 = component "command.requestTripDocumentVoidingHandler" "application command" ""  "APP"
                1611456128 = component "command.completeTripDocumentVoidingHandler" "application command" ""  "APP"
                1438553552 = component "tripdocuments.MockRepository" "adapter repository" ""  "ADAPTER"
                2411668577 = component "command.reissueTripDocumentHandler" "application command" ""  "APP"
                3816622744 = component "command.generateTripDocumentPDFHandler" "application command" ""  "APP"
                847409183 = component "command.abortTripDocumentVoidingHandler" "application command" ""  "APP"
                2527900488 = component "contract.ClientMock" "adapter client" ""  "ADAPTER"
                619983575 = component "command.handleTripReleasedHandler" "application command" ""  "APP"
                3476862434 = component "trips.ClientMock" "adapter client" ""  "ADAPTER"
                2694411847 = component "organisations.ClientMock" "adapter client" ""  "ADAPTER"
                2388691745 = component "partnerdetails.MockRepository" "adapter repository" ""  "ADAPTER"
                3061501813 = component "query.getInvoiceSnapshotHandler" "application query" ""  "APP"
                864625067 = component "farebreakdown.ClientMock" "adapter client" ""  "ADAPTER"
                4058262459 = component "command.issueTripDocumentsHandler" "application command" ""  "APP"
                19188921 = component "query.getPrintedDocumentOptionsHandler" "application query" ""  "APP"
                4191222838 = component "creditcardfeerates.ClientMock" "adapter client" ""  "ADAPTER"
                285429996 -> 1438553552
                847409183 -> 1438553552
                1817426477 -> 1438553552
                19188921 -> 1438553552
                4058262459 -> 3476862434
                4058262459 -> 2694411847
                4058262459 -> 2388691745
                4058262459 -> 1438553552
                4058262459 -> 1867705285
                2411668577 -> 1438553552
                2411668577 -> 1867705285
                2411668577 -> 3476862434
                2411668577 -> 2694411847
                2411668577 -> 2388691745
                3061501813 -> 1438553552
                3816622744 -> 1438553552
                3816622744 -> 3476862434
                1611456128 -> 1438553552
            }
        }
    }

    views {
        component app "Test" {
            include *
            autoLayout
        }

        styles {
            element "APP" {
                background #08427b
                color #ffffff
            }
            element "ADAPTER" {
                background #1168bd
                color #ffffff
            }
            element "DOMAIN" {
                background #ffffff
                color #000000
            }
        }
    }
}