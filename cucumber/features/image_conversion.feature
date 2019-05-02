Feature: Image Conversion

    Imageswitch's primary purpose is to convert images from one format to another

    Scenario: Convert a jpg to png
        Given a base jpg image
        When that image is converted to a png
        Then a png must be returned

    Scenario: Convert a bmp to png
        Given a base bmp image
        When that image is converted to a png
        Then a png must be returned
