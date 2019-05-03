Feature: Image Conversion

    Imageswitch's primary purpose is to convert images from one format to another

    Scenario: Convert a jpg to png
        Given a base jpg image
        When that image is converted to a png
        Then a png must be returned

    Scenario: Convert a png to jpg
        Given a base png image
        When that image is converted to a jpg
        Then a jpg must be returned
