workspace "Getting Started" "This is a model of my software system." {

    model {
        ss = softwareSystem "test" {
            app = container "app" {
                619983575 = component "command.handleTripReleasedHandler" "application command" ""  "APP"
                2609504327 = component "generator.TripInvoiceSnapshotGenerator" "domain component" ""  "DOMAIN"
                3816622744 = component "command.generateTripDocumentPDFHandler" "application command" ""  "APP"
                847409183 = component "command.abortTripDocumentVoidingHandler" "application command" ""  "APP"
                1611456128 = component "command.completeTripDocumentVoidingHandler" "application command" ""  "APP"
                3061501813 = component "query.getInvoiceSnapshotHandler" "application query" ""  "APP"
                4058262459 = component "command.issueTripDocumentsHandler" "application command" ""  "APP"
                1817426477 = component "query.getPrintedDocumentHandler" "application query" ""  "APP"
                19188921 = component "query.getPrintedDocumentOptionsHandler" "application query" ""  "APP"
                2388691745 = component "partnerdetails.MockRepository" "adapter repository" ""  "ADAPTER"
                1438553552 = component "tripdocuments.MockRepository" "adapter repository" ""  "ADAPTER"
                1867705285 = component "generator.InvoiceSnapshotGenerators" "domain component" ""  "DOMAIN"
                4199524036 = component "generator.FeeInvoiceSnapshotGenerator" "domain component" ""  "DOMAIN"
                1819222398 = component "generator.TripReceiptSnapshotGenerator" "domain component" ""  "DOMAIN"
                285429996 = component "command.requestTripDocumentVoidingHandler" "application command" ""  "APP"
                2411668577 = component "command.reissueTripDocumentHandler" "application command" ""  "APP"
                285429996 -> 1438553552
                847409183 -> 1438553552
                1611456128 -> 1438553552
                3061501813 -> 1438553552
                1867705285 -> 4199524036
                1867705285 -> 2609504327
                1867705285 -> 1819222398
                4199524036 -> 2388691745
                2609504327 -> 2388691745
                3816622744 -> 1438553552
                1817426477 -> 1438553552
                19188921 -> 1438553552
                4058262459 -> 2388691745
                4058262459 -> 1438553552
                4058262459 -> 1867705285
                1819222398 -> 2388691745
                2411668577 -> 2388691745
                2411668577 -> 1438553552
                2411668577 -> 1867705285
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